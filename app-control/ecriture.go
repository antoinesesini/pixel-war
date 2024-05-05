package main

import (
	"fmt"
	"strconv"
	"utils"
)

// Envoi une chaine de caractères sur la sortie standard
func envoyerMessage(message string) {
	mutex.Lock()
	fmt.Println(message)
	mutex.Unlock()
}

// Envoi un type Message pour les applis de contrôle
func envoyerMessageControle(message utils.Message) {
	envoyerMessage(utils.MessageToString(message))
}

// Envoi un type MessagePixel pour l'appli de base
func envoyerMessageBase(messagePixel utils.MessagePixel) {
	envoyerMessage("A" + utils.MessagePixelToString(messagePixel))
}

func envoyerMessageDemandeSC(SC utils.TypeSC, estampille utils.Estampille) {
	msg := "C"
	msg = strconv.Itoa(int(SC))
	msg += strconv.Itoa(estampille.Site)
	msg += strconv.Itoa(estampille.Horloge)
	envoyerMessage(msg)
}

func envoyerMessageFinSC(SC utils.TypeSC, estampille utils.Estampille) {
	msg := "C"
	msg = strconv.Itoa(int(SC))
	msg += strconv.Itoa(estampille.Site)
	msg += strconv.Itoa(estampille.Horloge)
	envoyerMessage(msg)
}
