package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"utils"
)

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
			utils.DisplayError("app-de-base", "lecture", "Réception de : "+rcvmsg[1:])
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

	colorWheel, _, err := ebitenutil.NewImageFromFile("color_wheel.png")
	if err != nil {
		panic(err)
	}

	game := &utils.Game{
		Matrix:        matrix,
		ColorWheel:    colorWheel,
		SelectedColor: color.RGBA{R: 0, G: 0, B: 0, A: 0xFF},
	}

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " ")
			if len(parts) == 5 {
				x, err := strconv.Atoi(parts[0])
				if err != nil {
					continue
				}
				y, err := strconv.Atoi(parts[1])
				if err != nil {
					continue
				}
				cr, err := strconv.Atoi(parts[2])
				if err != nil {
					continue
				}
				cg, err := strconv.Atoi(parts[3])
				if err != nil {
					continue
				}
				cb, err := strconv.Atoi(parts[4])
				if err != nil {
					continue
				}
				game.UpdateMatrix(x, y, uint8(cr), uint8(cg), uint8(cb))
				fmt.Printf("Updated pixel at (%d, %d) to (%d, %d, %d)\n", x, y, cr, cg, cb)
			}
		}
	}()

	ebiten.RunGame(game)

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Pixel-War")

	//Création de 2 go routines qui s'exécutent en parallèle
	go sendperiodic()
	go lecture(game)
	//On décide de bloquer le programme principal
	for {
		time.Sleep(time.Duration(60) * time.Second)
	} // Pour attendre la fin des goroutines...
}
