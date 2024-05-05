package main

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
	"utils"
)

// Pour l'instant, boucle sur l'entrée standard, lit et communique le résultat à la routine d'écriture
func lecture() {
	var rcvmsg string
	for {
		fmt.Scanln(&rcvmsg)
		mutex.Lock()
		// On traite uniquement les messages qui ne commencent pas par un 'A'
		if rcvmsg[0] != uint8('A') {
			//TRAITEMENT DES MESSAGES DE CONTRÔLE
			if utils.TrouverValeur(rcvmsg, "horloge") != "" {
				if utils.TrouverValeur(rcvmsg, "prepost") == "true" {
					traiterMessagePrepost(rcvmsg)
				} else {
					//L'affichage sur stderr se fait dans le traitement pour ce type de message
					traiterMessageControle(rcvmsg)
				}
			} else if utils.TrouverValeur(rcvmsg, "etat") != "" {
				traiterMessageEtat(rcvmsg)
			} else {
				utils.DisplayWarning(monNom, "lecture", "Message pixel reçu : "+rcvmsg)
				traiterMessagePixel(rcvmsg)
			}
		}
		mutex.Unlock()
	}
}

// TO DO implémenter le traitement B

// TRAITEMENT DES CONTRÔLES NORMAUX : on extrait le pixel que l'on exploite dans l'app-base et on fait suivre l'information
// et tout cela avec les bonnes informations mises à jour dans le message : horloge, couleur
func traiterMessageControle(rcvmsg string) {
	message := utils.StringToMessage(rcvmsg)

	if message.Nom != monNom { // On traite le message uniquement s'il ne vient pas de nous
		utils.DisplayWarning(monNom, "main", "Message de contrôle reçu : "+rcvmsg)
		//Extraction de la partie pixel
		messagePixel := message.Pixel
		//Recalage de l'horloge locale et mise à jour de sa valeur dans le message également
		H = utils.Recaler(H, message.Horloge)
		message.Horloge = H

		//ATTENTION ICI, METTRE À JOUR L'ÉTAT GLOBAL AVANT D'ENVOYER QUOI QUE CE SOIT

		//Avertissement d'une coupure demandée et actions en conséquence
		if message.Couleur == utils.Jaune && maCouleur == utils.Blanc {
			maCouleur = utils.Jaune
			messageEtat := utils.MessageEtat{list.List(monEtatLocal), monBilan}
			go envoyerMessage(utils.MessageEtatToString(messageEtat))
			//Réception d'un message prépost pas encore marqué comme prépost
		} else if message.Couleur == utils.Blanc && maCouleur == utils.Jaune {
			if jeSuisInitiateur {
				// Ajouter le message reçu à la sauvegarde générale
			} else {
				messagePrepost := message
				messagePrepost.Prepost = true
				go envoyerMessageControle(messagePrepost)
			}
		}

		message.Couleur = maCouleur
		go envoyerMessageControle(message)  // Pour la prochaine app de contrôle de l'anneau
		go envoyerMessageBase(messagePixel) // Pour l'app de base
	}
}

func traiterMessagePrepost(rcvmsg string) {
	if jeSuisInitiateur {
		//message := utils.StringToMessage(rcvmsg)
		// Traiter l'ajout du message à l'état de sauvegarde
	} else {
		go envoyerMessage(rcvmsg) // On fait suivre le message sur l'anneau
	}
}

func traiterMessageEtat(rcvmsg string) {
	if jeSuisInitiateur {
		// Traiter l'ajout de l'état à la sauvegarde générale
	} else {
		go envoyerMessage(rcvmsg)
	}
}

func traiterMessagePixel(rcvmsg string) {
	messagePixel := utils.StringToMessagePixel(rcvmsg)
	H++
	//ATTENTION ICI, METTRE À JOUR L'ÉTAT GLOBAL
	message := utils.Message{messagePixel, H, monNom, maCouleur, false}
	go envoyerMessageControle(message)
}

func traiterMessageDemandeSC(rcvmsg string) {
	dem_str := string(rcvmsg[1])
	dem, _ := strconv.Atoi(dem_str)
	site := 1
	demande := utils.MessageExclusionMutuelle{Type: utils.TypeSC(dem), Estampille: utils.Estampille{site, H}}
	tabSC = append(tabSC, demande)
	sort.Sort(utils.MessageExclusionMutuelleSlice(tabSC))
	estampille := utils.Estampille{Site: 1, Horloge: H}
	envoyerMessageDemandeSC(utils.TypeSC(dem), estampille)
}

func traiterMessageFinSC(rcvmsg string) {
	site_str := string(rcvmsg[2])
	site, _ := strconv.Atoi(site_str)
	horloge_str := string(rcvmsg[3])
	horloge, _ := strconv.Atoi(horloge_str)
	tabSC = utils.SupprimerMessageExclusionMutuelle(tabSC, site, horloge)
}

func traiterMessageRequete(rcvmsg string) {
	dem_str := string(rcvmsg[1])
	dem, _ := strconv.Atoi(dem_str)
	site_str := string(rcvmsg[2])
	site, _ := strconv.Atoi(site_str)
	// Variable Site pas encore définis
	if site == Site {
		return
	}
	demande := utils.MessageExclusionMutuelle{Type: utils.TypeSC(dem), Estampille: utils.Estampille{site, H}}
	tabSC = append(tabSC, demande)
	sort.Sort(utils.MessageExclusionMutuelleSlice(tabSC))
	estampille := utils.Estampille{Site: 1, Horloge: H}
	envoyerMessageDemandeSC(utils.TypeSC(dem), estampille)
}

func traiterMessageLiberation(rcvmsg string) {

}

func traiterMessageAccuse(rcvmsg string) {

}
