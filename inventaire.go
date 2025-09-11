package personnage

import "fmt"

type Potions struct {
	Name string
	Nb int
	Prix int
}

var Potions01 = Potions{"Potion de vie", 0, 20}
var Potions02 = Potions{"Potions de dégas", 0, 20}

type Armes struct {
	Name string
	Possède bool
	Multiplicateur float32
}

var Arme01 = Armes{"épées d'entraînements", false, 0.15}

type Armures struct {
	Name string
	possède bool
	Protecte float32
}

var Armures01 = Armures{"armure d'entraînment", false, 0.15}

func DisplayInventory() {
	fmt.Print("             invetaire             \n")
	fmt.Print("-", Potions01.Name, " x ", Potions01.Nb, "\n","-", Potions02.Name, " x ", Potions02.Nb, "\n")
	if Arme01.Possède == true {
		fmt.Print("-", Arme01.Name)
	}
	if Armures01.possède == true {
		fmt.Print("-", Armures01.Name)
	}
}