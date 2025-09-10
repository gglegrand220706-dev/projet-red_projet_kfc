package menu

import (
	"fmt"
	"personnage"
)

func MenuGeneral() {
	var selection int
	fmt.Print("\033[H\033[2J")
	fmt.Print("          TeKnologia          \n")
	fmt.Print("Options\n")
	fmt.Print("1. Base de donnees\n 2. Boutique\n 3. Entrainement\n 4. Quete\n")
	fmt.Scan(&selection)
	if selection == 1 {
		var joueur personnage.Character
		fmt.Print("\033[H\033[2J")
		personnage.DisplayInfo(joueur)
	}
	if selection == 2 {
		fmt.Print("\033[H\033[2J")
	}
	if selection == 3 {
		fmt.Print("\033[H\033[2J")
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
	fmt.Print("1. Fiche du personnage\n 2. Inventaire\n 3. Retour")
	fmt.Scan(&selection)

	if selection == 1 {
		fmt.Print("\033[H\033[2J")
	}
	if selection == 2 {
		fmt.Print("\033[H\033[2J")
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
	}
	if selection == 2 {
		fmt.Print("\033[H\033[2J")
	}
	if selection == 3 {
		fmt.Print("\033[H\033[2J")
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
