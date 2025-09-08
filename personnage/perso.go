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
	joueur.RecupInfo()
}

func (u Character) Display() {
	fmt.Printf("votre pseudo est donc: " + u.Nom)
}


func (u *Character) RecupInfo() {
	fmt.Println("choisisir un pseudo : ")
	fmt.Scanln(&u.Nom) 
	u.Display()

}
