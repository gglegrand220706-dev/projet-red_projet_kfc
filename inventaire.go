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
	Possede bool
	Multiplicateur float32
	Prix int
}


var Arme01 = Armes{"épées d'entraînements", false, 1.15, 30}
var Arme02 = Armes{"gant de Thanos", false, 10, 1000}
var Arme03 = Armes{"marteau de Thor", false, 4, 200}

type Armures struct {
	Name string
	Possede bool
	Protect float32
	Prix int
}

var Armure01 = Armures{"armure d'entraînment", false, 0.15, 15}

func DisplayInventory() {
	fmt.Print("             invetaire             \n")
	fmt.Print("-", Potions01.Name, " x ", Potions01.Nb, "\n","-", Potions02.Name, " x ", Potions02.Nb, "\n")
	if Arme01.Possede == true {
		fmt.Print("-", Arme01.Name, "\n")
	}
	if Arme02.Possede == true {
		fmt.Print("-", Arme02.Name, "\n")
	}
	if Arme03.Possede == true {
		fmt.Print("-", Arme03.Name, "\n")
	}
	if Armure01.Possede == true {
		fmt.Print("-", Armure01.Name, "\n")
	}
}
