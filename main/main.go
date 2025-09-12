package main

import (
	"personnage"
	"personnage/menu"
)

func main() {	
	personnage.CharacterCreation()
	menu.MenuGeneral()
	personnage.LevelUp()
}
