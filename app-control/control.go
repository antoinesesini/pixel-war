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
var horlogeVectorielle = utils.HorlogeVectorielle{}
var monEtatLocal utils.EtatLocal
var etatGlobal utils.EtatGlobal
var monBilan = 0
var nbEtatsAttendus = 0
var nbMessagesAttendus = 0
var N = 3
var maCouleur = utils.Blanc
var jeSuisInitiateur = false

var tabSC = make([]utils.ElementExclusionMutuelle, N)

var pNom = flag.String("n", "controle", "nom")
var monNom string

func main() {
	flag.Parse()
	monNom = *pNom + "-" + strconv.Itoa(os.Getpid())

	for _, e := range tabSC {
		e.Type = utils.Liberation
		e.Horloge = 0
	}

	monEtatLocal.NomSite = monNom
	horlogeVectorielle[monNom] = 0

	go lecture()
	for {
		time.Sleep(time.Duration(60) * time.Second)
	} // Pour attendre la fin des goroutines...
}
