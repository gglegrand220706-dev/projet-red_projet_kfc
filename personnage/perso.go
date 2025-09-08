package main

import "fmt"

type Character struct {

Nom string 
Classe string 
Niveau int 
Viemax int 
Vieactuelle int 
Inventaire int   
}

func main() {
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
	fmt.Println("choisisir un pseudo : ")
	fmt.Scanln(&u.Classe) 
	u.DisplayClass()

}