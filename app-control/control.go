package main

import (
	"flag"
	"os"
	"strconv"
	"sync"
	"utils"
)

// Définition des variables
var mutex = &sync.Mutex{}
var K = 0
var maCouleur = utils.Blanc
var jeSuisInitiateur = false
var monEtatLocal utils.EtatGlobal
var monBilan int

var pNom = flag.String("n", "controle", "nom")
var monNom string

func main() {
	flag.Parse()
	monNom = *pNom + "-" + strconv.Itoa(os.Getpid())

	//go lecture()
	//go ecriture()
	//for {
	//	time.Sleep(time.Duration(60) * time.Second)
	//} // Pour attendre la fin des goroutines...
}
