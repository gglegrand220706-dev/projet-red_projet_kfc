package personnage

import (
	"fmt"
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
        WhatIsTheLastBoughtItem = ChoixArmures
        InstantEquipArmor()
        Continuer()
        ShopArmures()
    }
}

func BlackSmith() {
    var CraftingChoice int
    fmt.Print("\033[33mQue voulez-vous, Aventurier ?\033[0m\n")
    fmt.Scan(&CraftingChoice)
    var ArmesForgeable = []*Armes{&GantThanos, &StormBreaker}
	var ArmuresForgeable =[]*Armures{&ArmureBatman}
	var TeknologiaCraft = &TeknologiaItem
    var Enough int
    var IndexFreeFire int
	var IndexTek int

    if CraftingChoice == K +2 {
        MenuShop()
        return
    }
	if 1>CraftingChoice && CraftingChoice < 2 {
		for IndexFreeFire < len(ArmesForgeable[CraftingChoice-1].ObjectsCraft) {
        if ArmesForgeable[CraftingChoice-1].ObjectsCraft[IndexFreeFire].Nb >= ArmesForgeable[CraftingChoice-1].RequiredQuantity[IndexFreeFire] {
            Enough++
            IndexFreeFire++
        } else {
            fmt.Print("\033[31mVous n'avez pas les matériaux\033[0m\n")
            return
        }
    }
    if Enough == len(ArmesForgeable[CraftingChoice-1].ObjectsCraft) {
        for i := 0; i < len(ArmesForgeable[CraftingChoice-1].ObjectsCraft); i++ {
            ArmesForgeable[CraftingChoice-1].ObjectsCraft[i].Nb -= ArmesForgeable[CraftingChoice-1].RequiredQuantity[i]
        }
        fmt.Print("\033[32mVous retrouverez cette arme dans le magasin\033[0m\n")
        ArmesForgeable[CraftingChoice-1].InStrore = true
    }
	}
    if CraftingChoice == 3 {
		for IndexFreeFire < len(ArmuresForgeable[CraftingChoice-1].ObjectCraft) {
        if ArmuresForgeable[CraftingChoice-1].ObjectCraft[IndexFreeFire].Nb >= ArmuresForgeable[CraftingChoice-1].RequiredQuantity[IndexFreeFire] {
            Enough++
            IndexFreeFire++
        } else {
            fmt.Print("\033[31mVous n'avez pas les matériaux\033[0m\n")
            return
        }
    }
    if Enough == len(ArmuresForgeable[CraftingChoice-1].ObjectCraft) {
        for i := 0; i < len(ArmuresForgeable[CraftingChoice-1].ObjectCraft); i++ {
            ArmuresForgeable[CraftingChoice-1].ObjectCraft[i].Nb -= ArmuresForgeable[CraftingChoice-1].RequiredQuantity[i]
        }
        fmt.Print("\033[32mVous retrouverez cette armure dans le magasin\033[0m\n")
        ArmuresForgeable[CraftingChoice-1].InStrore = true
    }
	}
	if CraftingChoice == 4 {
		for IndexTek < len(TeknologiaCraft.ObjectsCraft){
			if TeknologiaCraft.ObjectsCraft[IndexTek].Nb == TeknologiaCraft.ObjectsQuentity[IndexTek]{
				Enough++
				IndexTek ++
			} else {
				fmt.Print("vous n'avez pas les matériaux")
				return
			}
		}
		if Enough == len(TeknologiaCraft.ObjectsCraft) {
			for i :=0; i< len(TeknologiaCraft.ObjectsCraft); i++ {
				TeknologiaCraft.ObjectsCraft[i].Nb -= TeknologiaCraft.ObjectsQuentity[i]
			}
			fmt.Print("Vous avez réparer Teknologia, ouvrez votre invetaire pour l'utiliser")
			TeknologiaCraft.Possede = true
		}
	}
}

