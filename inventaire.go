package personnage

import "fmt"

type Potions struct {
	Name string
	Nb int
	EffectLife int
	Prix int
}

var PotionsSoins  = Potions{"\033[31mPotion de vie ❤️ \033[0m" ,  0, 10, 20}
var PotionsDamage = Potions{"\033[32mPotion de poison ☠️ \033[0m" ,  0, -10, 20}
var AllPotions = []*Potions{&PotionsSoins, &PotionsDamage}

type Armes struct {
	Name string
	Possede bool
	InStrore bool
	Multiplicateur float32
	Prix int	
	ObjectsCraft []*Objects
	RequiredQuantity []int
	Attaques	[]Attaques
    Mastery int
    MaxMastery int
}


var EpeeTraining=Armes{"\033[36m⚔️  Épées d'entraînement\033[0m",false,true,10,30, []*Objects{},[]int{},[]Attaques{AttaquesEpeeTraining},0,5,}
var GantThanos=Armes{"\033[35m🧤 Gant de Thanos\033[0m",false,false,100,1000,[]*Objects{&ScrapMetal,&InfinityStoneMind,&InfinityStoneReality,&InfinityStoneTime,&InfinityStonePouvoir,&InfinityStoneSpace,&InfinityStoneSoule},[]int{10,1,1,1,1,1,1},[]Attaques{AttaqueGantThanos,AttaqueGantThanos02,AttaqueGantThanos03,AttaqueGantThanos04,AttaqueGantThanos05,AttaqueGantSnap},0,10}
var MarteauThor=Armes{"\033[34m🔨 Marteau de Thor\033[0m",false,true,20,200,[]*Objects{},[]int{},[]Attaques{AttaquesMarteauLancé,AttaquesMarteauFoudre},0,15}
var StormBreaker=Armes{"\033[33m🪓 StormBreaker\033[0m",false,false,30,400,[]*Objects{&ScrapMetal,&GrootsBranches},[]int{5,5},[]Attaques{AttaqueHache,AttaqueHache02},0,20}
var AllWeapons = []*Armes{&EpeeTraining, &GantThanos, &MarteauThor, &StormBreaker}

type Armures struct {
	Name string
	Possede bool
	InStrore bool
	Protect int
	Prix int
	ObjectCraft []*Objects
	RequiredQuantity []int
}

var ArmureTraining=Armures{"\033[36m🥋 Armure d'entraînement\033[0m",false,true,10,15,[]*Objects{},[]int{}}
var ArmureBatman=Armures{"\033[34m🦇 Tenue de Batman\033[0m",false,false,20,30,[]*Objects{&CuireBat},[]int{15}}
var Armureamazone=Armures{"\033[35m🛡️ Tenue des Amazones\033[0m",false,true,15,20,[]*Objects{},[]int{}}
var AllArmures = []*Armures{&ArmureTraining, &ArmureBatman, &Armureamazone}

type Objects struct {
	Name string
	Nb int
	DropRate int
	MaxAmount int
}

type Teknologia struct {
    Name string
    ObjectsCraft []*Objects
    ObjectsQuentity []int
    Possede bool
}

var TeknologiaItem = Teknologia{"Teknologia", []*Objects{&KeyTK, &PlansMachine, &KeyTK}, []int{1, 1, 1}, false}

var ScrapMetal=Objects{"\033[90m⚙️  Scrape de Uruu\033[0m",10,75,10}
var InfinityStoneMind=Objects{"\033[35m🧠 Pierre de l'esprit\033[0m",1,15,1}
var InfinityStoneReality=Objects{"\033[35m🌌 Pierre de la réalité\033[0m",1,15,1}
var InfinityStoneTime=Objects{"\033[35m⏳ Pierre du temps\033[0m",1,15,1}
var InfinityStonePouvoir=Objects{"\033[35m⚡ Pierre du pouvoir\033[0m",1,15,1}
var InfinityStoneSpace=Objects{"\033[35m🪐 Pierre de l'espace\033[0m",1,15,1}
var InfinityStoneSoule=Objects{"\033[35m🧿 Pierre de l'âme\033[0m",1,15,1}
var GrootsBranches=Objects{"\033[32m🌿 Branche de Groot\033[0m",10,50,10}
var CuireBat=Objects{"\033[34m🦇 Aile de chauve-souris\033[0m",0,75,20}
var PlansMachine=Objects{"\033[33m📐 Plan pour réparer TeKnologia\033[0m",1,100,1}
var KeyTK=Objects{"\033[33m📐 Clé de TeKnologia\033[0m",1,100,1}
var AllObjects = []*Objects{&ScrapMetal, &InfinityStoneMind, &InfinityStoneReality, &InfinityStoneTime, &InfinityStonePouvoir, &InfinityStoneSpace, &InfinityStoneSoule, &GrootsBranches, &CuireBat, &KeyTK, &PlansMachine}


func DisplayInventoryObjects() {
    fmt.Print("\033[36m               Inventaire: Objets             \033[0m\n")
    var Choice int
    var Indexbrrr int
    for _, Obj := range AllObjects {
        if Obj.Nb > 0 {
            fmt.Printf("\033[33m- \033[0m%v x %d\n", Obj.Name, Obj.Nb)
        }
    }
    fmt.Print("\n")
    for i, Obj := range AllPotions {
        fmt.Printf("\033[33m%d --> \033[31m%v\033[0m x %d\n", i+1, Obj.Name, Obj.Nb)
        Indexbrrr ++
    }
    if TeknologiaItem.Possede {
        fmt.Print(Indexbrrr +1, " --> Teknologia")
    }
    fmt.Print("\n\033[33m3 --> \033[31mRetour\033[0m\n")
    fmt.Print("\033[33mQue voulez-vous faire ?\033[0m\n")
    fmt.Scan(&Choice)
    switch Choice {
    case 1:
        PotionsVie()
    case 2:
        fmt.Print("\033[31mCeci n'est pas une bonne idée, gardez-la pour d'autres occasions\033[0m\n")
        DisplayInventoryObjects()
    case 3:
        if !TeknologiaItem.Possede{
            DisplayInventory()
        } else {
            //Gio faut mettre ici ta fonction pour les crédits du jeu
        }
    case 4 :
        if !TeknologiaItem.Possede {
            fmt.Print("\033[31mCe n'est pas une option disponible\033[0m\n")
            DisplayInventoryObjects()
        } else {
            DisplayInventory()
        }
    default:
        fmt.Print("\033[31mCe n'est pas une option disponible\033[0m\n")
        DisplayInventoryObjects()
    }
}

func DisplayInventoryWeapons() {
    fmt.Print("\033[36m               Inventaire: Armes             \033[0m\n")
    var Choice int
    var index int = 1
    var Mapping []int
    for i, Obj := range AllWeapons {
        if Obj.Possede {
            if Obj.Name == Joueur.EquipedWeapon.Name {
                fmt.Printf("\033[33m%d --> \033[31m%v\033[0m \033[32m(équipée)\033[0m\n", index, Obj.Name)
            } else {
                fmt.Printf("\033[33m%d --> \033[31m%v\033[0m\n", index, Obj.Name)
            }
            Mapping = append(Mapping, i)
            index++
        }
    }
    fmt.Printf("\033[33m%d --> \033[31mRetour\033[0m\n", index)
    fmt.Print("\033[33mQue voulez-vous faire ?\033[0m\n")
    fmt.Scan(&Choice)
    if Choice == index {
        DisplayInventory()
        return
    }
    if Choice <= 0 || Choice > index {
        fmt.Println("\033[31mCe n'est pas une option disponible\033[0m")
        DisplayInventoryWeapons()
        return
    }
    Joueur.EquipedWeapon = *AllWeapons[Mapping[Choice-1]]
    fmt.Println("\033[33mVous avez équipé \033[0m", Joueur.EquipedWeapon.Name)
    Joueur.Puissance += float64(AllWeapons[Mapping[Choice-1]].Multiplicateur)
}

func DisplayInventoryArmures() {
    fmt.Print("\033[36m               Inventaire: Armures             \033[0m\n")
    var Choice int
    var index int = 1
    var mapping []int
    for i, Obj := range AllArmures {
        if Obj.Possede {
            if Obj.Name == Joueur.EquipedArmure.Name {
                fmt.Printf("\033[33m%d --> \033[31m%v\033[0m \033[32m(équipée)\033[0m\n", index, Obj.Name)
            } else {
                fmt.Printf("\033[33m%d --> \033[31m%v\033[0m\n", index, Obj.Name)
            }
            mapping = append(mapping, i)
            index++
        }
    }
    fmt.Printf("\033[33m%d --> \033[31mRetour\033[0m\n", index)
    fmt.Print("\033[33mQue voulez-vous faire ?\033[0m\n")
    fmt.Scan(&Choice)
    if Choice == index {
        DisplayInventory()
        return
    }
    if Choice <= 0 || Choice > index {
        fmt.Println("\033[31mCe n'est pas une option disponible\033[0m")
        DisplayInventoryArmures()
        return
    }
    Joueur.EquipedArmure = *AllArmures[mapping[Choice-1]]
    fmt.Println("\033[33mVous avez équipé \033[0m", Joueur.EquipedArmure.Name)
    Joueur.Viemax += AllArmures[mapping[Choice-1]].Protect
}

var LastArmorBaught  = []Armures{ArmureBatman, ArmureTraining, Armureamazone}
var LastWeaponBaught = []Armes{EpeeTraining, MarteauThor, GantThanos, StormBreaker}

func InstantEquipArmor() {
    var WantedEquip int
    fmt.Print("\033[33m\nVoulez-vous l'équiper maintenant ?\033[0m\n    \033[32m1. Oui\033[0m\n    \033[31m2. Non\033[0m\n")
    fmt.Scan(&WantedEquip)
    switch WantedEquip {
    case 1:
        Joueur.EquipedArmure = LastArmorBaught[WhatIsTheLastBoughtItem]
        fmt.Print("\033[33m\nVous avez équipé \033[31m", Joueur.EquipedArmure.Name, "\033[0m\n")
    case 2:
        Continuer()
        ShopArmures()
    default:
        fmt.Print("\033[31m\nCe n'est pas une des deux propositions, tu le fais exprès ?\033[0m\n")
        InstantEquipArmor()
    }
}

func InstantEquipWeapon() {
    var WantedEquip int
    fmt.Print("\033[33m\nVoulez-vous l'équiper maintenant ?\033[0m\n    \033[32m1. Oui\033[0m\n    \033[31m2. Non\033[0m\n")
    fmt.Scan(&WantedEquip)
    switch WantedEquip {
    case 1:
        Joueur.EquipedWeapon = LastWeaponBaught[WhatIsTheLastBoughtItem]
        fmt.Print("\033[33m\nVous avez équipé \033[31m", Joueur.EquipedWeapon.Name, "\033[0m\n")
    case 2:
        Continuer()
        ShopArmes()
    default:
        fmt.Print("\033[31m\nCe n'est pas une des deux propositions, tu le fais exprès ?\033[0m\n")
        InstantEquipWeapon()
    }
}