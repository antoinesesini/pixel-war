package main

import "utils"

// Message commencant par un B, APP Base -> APP CONTROL
// / A REPRENDRE
func traiterMessageDemandeSC(rcvmsg string) {
	H++
	demande := utils.StringToMessageElementExclusionMutuelle(rcvmsg)
	tabSC[Site] = utils.MessageExclusionMutuelle{
		Type:       demande,
		Estampille: utils.Estampille{Site: Site, Horloge: H},
	}
	/*
		if utils.QuestionEntreeSC(Site, tabSC) {
			envoyerMessageSCBase(tabSC[Site].Type)
		}
	*/
}

// Message commencant par un D
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

// Message commencant par un C
func traiterMessageRequete(rcvmsg string) {
	demande := utils.StringToMessageExclusionMutuelle(rcvmsg)
	H = max(demande.Estampille.Horloge, H) + 1
	tabSC[demande.Estampille.Site] = utils.MessageExclusionMutuelle{
		Type:       utils.Requete,
		Estampille: demande.Estampille,
	}
	envoyerMessageSCControle(demande)

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
