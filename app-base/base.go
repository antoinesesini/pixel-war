package main

import (
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"os"
	"strconv"
	"sync"
	"time"
	"utils"
)

func frontend(msg utils.MessagePixel) {
	//SectionCritique := true
	//recupPixelAChangerSansLeChangerNiEnLocalNiEnPrevenantLesAutres
	//x := msg.PositionX
	//y := msg.PositionY
	//r := msg.Rouge
	//g := msg.Vert
	//b := msg.Bleu

	//Keep the pixel
	//PixelTmp := utils.MessagePixel{x, y, r, g, b}
	//demandeSC = envoi d'un message demandeSC à l'app de control$
	//utils.MessageExclusionMutuelle {Type: 0,Estampille: (numSite, Horloge) } SecCrit
	//MessageSC( SecCrit )

	//Wait few minutes
}

// Le programme envoie périodiquement des messages sur stdout
func sendperiodic() {
	for i := 0; i < 4; i++ {
		mutex.Lock()
		envoyerPixel(i, i, 255, 0, 0)
		mutex.Unlock()
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func envoyerPixel(positionX int, positionY int, rouge int, vert int, bleu int) {
	messagePixel := utils.MessagePixel{positionX, positionY, rouge, vert, bleu}
	fmt.Println(utils.MessagePixelToString(messagePixel))
}

// Quand le programme n'est pas en train d'écrire, il lit
func lecture(game *utils.Game) {
	var rcvmsg string

	for {
		fmt.Scanln(&rcvmsg)

		if rcvmsg[0] == uint8('A') { // On traite le message s'il commence par un 'A'
			//Différencier les messages de sections des messages pixels
			//Si message de type debutSC, récup pixel dans le fifo et le changer en local et prévenir les autres
			//Si message pixel faire maj locale
			utils.DisplayError(monNom, "lecture", "Réception de : "+rcvmsg[1:])
			mutex.Lock()
			messagePixel := utils.StringToMessagePixel(rcvmsg[1:])
			utils.DisplayError(monNom, "changerPixel", "Et là bim on change le pixel")
			messageString := utils.MessagePixelToString(messagePixel)
			cr, _ := strconv.Atoi(utils.TrouverValeur(messageString, "R"))
			cb, _ := strconv.Atoi(utils.TrouverValeur(messageString, "B"))
			cg, _ := strconv.Atoi(utils.TrouverValeur(messageString, "G"))
			x, _ := strconv.Atoi(utils.TrouverValeur(messageString, "positionX"))
			y, _ := strconv.Atoi(utils.TrouverValeur(messageString, "positionY"))
			game.UpdateMatrix(x, y, uint8(cr), uint8(cg), uint8(cb))
			mutex.Unlock()
			envoyerPixel(x, y, cr, cg, cb)
		}
		rcvmsg = ""
	}
}

func changerPixel(messagePixel utils.MessagePixel, game utils.Game) {

}

var mutex = &sync.Mutex{}
var pNom = flag.String("n", "base", "nom")
var monNom string

func main() {

	flag.Parse()
	monNom = *pNom + "-" + strconv.Itoa(os.Getpid())

	matrix := make([][]utils.Pixel, 100)
	for y := 0; y < 100; y++ {
		matrix[y] = make([]utils.Pixel, 100)
		for x := 0; x < 100; x++ {
			matrix[y][x] = utils.Pixel{
				R: 255,
				G: 255,
				B: 255,
			}
		}
	}

	colorWheel, _, err := ebitenutil.NewImageFromFile("app-base/color_wheel.png")
	if err != nil {
		panic(err)
	}

	game := &utils.Game{
		Matrix:        matrix,
		ColorWheel:    colorWheel,
		SelectedColor: color.RGBA{R: 0, G: 0, B: 0, A: 0xFF},
	}

	//Création de 2 go routines qui s'exécutent en parallèle
	go sendperiodic()
	go lecture(game)

	ebiten.RunGame(game)

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Pixel-War")

	//On décide de bloquer le programme principal
	for {
		time.Sleep(time.Duration(60) * time.Second)
	} // Pour attendre la fin des goroutines...
}
