package utils

import "container/list"

// Définition des types
const sepM = "/" //séparateur dans les messages
const sepP = "=" //séparateur ddans les paires clé/valeur

type Couleur bool

const (
	Blanc Couleur = false
	Jaune Couleur = true
)

type MessagePixel struct {
	PositionX int
	PositionY int
	Rouge     int
	Vert      int
	Bleu      int
}

type Message struct {
	Pixel   MessagePixel
	Horloge int
	Nom     string
	Couleur Couleur
	Prepost bool //false pour les messages normaux
}

type EtatGlobal list.List //Sous-entendu une liste de MessagePixel

type MessageEtat struct {
	EG    list.List
	Bilan int
}
