package menu

import (
	"fmt"
	"personnage"
)

func MenuGeneral() {
	var selection int
	fmt.Print("\033[H\033[2J")
	fmt.Print("          TeKnologia : Battle World         \n")
	fmt.Print("Options\n")
	fmt.Print("1. Base de donnees\n 2. Boutique\n 3. Entrainement\n 4. Quete\n")
	fmt.Scan(&selection)
	if selection == 1 {
		fmt.Print("\033[H\033[2J")
		MenuDataBase()
	}
	if selection == 2 {
		fmt.Print("\033[H\033[2J")
		MenuShop()
	}
	if selection == 3 {
		fmt.Print("\033[H\033[2J")
		personnage.Adversery01.Vieactuelle = personnage.Adversery01.Viemax
		DisplayAtack()
		for personnage.Joueur.Vieactuelle > 0 && personnage.Adversery01.Vieactuelle > 0 {
				personnage.AtackSystème()
				personnage.Response()
				DisplayAtack()
			}
			if personnage.Joueur.Vieactuelle <= 0 {
				IsDead()
			}
			if personnage.Adversery01.Vieactuelle <= 0{
				personnage.AddExp()
				RetourMenu()
				personnage.BourseReward() 
			}
	}
	if selection == 4 {
		fmt.Print("\033[H\033[2J")
	
	}
	if selection > 4 || selection <= 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Option indisponible, veuillez choisir parmi les 4 propositions\n")
		MenuGeneral()

	}
}

func MenuDataBase() {
	var selection int
	fmt.Print("\033[H\033[2J")
	fmt.Print("          BASE DE DONNEES          \n")
	fmt.Print("Options\n")
	fmt.Print("1. Fiche du personnage\n 2. Inventaire\n 3. Retour\n")
	fmt.Scan(&selection)

	if selection == 1 {
		personnage.DisplayInfo()
		RetourMenu()
	}
	if selection == 2 {
		fmt.Print("\033[H\033[2J")
		personnage.DisplayInventory()
		RetourMenu()
	}
	if selection == 3 {
		fmt.Print("\033[H\033[2J")
		MenuGeneral()
	}

	if selection > 3 || selection <= 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Option indisponible, veuillez choisir parmi les 3 propositions\n")
		MenuDataBase()
	}
}

func MenuShop() {
	var selection int
	fmt.Print("\033[H\033[2J")
	fmt.Print("          BOUTIQUE          \n")
	fmt.Print("Options\n")
	fmt.Print("1. Potions\n 2. Armes\n 3. Armures\n 4. Retour\n")
	fmt.Scan(&selection)
	if selection == 1 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("          BOUTIQUE          \n")
		fmt.Print("Quelle potion voulez vous acheter\n")
		personnage.ShopPotion()
		RetourMenu()
	}
	if selection == 2 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("          BOUTIQUE          \n")
		fmt.Print("Quelle arme voulez vous acheter\n")
		personnage.ShopArmes()
		RetourMenu()
	}
	if selection == 3 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("          BOUTIQUE          \n")
		fmt.Print("Quelle armure voulez vous acheter\n")
		personnage.ShopArmures() 
		RetourMenu()
	}
	if selection == 4 {
		fmt.Print("\033[H\033[2J")
		MenuGeneral()
	}
	if selection > 4 || selection <= 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Option indisponible, veuillez choisir parmi les 4 propositions\n")
		MenuGeneral()

	}
}

func RetourMenu() {
	var Retour int
	fmt.Print("\nappuyer sur 1 pour retourner au menue ")
	fmt.Scan(&Retour)
	if Retour == 1 {
		MenuGeneral()
	}

}

func DisplayAtack() {
	if personnage.Joueur.Vieactuelle > 0 && personnage.Adversery01.Vieactuelle > 0 {
		var Attaques = []string{personnage.Joueur.Attaques[0].Name, personnage.Joueur.Attaques[1].Name, personnage.Joueur.Attaques[2].Name, personnage.Joueur.Attaques[3].Name}
    	fmt.Println("\nQue voulez-vous faire :")
    	for i, cap := range Attaques {
			fmt.Printf("%d. %v\n", i+1, cap )
			i++	
    	}
	}	
}

func IsDead() {
		fmt.Print("Vous êtes mort sale noob")
		fmt.Print("vous récussitez avec ", personnage.Joueur.Viemax/2)
		personnage.Joueur.Vieactuelle = personnage.Joueur.Viemax/2
		RetourMenu()
}