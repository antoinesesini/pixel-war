package main

import (
	"flag"
	"fmt"
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
var tabSC []utils.MessageExclusionMutuelle

var pNom = flag.String("n", "controle", "nom")
var monNom string
var Site int

func main() {
	flag.Parse()
	Site = utils.InitialisationNumSite(*pNom) - 1

	// DEBUG

	fmt.Println("Numéro site = ", Site)
	msg := utils.MessageExclusionMutuelle{Type: utils.Liberation, Estampille: utils.Estampille{Site: Site, Horloge: 1}}
	fmt.Println("Message vaut : ", utils.MessageExclusionMutuelleToString(msg))
	msgSC := utils.Liberation
	fmt.Println(utils.MessageTypeSC(msgSC))
	test := "B/=typeSC=1/=estampilleSite=18/=estampilleHorloge=1"
	test_conv := utils.StringToMessageExclusionMutuelle(test)
	test = "B"
	test += utils.MessageExclusionMutuelleToString(test_conv)
	fmt.Println("COnversion en msg élém exclu mutuelle puis en string vaut : ", test)
	test_sc := "/=typeSC=1"
	testSC := utils.StringToMessageTypeSC(test_sc)
	fmt.Println("Conversion en type SC puis en string vaut : ", utils.MessageTypeSC(testSC))
	tabSC = make([]utils.MessageExclusionMutuelle, 3)
	tabSC[0] = utils.MessageExclusionMutuelle{Type: utils.Requete, Estampille: utils.Estampille{Site: Site, Horloge: 1}}
	tabSC[1] = utils.MessageExclusionMutuelle{Type: utils.Requete, Estampille: utils.Estampille{Site: 2, Horloge: 0}}
	tabSC[2] = utils.MessageExclusionMutuelle{Type: utils.Requete, Estampille: utils.Estampille{Site: 3, Horloge: 5}}
	for x := range tabSC {
		fmt.Println(" TabSC [", x, "] vaut = ", utils.MessageExclusionMutuelleToString(tabSC[x]))
	}
	if utils.QuestionEntreeSC(Site, tabSC) {
		fmt.Println("Je suis bien en Section critique, la fonction marche")
	} else {
		fmt.Println("Je suis pas en section critique, la fonction ne marche pas")
	}

	// Récupérer l'heure actuelle
	now := time.Now()

	// Convertir l'heure actuelle en chaîne de caractères
	timeString := now.Format("2006-01-02/15:04:05")

	// Afficher la chaîne de caractères
	fmt.Println(timeString)
	// FIN DEBUG

	/*
		monNom = *pNom + "-" + strconv.Itoa(os.Getpid())
			i := 1
			for _, e := range tabSC {
				e.Type = utils.Liberation
				e.Estampille = utils.Estampille{Site: i, Horloge: 0}
			}

			go lecture()
			for {
				time.Sleep(time.Duration(60) * time.Second)
			} // Pour attendre la fin des goroutines...
	*/
}
