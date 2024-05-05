package main

import (
	"flag"
	"os"
	"strconv"
	"sync"
	"time"
	"utils"
)

// Définition des variables
var mutex = &sync.Mutex{}
var H = 0
var maCouleur = utils.Blanc
var jeSuisInitiateur = false
var monEtatLocal utils.EtatGlobal
var monBilan int
var N = 3
var tabSC = []utils.MessageExclusionMutuelle{}

var pNom = flag.String("n", "controle", "nom")
var monNom string

func main() {
	flag.Parse()
	monNom = *pNom + "-" + strconv.Itoa(os.Getpid())

	/*for _, e := range tabSC {
		e.Type = utils.Liberation
		e.Horloge = 0
	}
	*/

	go lecture()
	for {
		time.Sleep(time.Duration(60) * time.Second)
	} // Pour attendre la fin des goroutines...
}
