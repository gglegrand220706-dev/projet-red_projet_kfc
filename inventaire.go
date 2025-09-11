package personnage

import "fmt"

type Potions struct {
	name string
	Nb int
}

var Potions01 = Potions{"Potion de vie", 0}
var Potions02 = Potions{"Potions de dégas", 0}

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
	fmt.Print("-", Potions01.name, " x ", Potions01.Nb, "\n","-", Potions02.name, " x ", Potions02.Nb, "\n")
	if Arme01.Possède == true {
		fmt.Print("-", Arme01.Name)
	}
	if Armures01.possède == true {
		fmt.Print("-", Armures01.Name)
	}
}