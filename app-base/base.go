package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
	"utils"
)

func frontend(msg utils.MessagePixel) {
	SectionCritique := true
	//recupPixelAChangerSansLeChangerNiEnLocalNiEnPrevenantLesAutres
	MsgPixelTmp := utils.MessagePixelToString(msg)

	//Keep the pixel
	PixelTmp := utils.MessagePixel{x, y, r, g, b}
	//demandeSC = envoi d'un message demandeSC à l'app de control$
	utils.MessageExclusionMutuelle {Type: 0,Estampille: (numSite, Horloge) } SecCrit
	MessageSC( SecCrit )

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
func lecture() {
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
			changerPixel(messagePixel)
			mutex.Unlock()
		}
		rcvmsg = ""
	}
}

func changerPixel(messagePixel utils.MessagePixel) {
	utils.DisplayError(monNom, "changerPixel", "Et là bim on change le pixel")
}

var mutex = &sync.Mutex{}
var pNom = flag.String("n", "base", "nom")
var monNom string

func main() {

	flag.Parse()
	monNom = *pNom + "-" + strconv.Itoa(os.Getpid())

	//Création de 2 go routines qui s'exécutent en parallèle
	go sendperiodic()
	go lecture()
	//On décide de bloquer le programme principal
	for {
		time.Sleep(time.Duration(60) * time.Second)
	} // Pour attendre la fin des goroutines...
}
