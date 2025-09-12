package personnage

import (
	"fmt"
)

func ShopPotion() {
	var ChoixPotion int
	fmt.Print("Vous avez ", Joueur.Bourse, " tekno-dollar\n")
	fmt.Print("1. potion de soins\n")
	fmt.Print("2. potion de poison\n")
	fmt.Scan(&ChoixPotion)
	if ChoixPotion == 1 {
		if Potions01.Prix <= Joueur.Bourse && Potions01.Nb < 5{
			Joueur.Bourse -= Potions01.Prix
			Potions01.Nb++
			fmt.Print("vous avez acheté une potion de soin\n")
			fmt.Print("il vous reste : ", Joueur.Bourse, " tekno-dollar\n")
		}else{
			fmt.Print("Vous ne pouvez pas acheter :P\n")
			ShopPotion()
		}
	}
	if ChoixPotion == 2 {
		if Potions02.Prix <= Joueur.Bourse {
			Joueur.Bourse -= Potions02.Prix
			Potions02.Nb++
			fmt.Print("vous avez acheté une potion de poison\n")
			fmt.Print("il vous reste : ", Joueur.Bourse, " tekno-dollar\n")
		}else{
			fmt.Print("Vous ne pouvez pas acheter :P\n")
			ShopPotion()
		}
	}
}

func ShopArmes() {
	var ChoixArmes int
	fmt.Print("Vous avez ", Joueur.Bourse, " tekno-dollar\n")
	fmt.Print("1. ", Arme01.Name, " ", Arme01.Prix, " tekno-dollar", "\n 2. ", Arme02.Name, " ", Arme02.Prix, " tekno-dollar", "\n 3. ", Arme03.Name, " ", Arme03.Prix, " tekno-dollar\n")
	fmt.Scan(&ChoixArmes)
	if ChoixArmes == 1 {
		if Arme01.Prix <= Joueur.Bourse && !Arme01.Possede{
			Joueur.Bourse -= Arme01.Prix
			Arme01.Possede = true
			fmt.Print("vous avez acheté ", Arme01.Name, "\n")
			fmt.Print("il vous reste : ", Joueur.Bourse, " tekno-dollar\n")
		}else{
			fmt.Print("Vous ne pouvez pas acheter :P\n")
			ShopArmes()
		}
	}
	if ChoixArmes == 2 {
		if Arme02.Prix <= Joueur.Bourse && !Arme02.Possede{
			Joueur.Bourse -= Arme02.Prix
			Arme02.Possede = true
			fmt.Print("vous avez acheté ", Arme02.Name, "\n")
			fmt.Print("il vous reste : ", Joueur.Bourse, " tekno-dollar\n")
		}else{
			fmt.Print("Vous ne pouvez pas acheter :P\n")
			ShopArmes()
		}
	}
	if ChoixArmes == 3 {
		if Arme03.Prix <= Joueur.Bourse && !Arme03.Possede{
			Joueur.Bourse -= Arme03.Prix
			Arme03.Possede = true
			fmt.Print("vous avez acheté ", Arme03.Name, "\n")
			fmt.Print("il vous reste : ", Joueur.Bourse, " tekno-dollar\n")
		}else{
			fmt.Print("Vous ne pouvez pas acheter :P\n")
			ShopArmes()
		}

	}
}

func ShopArmures() {
	fmt.Print("Vous avez ", Joueur.Bourse, " tekno-dollar\n")
	var ChoixArmures int
	fmt.Print("1. ", Armure01.Name, " ", Armure01.Prix, "\n")
	fmt.Scan(&ChoixArmures)
	if ChoixArmures == 1 {
		if Armure01.Prix <= Joueur.Bourse && !Armure01.Possede{
			Joueur.Bourse -= Armure01.Prix
			Armure01.Possede = true
			fmt.Print("vous avez acheté ", Arme01.Name, "\n")
			fmt.Print("il vous reste : ", Joueur.Bourse, " tekno-dollar\n")

		}else{
			fmt.Print("Vous ne pouvez pas acheter :P\n")
			ShopArmures()
		}
	}

}

func BourseReward() {
	Joueur.Bourse += Adversery01.GivenMoney
}