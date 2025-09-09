package menu

import "fmt"

func MenuGeneral () {
	var selection int 
	fmt.Println("          TeKnologia          \n")
	fmt.Println("Options\n")
	fmt.Println("1. Base de donnees\n 2. Boutique\n 3. Entrainement\n 4. Quete\n")
	fmt.Scan(&selection)
	if selection == 1{

	}
	if selection == 2{
		
	}
	if selection == 3{
		
	}
	if selection == 4{
		
	}
	if selection > 4 || selection <= 0{
		fmt.Println("Option indisponible, veuillez choisir parmi les 4 propositions\n")
		MenuGeneral()
		
	}
}

func MenuDataBase (){
	var selection int
	fmt.Println("          BASE DE DONNEES          \n")
	fmt.Println("Options\n")
	fmt.Println("1. Fiche du personnage\n 2. Inventaire\n 3. Retour")
	fmt.Scan(&selection)

	if selection == 1{

	}
	if selection == 2{
		
	}
	if selection == 3{
		MenuGeneral()
	}
	
	if selection > 3 || selection <= 0{
		fmt.Println("Option indisponible, veuillez choisir parmi les 3 propositions\n")
		MenuDataBase()
	}

}

func MenuShop (){
	var selection int
	fmt.Println("          BOUTIQUE          \n")
	fmt.Println("Options\n")
	fmt.Println("1. Potions\n 2. Armes\n 3. Armures\n 4. Retour\n")
	fmt.Scan(&selection)
	if selection == 1{

	}
	if selection == 2{
		
	}
	if selection == 3{
		
	}
	if selection == 4{
		MenuGeneral()
		
	}
	if selection > 4 || selection <= 0{
		fmt.Println("Option indisponible, veuillez choisir parmi les 4 propositions\n")
		MenuGeneral()
		
	}

}
