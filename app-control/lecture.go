package main

import (
	"container/list"
	"fmt"
	"utils"
)

// Pour l'instant, boucle sur l'entrée standard, lit et communique le résultat à la routine d'écriture
func lecture() {
	var rcvmsg string
	for {
		fmt.Scanln(&rcvmsg)
		mutex.Lock()
		// On traite uniquement les messages qui ne commencent pas par un 'A'
		if rcvmsg[0] != uint8('A') && rcvmsg[0] != uint8('B') && rcvmsg[0] != uint8('C') {
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
			// TO DO implémenter le traitement B et C
		}
		if rcvmsg[0] == 'B' {
			traiterMessageDemandeSC(rcvmsg)
		}
	}
	if rcvmsg[0] == 'C' {
		demande := utils.StringToMessageTypeSC(rcvmsg)
		switch demande {
		case utils.Requete:
			traiterMessageRequete(rcvmsg)
		case utils.Accuse:
			traiterMessageAccuse(rcvmsg)
		case utils.Liberation:
			traiterMessageLiberation(rcvmsg)
		default:
		}
	}
	mutex.Unlock()
}

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

//// PARTIE EXCLUSION MUTUELLE

// Message commencant par un B, APP Base -> APP CONTROL. Traite aussi bien requete que libération
func traiterMessageDemandeSC(rcvmsg string) {
	demande := utils.StringToMessageTypeSC(rcvmsg)
	H++
	tabSC[Site] = utils.MessageExclusionMutuelle{
		Type:       demande,
		Estampille: utils.Estampille{Site: Site, Horloge: H},
	}
	MessageSC := utils.MessageExclusionMutuelle{Type: demande, Estampille: utils.Estampille{Site: Site, Horloge: H}}
	envoyerMessageSCControle(MessageSC)
}

// PROBABLEMENT A SUPPRIMER
/*
// Message commencant par un C, APP CONTROLE -> APP CONTROLE
func traiterMessageFinSC(rcvmsg string) {
	fin := utils.StringToMessageExclusionMutuelle(rcvmsg)
	H++
	tabSC[fin.Estampille.Site] = utils.MessageExclusionMutuelle{
		Type:       fin.Type,
		Estampille: fin.Estampille,
	}
	// A RAJOUTER SENS MESSAGE
	envoyerMessageSCControle(fin)

	if utils.QuestionEntreeSC(Site, tabSC) {
		envoyerMessageSCBase(tabSC[Site].Type)
	}
}

*/

// Message commencant par un C
func traiterMessageRequete(rcvmsg string) {
	demande := utils.StringToMessageExclusionMutuelle(rcvmsg)
	H = max(demande.Estampille.Horloge, H) + 1
	tabSC[demande.Estampille.Site] = utils.MessageExclusionMutuelle{
		Type:       utils.Requete,
		Estampille: demande.Estampille,
	}

	Accuse := utils.MessageExclusionMutuelle{Type: utils.Liberation, Estampille: utils.Estampille{Site: Site, Horloge: H}}
	envoyerMessageSCControle(Accuse)

	if utils.QuestionEntreeSC(Site, tabSC) {
		envoyerMessageSCBase(tabSC[Site].Type)
	}
}

// Message commencant par un C

func traiterMessageLiberation(rcvmsg string) {
	liberation := utils.StringToMessageExclusionMutuelle(rcvmsg)
	H = max(liberation.Estampille.Horloge, H) + 1

	tabSC[liberation.Estampille.Site] = utils.MessageExclusionMutuelle{
		Type:       liberation.Type,
		Estampille: liberation.Estampille,
	}
	envoyerMessageSCControle(liberation)
	if utils.QuestionEntreeSC(Site, tabSC) {
		envoyerMessageSCBase(tabSC[Site].Type)
	}
}

// Message commencant par un C
func traiterMessageAccuse(rcvmsg string) {
	mess := utils.StringToMessageExclusionMutuelle(rcvmsg)
	H = max(mess.Estampille.Horloge, H) + 1
	if tabSC[mess.Estampille.Site].Type != utils.Requete {
		tabSC[mess.Estampille.Site] = utils.MessageExclusionMutuelle{
			Type:       mess.Type,
			Estampille: mess.Estampille,
		}
	}
	tabSC[mess.Estampille.Site] = utils.MessageExclusionMutuelle{Type: utils.Accuse, Estampille: utils.Estampille{Site: mess.Estampille.Site, Horloge: H}}
	if utils.QuestionEntreeSC(Site, tabSC) {
		envoyerMessageSCBase(tabSC[Site].Type)
	}
}
