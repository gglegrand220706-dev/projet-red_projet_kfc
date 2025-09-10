package personnage

import "fmt"

type Character struct {

Nom string 
Classe string 
vieactuelle int 
Niveau int 
}

func CharacterCreation() {
	var joueur Character
	joueur.RecupInfoName()
	var classe Character
	classe.RecupInfoClass()
}

func (u Character) DisplayName() {
	fmt.Printf("votre pseudo est donc: " + u.Nom)
}



func (u *Character) RecupInfoName() {
	fmt.Println("choisisir un pseudo : ")
	fmt.Scanln(&u.Nom) 
	u.DisplayName()

}

func (u Character) DisplayClass() {
	fmt.Printf("votre classe est donc: " + u.Classe)
}

func (u *Character) RecupInfoClass() {
	var ClasseSelection int
	fmt.Println("\n choisisir un classe :\n 1.Kryptonien \n 2.Bat family \n 3.Hulk \n 4.Wakanda")
	fmt.Scanln(&ClasseSelection)
	if ClasseSelection == 1 {
		u.Classe = "Kryptonien"
	}
	if ClasseSelection == 2 {
		u.Classe = "Bat family"
	}
	if ClasseSelection == 3 {
		u.Classe = "Hulk"
	}
	if ClasseSelection == 4 {
		u.Classe = "Wakanda"
	}
	if ClasseSelection > 4 || ClasseSelection < 0 {
		fmt.Println("ce n'est pas une option disponible")
		u.RecupInfoClass()
	}
	u.DisplayClass()
}