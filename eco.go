package personnage

import (
	"fmt"
)

func ShopPotion() {
	var ChoixPotion int
	fmt.Print("1. potion de soins\n")
	fmt.Print("2. potion de poison\n")
	fmt.Scan(&ChoixPotion)
	if ChoixPotion == 1 {
		if Potions01.Prix <=  Joueur.Bourse {
			Joueur.Bourse -= Potions01.Prix
			Potions01.Nb ++
	  		fmt.Print("vous avez acheté une potion de soin\n")
	  		fmt.Print("il vous reste : ", Joueur.Bourse, " Teknodollar\n")
		}
	}
	if ChoixPotion == 2 {
		if Potions02.Prix <= Joueur.Bourse {
			Joueur.Bourse -= Potions02.Prix
			Potions02.Nb ++
	  		fmt.Print("vous avez acheté une potion de poison\n")
	  		fmt.Print("il vous reste : ", Joueur.Bourse, " Teknodollar\n")
		}
	}
}

func ShopArmes() {
	var ChoixArmes int
	fmt.Print("1. Arme d'entraînement\n")
	fmt.Scan(&ChoixArmes)
	if ChoixArmes == 1 {
		if Arme01.Prix <= Joueur.Bourse {
			Joueur.Bourse -= Arme01.Prix
			Arme01.Possède = true
	  		fmt.Print("vous avez acheté une arme d'entraînement\n")
	  		fmt.Print("il vous reste : ", Joueur.Bourse, " Teknodollar\n")
		}
	}
}
