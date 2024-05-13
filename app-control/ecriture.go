package main

import (
	"fmt"
	"utils"
)

// Envoi une chaine de caractères sur la sortie standard
func envoyerMessage(message string) {
	mutex.Lock()
	fmt.Println(message)
	mutex.Unlock()
}

func envoiSequentiel(message string) {
	fmt.Println(message)
}

// Envoi un type Message pour les applis de contrôle
func envoyerMessageControle(message utils.Message) {
	go envoyerMessage(utils.MessageToString(message))
}

// Envoi un type MessageEtat pour les applis de contrôle
func envoyerMessageEtat(messageEtat utils.MessageEtat) {
	go envoyerMessage(utils.MessageEtatToString(messageEtat))
}

// Envoi un type MessagePixel pour l'appli de base
func envoyerMessageBase(messagePixel utils.MessagePixel) {
	go envoyerMessage("A" + utils.MessagePixelToString(messagePixel))
}

// Envoi un type MessageSauvegarde pour l'appli de base
func envoyerMessageBaseSauvegarde(messageSauvegarde utils.MessageSauvegarde) {
	go envoyerMessage("A" + utils.MessageSauvegardeToString(messageSauvegarde))
}

/////////////////////
// Exclusion mutuelle
/////////////////////

func envoyerMessageSCControle(msgSC utils.MessageExclusionMutuelle) {
	msg := utils.MessageExclusionMutuelleToString(msgSC)
	go envoyerMessage(msg)
}

func envoyerMessageAccuse(msgAcc utils.MessageAccuse) {
	msg := utils.MessageAccuseToString(msgAcc)
	envoiSequentiel(msg)
}

func envoyerMessageSCBase(msgSC utils.TypeSC) {
	msg := "A" + utils.MessageTypeSCToString(msgSC)
	envoiSequentiel(msg)
}
