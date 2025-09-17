package personnage

import (
	"fmt"
)

func ShopPotion() {
	var ChoixPotion int
	fmt.Print("\nVous avez ", Joueur.Bourse, " TeKno-dollar\n")
	fmt.Print("\n1. potion de soins\n")
	fmt.Print("2. potion de poison\n")
	fmt.Print("3. Retour\n")
	var PotionDispo = []*Potions{&PotionsSoins, &PotionsDamage}
	fmt.Scan(&ChoixPotion)	
	if PotionDispo[ChoixPotion - 1].Prix <= Joueur.Bourse && PotionDispo[ChoixPotion - 1].Nb < 5{
		Joueur.Bourse -= PotionDispo[ChoixPotion - 1].Prix
		PotionDispo[ChoixPotion - 1].Nb ++
		fmt.Print("Vous avez acheter ", PotionDispo[ChoixPotion -1].Name, " il vous reste donc ", Joueur.Bourse, " TeKno-Dollars")
		if ChoixPotion == 3 {
			MenuShop()
		}
		if ChoixPotion < 0 || ChoixPotion > 4 {
			fmt.Print("se n'est pas une proposition disponible")
			ShopPotion()
		}
	}
}

func ShopArmes() {
	var ChoixArmes int
	fmt.Print("\nVous avez ", Joueur.Bourse, " TeKno-Dollar\n")
	var ArmesDispo = []*Armes{&EpeeTraining, &MarteauThor, &GantThanos, &StormBreaker}
	var Index int
	for i, Name := range ArmesDispo {
		if Name.InStrore {
			fmt.Println(i + 1, Name.Name,  Name.Prix, " TKD")
			Index = i
		}
	}
	fmt.Print(Index + 2, ". Retour\n")
	fmt.Scan(&ChoixArmes)
	if ChoixArmes == Index + 2 {
		MenuShop()
	}
	if ChoixArmes < 0 || ChoixArmes > len(ArmesDispo) {
		fmt.Print("se n'est pas une option disponible")
		Continuer()
		ShopArmes()
	}
	if ArmesDispo[ChoixArmes -1].Prix <= Joueur.Bourse && !ArmesDispo[ChoixArmes-1].Possede{
		Joueur.Bourse -= ArmesDispo[ChoixArmes -1].Prix
		ArmesDispo[ChoixArmes -1].Possede = true
		fmt.Print("vous avez acheté ", ArmesDispo[ChoixArmes -1].Name, "\n")
		fmt.Print("il vous reste : ", Joueur.Bourse, " TeKno-Dollar\n")
		Continuer()
		MenuShop()
	}else {
		fmt.Print("Vous ne pouvez pas acheter :P\n")
		Continuer()
		ShopArmes()
	}
}

func ShopArmures() {
	fmt.Print("Vous avez ", Joueur.Bourse, " TeKno-Dollar\n")
	var ChoixArmures int
	var ArmuresDispo = []*Armures{&ArmureTraining}
	for i, Name := range ArmuresDispo {
		fmt.Println(i+1, Name.Name, Name.Prix, "TKD")
	}
	fmt.Scan(&ChoixArmures)
	if ChoixArmures < 0 || ChoixArmures > len(ArmuresDispo) {
		fmt.Print("se n'est pas un choix disponible")
		Continuer()
		ShopArmures()
	}
	if Joueur.Bourse >= ArmuresDispo[ChoixArmures -1].Prix && !ArmuresDispo[ChoixArmures -1].Possede {
		Joueur.Bourse -= ArmuresDispo[ChoixArmures -1].Prix
		ArmuresDispo[ChoixArmures -1].Possede = true
		fmt.Print("Vous avez acheter ", ArmuresDispo[ChoixArmures -1].Name, " il vous reste ", Joueur.Bourse, " TeKno-Dollars")
		Continuer()
		MenuShop()
	}
}

func BlackSmith() {
    var CraftingChoice int
    fmt.Print("Que voulez vous Aventurier ?\n")
    fmt.Scan(&CraftingChoice)
    var ArmesForgeable = []*Armes{&GantThanos, &StormBreaker}
    var Enough int
    var IndexFreeFire int
	if CraftingChoice == K{
		MenuShop()
	}
    for IndexFreeFire < len(ArmesForgeable[CraftingChoice -1].ObjectsCraft) {
        if ArmesForgeable[CraftingChoice -1].ObjectsCraft[IndexFreeFire].Nb >= ArmesForgeable[CraftingChoice -1].RequiredQuantity[IndexFreeFire] {
            Enough++
            IndexFreeFire++
        } else {
            fmt.Print("Vous n'avez pas les matériaux\n")
            return
        }
    }
    if Enough == len(ArmesForgeable[CraftingChoice -1].ObjectsCraft) {
        for i := 0; i < len(ArmesForgeable[CraftingChoice -1].ObjectsCraft); i++ {
            ArmesForgeable[CraftingChoice -1].ObjectsCraft[i].Nb -= ArmesForgeable[CraftingChoice -1].RequiredQuantity[i]
        }
        fmt.Print("Vous retrouverez cette arme dans le magasin\n")
		ArmesForgeable[CraftingChoice -1].InStrore = true
    }
}
