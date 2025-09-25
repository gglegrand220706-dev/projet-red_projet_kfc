package personnage

import (
	"fmt"
	"math/rand"
    "time"
)

func ShopPotion() {
    var ChoixPotion int
    fmt.Printf("\n\033[33mVous avez \033[32m%d 🪙  TeKno-Dirham\033[0m\n\n", Joueur.Bourse)
    fmt.Print("\033[33m1 --> \033[31mPotion de soins ❤️\033[0m\n")
    fmt.Print("\033[33m2 --> \033[32mPotion de poison ☠️\033[0m\n")
    fmt.Print("\033[33m3 --> \033[31mRetour\033[0m\n")
    var PotionDispo = []*Potions{&PotionsSoins, &PotionsDamage}
    fmt.Scan(&ChoixPotion)

    if ChoixPotion == 3 {
        MenuShop()
        return
    }
    if ChoixPotion < 1 || ChoixPotion > 2 {
        fmt.Print("\033[31mCe n'est pas une proposition disponible\033[0m\n")
        ShopPotion()
        return
    }
    if PotionDispo[ChoixPotion-1].Prix <= Joueur.Bourse && PotionDispo[ChoixPotion-1].Nb < 5 {
        Joueur.Bourse -= PotionDispo[ChoixPotion-1].Prix
        PotionDispo[ChoixPotion-1].Nb++
        fmt.Printf("\033[32m✨ Vous avez acheté \033[0m%v\033[32m  il vous reste \033[32m%d 🪙  TeKno-Dirham\033[0m\n",
            PotionDispo[ChoixPotion-1].Name, Joueur.Bourse)
        ShopPotion()
    }
}

var WhatIsTheLastBoughtItem int

func ShopArmes() {
    fmt.Printf("\n\033[33mVous avez \033[32m%d 🪙  TeKno-Dirham\033[0m\n\n", Joueur.Bourse)
    var ArmesDispo = []*Armes{&EpeeTraining, &MarteauThor, &GantThanos, &StormBreaker}
    var ChoixArmes int
    var Index int

    for i, Name := range ArmesDispo {
        if Name.InStrore {
            if Name.Possede {
                fmt.Printf("\033[33m%d --> \033[31m%v\033[0m \033[31m(déjà possédée)\033[0m\n", i+1, Name.Name)
            } else {
                fmt.Printf("\033[33m%d --> \033[31m%v\033[0m \033[32m%d 🪙  TKD\033[0m\n", i+1, Name.Name, Name.Prix)
            }
            Index = i
        }
    }

    fmt.Printf("\033[33m%d --> \033[31mRetour\033[0m\n", Index+2)
    fmt.Scan(&ChoixArmes)

    if ChoixArmes == Index+2 {
        MenuShop()
        return
    }
    if ChoixArmes < 1 || ChoixArmes > len(ArmesDispo) {
        fmt.Print("\033[31mCe n'est pas une option disponible\033[0m\n")
        Continuer()
        ShopArmes()
        return
    }
    if ArmesDispo[ChoixArmes-1].Prix <= Joueur.Bourse && !ArmesDispo[ChoixArmes-1].Possede {
        Joueur.Bourse -= ArmesDispo[ChoixArmes-1].Prix
        ArmesDispo[ChoixArmes-1].Possede = true
        fmt.Printf("\033[32m✨ Vous avez acheté \033[0m%v\033[32m\n Il vous reste : \033[32m%d  🪙  TeKno-Dirham\033[0m\n",
            ArmesDispo[ChoixArmes-1].Name, Joueur.Bourse)
        WhatIsTheLastBoughtItem = ChoixArmes - 1
        InstantEquipWeapon()
        Continuer()
        ShopArmes()
    } else {
        fmt.Print("\033[31mVous ne pouvez pas acheter :P\033[0m\n")
        Continuer()
        ShopArmes()
    }
}	



func ShopArmures() {
    fmt.Printf("\033[33mVous avez \033[32m%d  🪙  TeKno-Dirham\033[0m\n\n", Joueur.Bourse)
    var ChoixArmures int
    var Index int
    var ArmuresDispo = []*Armures{&ArmureTraining, &Armureamazone, &ArmureBatman}

    for i, Name := range ArmuresDispo {
        if Name.InStrore {
            if Name.Possede {
                fmt.Printf("\033[33m%d --> \033[31m%v\033[0m \033[31m(déjà possédée)\033[0m\n", i+1, Name.Name)
            } else {
                fmt.Printf("\033[33m%d --> \033[31m%v\033[0m \033[32m%d  🪙 TKD\033[0m\n", i+1, Name.Name, Name.Prix)
            }
            Index = i
        }
    }

    fmt.Printf("\033[33m%d --> \033[31mRetour\033[0m\n", Index+2)
    fmt.Scan(&ChoixArmures)

    if ChoixArmures == Index+2 {
        MenuShop()
        return
    }
    if ChoixArmures < 1 || ChoixArmures > len(ArmuresDispo) {
        fmt.Print("\033[31mCe n'est pas un choix disponible\033[0m\n")
        Continuer()
        ShopArmures()
        return
    }
    if Joueur.Bourse >= ArmuresDispo[ChoixArmures-1].Prix && !ArmuresDispo[ChoixArmures-1].Possede {
        Joueur.Bourse -= ArmuresDispo[ChoixArmures-1].Prix
        ArmuresDispo[ChoixArmures-1].Possede = true
        fmt.Printf("\033[32m✨ Vous avez acheté \033[0m%v\033[32m,  il vous reste \033[32m%d  🪙  1TeKno-Dirham\033[0m\n",
            ArmuresDispo[ChoixArmures-1].Name, Joueur.Bourse)
        	WhatIsTheLastBoughtItem = ChoixArmures -1
        InstantEquipArmor()
        Continuer()
        ShopArmures()
    }
}

func BlackSmithArmes() {
    var CraftingChoice int
    DynamiqueType("\033[33mQue voulez-vous, Aventurier ?\033[0m\n", 60)
    fmt.Scan(&CraftingChoice)
    IndexArmesForgeable = 1
    var ArmesForgeable = []*Armes{&GantThanos, &StormBreaker}
    var Enough bool = false
    fmt.Print("ça marche ici")
    if CraftingChoice == len(ArmesForgeable) + 1 {
        MenueBlackSmith()
        return
    }
	if 1 == CraftingChoice || CraftingChoice == 2 {
        fmt.Print("ça marche ici")
        for IndexArmesForgeable < len(ArmesForgeable[CraftingChoice-1].ObjectsCraft) {
            if ArmesForgeable[CraftingChoice-1].ObjectsCraft[IndexArmesForgeable-1].Nb >= ArmesForgeable[CraftingChoice-1].RequiredQuantity[IndexArmesForgeable-1] {
                Enough = true
                IndexArmesForgeable++
            } else {
                fmt.Print("\033[31mVous n'avez pas les matériaux\033[0m\n")
                Enough = false
                return          
            }
        }
        if Enough {
            for i := 0; i < len(ArmesForgeable[CraftingChoice-1].ObjectsCraft); i++ {
                ArmesForgeable[CraftingChoice-1].ObjectsCraft[i].Nb -= ArmesForgeable[CraftingChoice-1].RequiredQuantity[i]
            }
            fmt.Print("\033[32mVous retrouverez cette arme dans le magasin\033[0m\n")
            ArmesForgeable[CraftingChoice-1].InStrore = true
            Enough = false
        }
	}	
}

func BlackSmithArmures() {
    var CraftingChoice int
    DynamiqueType("\033[33mQue voulez-vous, Aventurier ?\033[0m\n", 60)
    fmt.Scan(&CraftingChoice)
    var ArmuresForgeable = []*Armures{&ArmureBatman}
    var Enough bool = false

    if CraftingChoice == len(ArmuresForgeable) + 1 {
        MenueBlackSmith()
        return
    }
	if CraftingChoice == 1 {
        for IndexArmureForgeable <= len(ArmuresForgeable[CraftingChoice-1].ObjectCraft) {
            if ArmuresForgeable[CraftingChoice-1].ObjectCraft[IndexArmureForgeable-1].Nb >= ArmuresForgeable[CraftingChoice-1].RequiredQuantity[IndexArmureForgeable-1] {
                Enough = true
                IndexArmureForgeable++
            } else {
                fmt.Print("\033[31mVous n'avez pas les matériaux\033[0m\n")
                Enough = false
                return
            }
        }
        if Enough  {
            for i := 0; i < len(ArmuresForgeable[CraftingChoice-1].ObjectCraft); i++ {
                ArmuresForgeable[CraftingChoice-1].ObjectCraft[i].Nb -= ArmuresForgeable[CraftingChoice-1].RequiredQuantity[i]
            }
            fmt.Print("\033[32mVous retrouverez cette arme dans le magasin\033[0m\n")
            ArmuresForgeable[CraftingChoice-1].InStrore = true
        }
	}	
}

func BlackSmithObjets() {
    var CraftingChoice int
    DynamiqueType("\033[33mQue voulez-vous, Aventurier ?\033[0m\n", 60)
    fmt.Scan(&CraftingChoice)
    var ObjetForgeables = []*Teknologia{&TeknologiaItem}
    var Enough bool = false

    if CraftingChoice == len(ObjetForgeables) +1 {
        MenueBlackSmith()
        return
    }
	if 1 == CraftingChoice {
		for IndexObjetsForgeable < len(ObjetForgeables[CraftingChoice-1].ObjectsCraft) {
            if ObjetForgeables[CraftingChoice-1].ObjectsCraft[IndexObjetsForgeable-1].Nb >=  ObjetForgeables[CraftingChoice-1].ObjectsQuentity[IndexObjetsForgeable-1] {
                Enough = true
                IndexObjetsForgeable++
            } else {
                fmt.Print("\033[31mVous n'avez pas les matériaux\033[0m\n")
                Enough = false
                return
            }
        }
        if Enough {
            for i := 0; i < len(ObjetForgeables[CraftingChoice-1].ObjectsCraft); i++ {
                ObjetForgeables[CraftingChoice-1].ObjectsCraft[i].Nb -= ObjetForgeables[CraftingChoice-1].ObjectsQuentity[i]
            }
            DynamiqueType("\033[32mVous retrouverez cette Objets dans votre inventaire\033[0m\n", 60)
            ObjetForgeables[CraftingChoice-1].Possede = true
        }
	}	
}

func CasinoMachineSous() {
    fmt.Print("\033[H\033[2J")
    fmt.Print("               Bien venue au TeKno -Casino")
    var Replay string
    var FDigit int
    var SDigit  int
    var TDigit int
    var ForthDigit int
    var AllDigits []int
    var PlayedMoney int
    fmt.Print("\n\nCombien voulez vous miser ? (entrez 0 pour quitter)\n\n")
    fmt.Scan(&PlayedMoney)
    if PlayedMoney == 0 {
        CasinoMenu()
    }
    for PlayedMoney > Joueur.Bourse {
        if PlayedMoney > Joueur.Bourse {
            fmt.Print("vous n'avez pas autant d'argent que ça voyons\n")
            fmt.Print("\nCombien voulez vous miser ?\n\n")
            fmt.Scan(&PlayedMoney)
        }
    }
    Joueur.Bourse -= PlayedMoney
    FDigit = rand.Intn(8)
    SDigit = rand.Intn(8)
    TDigit = rand.Intn(8)
    ForthDigit = rand.Intn(8)
    AllDigits = append(AllDigits, FDigit, SDigit, TDigit, ForthDigit)
    time.Sleep(1 * time.Second)
    for i, Numbers := range AllDigits {
        if i == len(AllDigits) -1 {
            fmt.Print( Numbers)
            time.Sleep(1 * time.Second)
        } else {
            fmt.Print( Numbers, "-")
        time.Sleep(1 * time.Second)
        }       
    }
    if FDigit == SDigit && TDigit == ForthDigit && SDigit == TDigit{
        fmt.Print("\nVous Avez Gangné !!!\nVotre Bourse Gagne ", 9*PlayedMoney)
        Joueur.Bourse += 10*PlayedMoney
        fmt.Print("Voulez Vous Rejouer ? (o/n)")
        fmt.Scan(&Replay)
        for Replay != "o" || Replay != "n" {
            switch Replay {
                case "o" :
                    CasinoMachineSous()
                case "n" :
                    MenuTeknoTown()
                    default : 
                        fmt.Print("veuillez choisire parmis les options proposée \n")
                        fmt.Print("Voulez Vous Rejouer ? (o/n)\n")
                        fmt.Scan(&Replay)
            }
        }
        
    } else {
        fmt.Print("\nDommage, Voulez vous rejouer ? (o/n) ")
        fmt.Scan(&Replay)
       for Replay != "o" || Replay != "n" {
            switch Replay {
                case "o" :
                    CasinoMachineSous()
                case "n" :
                    MenuTeknoTown()
                    default : 
                        fmt.Print("veuillez choisire parmis les options proposée \n")
                        fmt.Print("Voulez Vous Rejouer ? (o/n)\n")
                    fmt.Scan(&Replay)
                }
            }
    }
}

func RouletteCasino() {
    fmt.Print("\033[H\033[2J")
    var Couleur string
    var Replay string
    var PlayedMoney int
    var ChoisirNombreConf string
    var NombreChoisi int
    var RandColorCompare string
    var RandColor string
    var RandNumber int
    fmt.Print("               Bien venue à la roulette du TeKno -Casino\n\nCombien voulez Vous miser ? (entrez 0 pour quitter)\n")
    fmt.Scan(&PlayedMoney)
    if PlayedMoney == 0 {
        CasinoMenu()
    }
    fmt.Print("\n Quelle Couleur voulez vous choisire ? (r/n) ")
    fmt.Scan(&Couleur)
    for ChoisirNombreConf != "o" && ChoisirNombreConf != "n" {
        fmt.Print("\n\n voulez vous Choisir un nombre ? (o/n) ")
        fmt.Scan(&ChoisirNombreConf)
    }   
    switch ChoisirNombreConf {
    case "o" :
        fmt.Print("\n\nQuelle Nombre voulez vous choisire entre 0 et 36 \n")
        fmt.Scan(&NombreChoisi)
        fmt.Print("\nParfait êtes vous prêts ? c'est Partit !!!\n")
        RandNumber = rand.Intn(37)
        if  RandNumber%2 == 0 {
            RandColor = "rouge"
            RandColorCompare = "r"
            fmt.Print("\nvous êtes tomber sur ", RandNumber, " qui est une case ", RandColor, "\n")
        } else {
            RandColor = "noir"
            RandColorCompare = "n"
            fmt.Print("\nvous êtes tomber sur ", RandNumber, " qui est une case ", RandColor, "\n")
        }
        if RandColorCompare == Couleur && RandNumber == NombreChoisi {
            fmt.Print("\nCous avez gagné le gros lots !!! vous repartez avec ", PlayedMoney*36, "\n")
            Joueur.Bourse += PlayedMoney
        } else {
            fmt.Print("\nDommage, voulez vous rejouer ?? (o/n) ")
            Joueur.Bourse -= PlayedMoney
            fmt.Scan(&Replay)
            for Replay != "o" || Replay != "n"{
            switch Replay {
            case "o" :
                RouletteCasino()
            case "n" :
                CasinoMenu()
                default : 
                fmt.Print("\nse n'est pas une option disponible !!!\n")
            }
        }
        }
    case "n" :
        fmt.Print("\nParfait êtes vous prêts ? c'est Partit !!!\n")
        RandNumber = rand.Intn(37)
        if  RandNumber%2 == 0 {
            RandColor = "rouge"
            RandColorCompare = "r"
            fmt.Print("\nvous êtes tomber sur ", RandNumber, " qui est une case ", RandColor, "\n")
        } else {
            RandColor = "noir"
            RandColorCompare = "n"
            fmt.Print("\nvous êtes tomber sur ", RandNumber, " qui est une case ", RandColor, "\n")
        }
        if RandColorCompare == Couleur {
            fmt.Print("\nCous avez gagné !!! vous repartez avec ", PlayedMoney*2, "\n")
            Joueur.Bourse += PlayedMoney
        } else {
            fmt.Print("\nDommage, voulez vous rejouer ?? (o/n) ")
            Joueur.Bourse -= PlayedMoney
            fmt.Scan(&Replay)
            for Replay != "o" || Replay != "n"{
                switch Replay {
                case "o" :
                    RouletteCasino()
                case "n" :
                    CasinoMenu()
                    default :
                    fmt.Print("\ntoujours pas une option disponible\n")
                }
            }
        }
        default : 
        fmt.Print("\nse n'est pas une option disponible !!!!\n")
    }
}
