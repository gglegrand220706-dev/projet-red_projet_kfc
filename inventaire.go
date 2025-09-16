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
var AllPotions = []*Potions{&PotionsSoins, &PotionsDamage}

type Armes struct {
	Name string
	Possede bool
	InStrore bool
	Multiplicateur float32
	Prix int
	ObjectsCraft []Objects
	RequiredQuantity []int
}


var EpeeTraining = Armes{"épées d'entraînements", false, true, 1.15, 30, []Objects{}, []int{}}
var GantThanos = Armes{"gant de Thanos", false, false, 10, 1000, []Objects{ScrapMetal, InfinityStoneMind, InfinityStoneReality, InfinityStoneTime, InfinityStonePouvoir, InfinityStoneSpace, InfinityStoneSoule}, []int{5, 1, 1, 1, 1, 1, 1}}
var MarteauThor = Armes{"marteau de Thor", false, true, 4, 200, []Objects{}, []int{}}
var StormBreaker = Armes{"Hache de Thor", false, false, 5, 400, []Objects{ScrapMetal, GrootsBranches}, []int{5, 5}}
var AllWeapons = []*Armes{&EpeeTraining, &GantThanos, &MarteauThor, &StormBreaker}

type Armures struct {
	Name string
	Possede bool
	InStrore bool
	Protect float32
	Prix int
}

var ArmureTraining = Armures{"armure d'entraînment", false, true, 0.15, 15}
var AllArmures = []*Armures{&ArmureTraining}

type Objects struct {
	Name string
	Nb int
	DropRate int
	MaxAmount int
}

var ScrapMetal = Objects{"scrape de Uruu", 10, 75, 10}
var InfinityStoneMind = Objects{"Pierre de l'esprit", 1, 15, 1}
var InfinityStoneReality = Objects{"Pierre de la réalitée", 1, 15, 1}
var InfinityStoneTime = Objects{"pierre du temps", 1, 15, 1}
var InfinityStonePouvoir = Objects{"Pierre du pouvoir", 1, 15, 1}
var InfinityStoneSpace = Objects{"Pierre de l'espace", 1, 15, 1}
var InfinityStoneSoule = Objects{"Pierre de l'âme", 1, 15, 1}
var GrootsBranches = Objects{"Branche de groot", 10, 50, 10}
var AllObjects = []*Objects{&ScrapMetal, &InfinityStoneMind, &InfinityStoneReality, &InfinityStoneTime, &InfinityStonePouvoir, &InfinityStoneSpace, &InfinityStoneSoule, &GrootsBranches}

func DisplayInventoryObjects() {
	fmt.Print("               invetaire: Objets             \n")
	var Choice int
	for _, Obj := range AllObjects {
		if Obj.Nb > 0 {
			fmt.Print("- ", Obj.Name, " x ", Obj.Nb, "\n")
		}		
	}
	fmt.Print("\n")
	for i, Obj := range AllPotions {
		fmt.Print(i +1, " ", Obj.Name, " x ", Obj.Nb, "\n")
	}	
	fmt.Print("\n3. retour\n")
	fmt.Print("Que voulez vous faire ?\n")
	fmt.Scan(&Choice)
	switch Choice {
	case 1:
		PotionsVie()
	case 2 :
		fmt.Print("Ceci n'est pas une bonne idées, gardez la pour d'autres ocasions")
		DisplayInventoryObjects()
	case 3 :
		DisplayInventory()
	default :
		fmt.Print("Se n'est pas une option disponible")
		DisplayInventoryObjects()
	}
}

func DisplayInventoryWeapons() {
    fmt.Print("               inventaire: Armes             \n")
    var Choice int
    var index int = 1
    var mapping []int 
    for i, Obj := range AllWeapons {
        if Obj.Possede {
            if Obj.Name == Joueur.EquipedWeapon.Name {
                fmt.Printf("%d. %s (équipée)\n", index, Obj.Name)
            } else {
                fmt.Printf("%d. %s\n", index, Obj.Name)
            }
            mapping = append(mapping, i)
            index++
        }
    }
    fmt.Printf("\n%d. retour\n", index)
    fmt.Print("Que voulez-vous faire ?\n")
    fmt.Scan(&Choice)
    if Choice == index {
        DisplayInventory()
        return
    }
    if Choice <= 0 || Choice > index {
        fmt.Println("Ce n'est pas une option disponible")
        DisplayInventoryWeapons()
        return
    }
    Joueur.EquipedWeapon = *AllWeapons[mapping[Choice-1]]
    fmt.Println("Vous avez équipé", Joueur.EquipedWeapon.Name)
}

func DisplayInventoryArmures() {
    fmt.Print("               inventaire: Armures             \n")
    var Choice int
    var index int = 1
    var mapping []int 
    for i, Obj := range AllArmures {
        if Obj.Possede {
            if Obj.Name == Joueur.EquipedArmure.Name {
                fmt.Printf("%d. %s (équipée)\n", index, Obj.Name)
            } else {
                fmt.Printf("%d. %s\n", index, Obj.Name)
            }
            mapping = append(mapping, i)
            index++
        }
    }
    fmt.Printf("\n%d. retour\n", index)
    fmt.Print("Que voulez-vous faire ?\n")
    fmt.Scan(&Choice)
    if Choice == index {
        DisplayInventory()
        return
    }
    if Choice <= 0 || Choice > index {
        fmt.Println("Ce n'est pas une option disponible")
        DisplayInventoryArmures()
        return
    }
    Joueur.EquipedArmure = *AllArmures[mapping[Choice-1]]
    fmt.Println("Vous avez équipé", Joueur.EquipedArmure.Name)
}

