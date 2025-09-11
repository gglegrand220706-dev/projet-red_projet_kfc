package personnage

import (
	"fmt"
)

func Shop() {
	var ChoixPotion int
	fmt.Print("1. potion de soins\n")
	fmt.Print("2. potion de poison\n")
	fmt.Scan(&ChoixPotion)
	if ChoixPotion == 1 {
		if Potions01.Prix < Joueur.Bourse {
			Joueur.Bourse = Joueur.Bourse - Potions01.Prix
			Potions01.Nb ++
	  		fmt.Print("vous avez acheté une potion de soin\n")
	  		fmt.Print("il vous reste : ", Joueur.Bourse, " de d'or")
		}
	}
	if ChoixPotion == 2 {
		if Potions02.Prix < Joueur.Bourse {
			Joueur.Bourse = Joueur.Bourse - Potions02.Prix
			Potions02.Nb ++
	  		fmt.Print("vous avez acheté une potion de soin\n")
	  		fmt.Print("il vous reste : ", Joueur.Bourse, " de d'or")
		}
	}
}
