package main

import (
	"fmt"
	"strconv"
	"utils"
)

// Pour l'instant, boucle sur l'entrée standard, lit et communique le résultat à la routine d'écriture
func lecture() {
	var rcvmsg string
	for {
		fmt.Scanln(&rcvmsg)
		if rcvmsg == "" {
			utils.DisplayWarning(monNom, "lecture", "Message vide reçu")
			continue
		}
		mutex.Lock()
		// On traite uniquement les messages qui ne commencent pas par un 'A'
		if rcvmsg[0] != uint8('A') {

			// Demande de sauvegarde
			if rcvmsg == "sauvegarde" {
				traiterDebutSauvegarde()

				// Traitement des messages de contrôle
			} else if utils.TrouverValeur(rcvmsg, "horloge") != "" {
				traiterMessageControle(rcvmsg)
			} else if utils.TrouverValeur(rcvmsg, "etat") != "" {
				traiterMessageEtat(rcvmsg)
			} else {
				traiterMessagePixel(rcvmsg)
			}
		}
		mutex.Unlock()
	}
}

// TRAITEMENT DES CONTRÔLES NORMAUX : on extrait le pixel que l'on exploite dans l'app-base et on fait suivre l'information
// et tout cela avec les bonnes informations mises à jour dans le message : horloge, couleur
func traiterMessageControle(rcvmsg string) {
	message := utils.StringToMessage(rcvmsg)
	monBilan--

	// On traite le message uniquement s'il ne vient pas de nous
	if message.Nom == monNom {
		return
	}

	utils.DisplayWarning(monNom, "Controle", "Message de contrôle reçu : "+rcvmsg)

	// Extraction de la partie pixel
	messagePixel := message.Pixel

	// Recalage de l'horloge locale et mise à jour de sa valeur dans le message également
	H = utils.Recaler(H, message.Horloge)
	message.Horloge = H

	// Mise à jour de l'horloge vectorielle locale et mise à jour de sa valeur dans le message également
	horlogeVectorielle = utils.MajHorlogeVectorielle(monNom, horlogeVectorielle, message.Vectorielle)
	message.Vectorielle = horlogeVectorielle

	// Première fois qu'on reçoit l'ordre de transmettre sa sauvegarde
	if message.Couleur == utils.Jaune && maCouleur == utils.Blanc {
		maCouleur = utils.Jaune
		utils.DisplayError(monNom, "Controle", "Passage en jaune")

		utils.DisplayError(monNom, "Controle", "EtatLocal : "+utils.EtatLocalToString(monEtatLocal))
		go envoyerMessageEtat(monEtatLocal)
	}

	message.Couleur = maCouleur

	// On met à jour l'état local
	monEtatLocal = utils.MajEtatLocal(monEtatLocal, messagePixel)
	monEtatLocal.Vectorielle = horlogeVectorielle

	go envoyerMessageControle(message) // Pour la prochaine app de contrôle de l'anneau
	monBilan++
	go envoyerMessageBase(messagePixel) // Pour l'app de base
	utils.DisplayInfo(monNom, "Controle", "monBilanActuel = "+strconv.Itoa(int(monBilan)))
}

func traiterMessageEtat(rcvmsg string) {

	if !jeSuisInitiateur {
		utils.DisplayError(monNom, "Etat", "Transfert message etat : "+rcvmsg)
		go envoyerMessage(rcvmsg)
		return
	}

	utils.DisplayError(monNom, "Etat", "MessageEtat recu")
	etatLocal := utils.StringToEtatLocal(rcvmsg)

	// On ajoute l'état local reçu à la sauvegarde générale
	etatGlobal = append(etatGlobal, utils.CopyEtatLocal(etatLocal))

	nbEtatsAttendus--

	utils.DisplayError(monNom, "Etat", "nbEtatsAttendus="+strconv.Itoa(nbEtatsAttendus)+" ; nbMessagesAttendus="+strconv.Itoa(nbMessagesAttendus))
	if nbEtatsAttendus == 0 {
		finSauvegarde()
	}
}

func traiterMessagePixel(rcvmsg string) {
	utils.DisplayWarning(monNom, "lecture", "Message pixel reçu : "+rcvmsg)

	messagePixel := utils.StringToMessagePixel(rcvmsg)
	H++

	horlogeVectorielle[monNom]++

	// Mise à jour de l'état local
	monEtatLocal = utils.MajEtatLocal(monEtatLocal, messagePixel)
	monEtatLocal.Vectorielle = horlogeVectorielle

	message := utils.Message{messagePixel, H, horlogeVectorielle, monNom, maCouleur}
	go envoyerMessageControle(message)
	monBilan++
	utils.DisplayInfo(monNom, "Pixel", "monBilanActuel = "+strconv.Itoa(int(monBilan)))
}

func traiterDebutSauvegarde() {
	utils.DisplayError(monNom, "Debut", "debut de la sauvegarde")
	maCouleur = utils.Jaune
	jeSuisInitiateur = true
	nbEtatsAttendus = N - 1
	nbMessagesAttendus = monBilan

	utils.DisplayError(monNom, "Debut", "nbEtatsAttendus="+strconv.Itoa(nbEtatsAttendus)+" ; nbMessagesAttendus="+strconv.Itoa(nbMessagesAttendus))

	// On ajoute l'état local à la sauvegarde générale
	for _, mp := range monEtatLocal.ListMessagePixel {
		utils.DisplayError(monNom, "Debut", utils.MessagePixelToString(mp))
	}
	etatGlobal = append(etatGlobal, utils.CopyEtatLocal(monEtatLocal))
}

func finSauvegarde() {
	utils.DisplayError(monNom, "Fin", "Sauvegarde complétée")
	for _, etatLocal := range etatGlobal {
		utils.DisplayInfo(monNom, "Fin", utils.EtatLocalToString(etatLocal))
	}

	if utils.CoupureEstCoherente(etatGlobal) {
		utils.DisplayInfo(monNom, "Fin", "COUPURE COHÉRENTE !")
	} else {
		utils.DisplayInfo(monNom, "Fin", "Coupure non cohérente...")
	}
}
