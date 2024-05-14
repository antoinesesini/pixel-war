package main

import (
	"flag"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"os"
	"strconv"
	"sync"
	"time"
	"utils"
)

func nSecondsSnapshot(n int) {
	// On envoie un message de sauvegarde automatiquement après n secondes
	time.Sleep(time.Duration(n) * time.Second)
	mutex.Lock()
	envoiSequentiel("sauvegarde")
	mutex.Unlock()
}

// Le programme envoie périodiquement des messages sur stdout
func sendPeriodic(nbMessages int, slower bool) {
	val, _ := strconv.Atoi(monNom[1:2])
	for i := 0; i < nbMessages; i++ {
		// On demande l'accès à la section critique
		demandeSC()
		//Le slower permet créer une différence de vitesse entre les sites et accentue la dispute pour la section critique
		//Ici que pour les 2 premiers sites
		if slower {
			if monNom[0:2] == "A1" {
				time.Sleep(time.Duration(3) * time.Second)
			}
			if monNom[0:2] == "A2" {
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
		// Variable val permet d'identifier le site initiateur du message
		envoyerPixel(i, i, 255, val, 0)
		// On libère la section critique
		relacherSC()
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
	utils.DisplayWarning(monNom, "sendPeriodic", "SEND PERIODIC FINIT")
}

func attenteDroit(n int) {
	// Définit une attente de n secondes avant de pouvoir poser un pixel de nouveau
	time.Sleep(time.Duration(n) * time.Second)
	jePeux = true
}

func clicGaucheMatrice(game *Game, positionX int, positionY int, rouge int, vert int, bleu int) {
	// Lors de la pose d'un pixel (clic gauche)
	if jePeux {
		// On demande l'accès à la section critique
		demandeSC()
		// Mise à jour de la matrice locale
		game.UpdateMatrix(positionX, positionY, uint8(rouge), uint8(vert), uint8(bleu))
		// On envoie le pixel posé à app-controle
		envoyerPixel(positionX, positionY, rouge, vert, bleu)
		// On libère la section critique
		relacherSC()
		// On empêche la pose de nouveaux pixels pendant 10 secondes
		jePeux = false
		go attenteDroit(10)
	}
}

func clicGaucheSaveLogo() {
	// Lancement de la sauvegarde lorsque le bouton de sauvegarde est cliqué
	mutex.Lock()
	envoiSequentiel("sauvegarde")
	mutex.Unlock()
}

// Variable globales d'interface
// jePeux donne la permission de poser un pixel
// saveAccess donne la permission de lancer une sauvegarde
var jePeux = true
var saveAccess = true

// Variables globales de répartition
var mutex = &sync.Mutex{}
var pNom = flag.String("n", "base", "nom")
var pPath = flag.String("p", "./sauvegardes", "path")
var monNom string
var cheminSauvegardes string
var accesSC = false // false si on n'est pas en section critique, true si est en section critique

// Variables globales d'utilisation
var pMode = flag.String("m", "a", "mode") //"g" ou "a" pour graphique ou automatique

func main() {
	flag.Parse()
	// On initialise le nom de l'application et le chemin vers le fichier de sauvegarde
	// ainsi que le mode de lancement
	monNom = *pNom + "-" + strconv.Itoa(os.Getpid())
	cheminSauvegardes = *pPath
	modeDeLancement := *pMode
	var game *Game

	//Si l'option m == "g" on lance l'interface graphique, sinon le mode terminal ou automatique
	if modeDeLancement == "g" {
		lancementModeGraphique(game)
	} else {
		lancementModeAutomatique(game)
	}
}

func lancementModeGraphique(game *Game) {
	// On initialise l'interface graphique sous forme d'une matrice
	matrix := make([][]Pixel, 100)
	for y := 0; y < 100; y++ {
		matrix[y] = make([]Pixel, 100)
		for x := 0; x < 100; x++ {
			matrix[y][x] = Pixel{
				R: 255,
				G: 255,
				B: 255,
			}
		}
	}

	//IMAGE
	// On initialise la roue de couleur
	colorWheel, _, err := ebitenutil.NewImageFromFile("app-base/color_wheel.png")
	if err != nil {
		panic(err)
	}

	//BOUTTON SAUVEGARDE
	// On initialise le bouton de sauvegarde
	saveLogo, _, err := ebitenutil.NewImageFromFile("app-base/saveLogo.png")
	if err != nil {
		panic(err)
	}

	// On initialise l'objet game qui sert d'interface graphique
	game = &Game{
		Matrix:        matrix,
		ColorWheel:    colorWheel,
		SaveLogo:      saveLogo,
		SelectedColor: color.RGBA{R: 0, G: 0, B: 0, A: 0xFF},
	}
	// On lance une goroutinede lecture des messages ainsi que l'interface graphique
	go lecture(game)
	err = ebiten.RunGame(game)
	if err != nil {
		return
	}
}

func lancementModeAutomatique(game *Game) {
	//On lance le snapshot sur A1 au bout de 10 secondes (A1 doit être en mode automatique biensûr)
	if monNom[0:2] == "A1" {
		go nSecondsSnapshot(10)
	}

	//On lance un envoi automatique périodique sur la diagonale sur les 2 premiers/seuls sites (ils doivent exister sous ce nom biensûr)
	if monNom[0:2] == "A1" || monNom[0:2] == "A2" {
		go sendPeriodic(20, true)
	}
	// Il n'y a pas d'interface graphique dans le mode automatique
	game = &Game{
		Matrix:        nil,
		ColorWheel:    nil,
		SelectedColor: color.RGBA{},
	}
	go lecture(game)
	//On décide de bloquer le programme principal
	for {
		time.Sleep(time.Duration(60) * time.Second)
	}
}
