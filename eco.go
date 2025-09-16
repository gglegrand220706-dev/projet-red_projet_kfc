package personnage

import (
	"fmt"
)

func ShopPotion() {
	var ChoixPotion int
	fmt.Print("\nVous avez ", Joueur.Bourse, " TeKno-dollar\n")
	fmt.Print("\n1. potion de soins\n")
	fmt.Print("2. potion de poison\n")
	fmt.Print("3. Retour\n")
	var PotionDispo = []*Potions{&PotionsSoins, &PotionsDamage}
	fmt.Scan(&ChoixPotion)	
	if PotionDispo[ChoixPotion - 1].Prix <= Joueur.Bourse && PotionDispo[ChoixPotion - 1].Nb < 5{
		Joueur.Bourse -= PotionDispo[ChoixPotion - 1].Prix
		PotionDispo[ChoixPotion - 1].Nb ++
		fmt.Print("Vous avez acheter ", PotionDispo[ChoixPotion -1].Name, " il vous reste donc ", Joueur.Bourse, " TeKno-Dollars")
		if ChoixPotion == 3 {
			MenuShop()
		}
		if ChoixPotion < 0 || ChoixPotion > 4 {
			fmt.Print("se n'est pas une proposition disponible")
			ShopPotion()
		}
	}
}

func ShopArmes() {
	var ChoixArmes int
	fmt.Print("\nVous avez ", Joueur.Bourse, " TeKno-Dollar\n")
	fmt.Print("1. ", Arme01.Name, " ", Arme01.Prix, " TeKno-Dollar", "\n2. ", Arme02.Name, " ", Arme02.Prix, " TeKno-Dollar", "\n3. ", Arme03.Name, " ", Arme03.Prix, " TeKno-Dollar\n")
	var ArmesDispo = []*Armes{&Arme01, &Arme02, &Arme03 }
	fmt.Scan(&ChoixArmes)
	if ChoixArmes < 0 || ChoixArmes > len(ArmesDispo) {
		fmt.Print("se n'est pas une option disponible")
		Continuer()
		ShopArmes()
	}
	if ArmesDispo[ChoixArmes -1].Prix <= Joueur.Bourse && !ArmesDispo[ChoixArmes-1].Possede{
		Joueur.Bourse -= ArmesDispo[ChoixArmes -1].Prix
		ArmesDispo[ChoixArmes -1].Possede = true
		fmt.Print("vous avez acheté ", ArmesDispo[ChoixArmes -1].Name, "\n")
		fmt.Print("il vous reste : ", Joueur.Bourse, " TeKno-Dollar\n")
		Continuer()
		MenuShop()
	}else {
		fmt.Print("Vous ne pouvez pas acheter :P\n")
		Continuer()
		ShopArmes()
	}
}

func ShopArmures() {
	fmt.Print("Vous avez ", Joueur.Bourse, " TeKno-Dollar\n")
	var ChoixArmures int
	fmt.Print("1. ", Armure01.Name, " ", Armure01.Prix, "\n")
	var ArmuresDispo = []*Armures{&Armure01}
	fmt.Scan(&ChoixArmures)
	if ChoixArmures < 0 || ChoixArmures > len(ArmuresDispo) {
		fmt.Print("se n'est pas un choix disponible")
		Continuer()
		ShopArmures()
	}
	if Joueur.Bourse >= ArmuresDispo[ChoixArmures -1].Prix && !ArmuresDispo[ChoixArmures -1].Possede {
		Joueur.Bourse -= ArmuresDispo[ChoixArmures -1].Prix
		ArmuresDispo[ChoixArmures -1].Possede = true
		fmt.Print("Vous avez acheter ", ArmuresDispo[ChoixArmures -1].Name, " il vous reste ", Joueur.Bourse, " TeKno-Dollars")
		Continuer()
		MenuShop()
	}
}
