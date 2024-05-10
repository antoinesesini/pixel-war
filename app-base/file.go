package main

import (
	"fmt"
	"utils"
)

func frontend(msg utils.MessagePixel) {
	SectionCritique := true
	//recupPixelAChangerSansLeChangerNiEnLocalNiEnPrevenantLesAutres
	// ------> paramètre msg

	//Keep the pixel
	PixelTmp := utils.MessagePixel{msg.PositionX, msg.PositionY, msg.Rouge, msg.Vert, msg.Bleu}

	//demandeSC = envoi d'un message demandeSC à l'app de control$
	SecCrit := utils.Requete
	envoyerMessageBaseSC(SecCrit)
	for SectionCritique == true {
		si je recois un message comme quoi la SectionCritique est disponible
		uPDATE PixelTmp Dans la matrice actuelle
		SecCrit == false
	}
	envoyer un message de liberation de la section critique
	envoyer le message de la modification


	//Wait few minutes

}
// Fonction qui permet d'envoyer un message concernant l'accès / libération de la section critique
func envoyerMessageBaseSC(Type utils.TypeSC) {
	fmt.Println("B",Type)
}
