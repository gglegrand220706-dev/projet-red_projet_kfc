package personnage

import "fmt"

type Potions struct {
	Name string
	Nb int
	EffectLife int
	Prix int
}

var PotionsSoins = Potions{"Potion de vie", 0, 10, 20}
var PotionsDamage = Potions{"Potions de dégas", 0, -10, 20}
var AllPotions = []Potions{PotionsSoins, PotionsDamage}

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

type Objects struct {
	Name string
	Nb int
	DropRate int
}

var ScrapMetal = Objects{"scrape de metal", 0, 75}
var InfinityStoneMind = Objects{"pierre de l'esprit", 0, 15}

func DisplayInventory() {
	fmt.Print("             invetaire             \n")
	fmt.Print("-", PotionsSoins.Name, " x ", PotionsSoins.Nb, "\n","-", PotionsDamage.Name, " x ", PotionsDamage.Nb, "\n")
	if Arme01.Possede {
		fmt.Print("-", Arme01.Name, "\n")
	}
	if Arme02.Possede {
		fmt.Print("-", Arme02.Name, "\n")
	}
	if Arme03.Possede {
		fmt.Print("-", Arme03.Name, "\n")
	}
	if Armure01.Possede {
		fmt.Print("-", Armure01.Name, "\n")
	}
}
