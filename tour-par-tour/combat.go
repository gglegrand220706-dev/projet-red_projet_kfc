package tourpartour

import (
	"fmt"
	"personnage"
)

func AtackSysteme() {
	var joueur *personnage.Character
	fmt.Print("que voulez-vous faire :\n")
	fmt.Print("1. ", joueur.Capacity[1], "\n2. ", joueur.Capacity[2])
}