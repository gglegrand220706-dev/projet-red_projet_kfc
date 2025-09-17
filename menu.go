package personnage

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func MenuGeneral() {
	var selection int
	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[91m  _____         _                _             _           \033[0m")
	fmt.Println("\033[91m |_   _|__  ___| | ___ __   ___ | | ___   __ _(_) __ _     \033[0m")
	fmt.Println("\033[31m   | |/ _ \\/ __| |/ / '_ \\ / _ \\| |/ _ \\ / _` | |/ _` |    \033[0m")
	fmt.Println("\033[31m   | |  __/ (__|   <| | | | (_) | | (_) | (_| | | (_| |    \033[0m")
	fmt.Println("\033[31m   |_|\\___|\\___|_|\\_\\_| |_|\\___/|_|\\___/ \\__, |_|\\__,_|    \033[0m")
	fmt.Println("\033[31m  ____        _   _   _       __        _|___/     _     _ \033[0m")
	fmt.Println("\033[31m | __ )  __ _| |_| |_| | ___  \\ \\      / /__  _ __| | __| |\033[0m")
	fmt.Println("\033[31m |  _ \\ / _` | __| __| |/ _ \\  \\ \\ /\\ / / _ \\| '__| |/ _` |\033[0m")
	fmt.Println("\033[31m | |_) | (_| | |_| |_| |  __/   \\ V  V / (_) | |  | | (_| |\033[0m")
	fmt.Println("\033[31m |____/ \\__,_|\\__|\\__|_|\\___|    \\_/\\_/ \\___/|_|  |_|\\__,_|\033[0m")
	fmt.Print("\033[33mOptions :\n\033[0m")
	fmt.Print("\033[33m1 -->\033[0m \033[36mBase de donnees\033[0m\n\033[33m2 -->\033[0m \033[35mBoutique\033[0m\n\033[33m3 -->\033[0m \033[32mEntrainement\033[0m\n\033[33m4 -->\033[0m \033[34mQuête\033[0m\n")
	fmt.Scan(&selection)
	if selection == 1 {
		fmt.Print("\033[H\033[2J")
		MenuDataBase()
	}
	if selection == 2 {
		fmt.Print("\033[H\033[2J")
		MenuShop()
	}
	if selection == 3 {
		fmt.Print("\033[H\033[2J")
		AdverseryChoice = 0
		CombatMode()
	}
	if selection == 4 {
		fmt.Print("\033[H\033[2J")
		MenuQuetes()
	}
	if selection > 4 || selection <= 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Option indisponible, veuillez choisir parmi les 4 propositions\n")
		MenuGeneral()

	}
}

func MenuDataBase() {
	var selection int
	fmt.Print("\033[H\033[2J")
	fmt.Print("          BASE DE DONNEES          \n")
	fmt.Print("\033[33mOptions :\n\033[0m")
	fmt.Print("1. Fiche du personnage\n2. Inventaire\n3. Retour\n")
	fmt.Scan(&selection)

	if selection == 1 {
		DisplayInfo()
		RetourMenu()
	}
	if selection == 2 {
		fmt.Print("\033[H\033[2J")
		DisplayInventory()
		RetourMenu()
	}
	if selection == 3 {
		fmt.Print("\033[H\033[2J")
		MenuGeneral()
	}
	if selection > 3 || selection <= 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Option indisponible, veuillez choisir parmi les 3 propositions\n")
		MenuDataBase()
	}
}

func MenuShop() {
	var selection int
	fmt.Print("\033[H\033[2J")
	fmt.Print("          BOUTIQUE          \n")
	fmt.Print("Options\n")
	fmt.Print("1. Potions\n2. Armes\n3. Armures\n4. Forgeron\n5. Retour\n")
	fmt.Scan(&selection)
	if selection == 1 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("          BOUTIQUE          \n")
		fmt.Print("Quelle potion voulez vous acheter\n")
		ShopPotion()
		Continuer()
		MenuShop()
	}
	if selection == 2 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("          BOUTIQUE          \n")
		fmt.Print("\nQuelle arme voulez vous acheter\n")
		ShopArmes()
		RetourMenu()
	}
	if selection == 3 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("          BOUTIQUE          \n")
		fmt.Print("Quelle armure voulez vous acheter\n")
		ShopArmures()
		RetourMenu()
	}
	if selection == 4 {
		fmt.Print("\033[H\033[2J")
		BlackSmithDisplay()
		BlackSmith()
		Continuer()
		MenuShop()
	}
	if selection == 5 {
		MenuGeneral()
	}
	if selection > 5 || selection <= 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Option indisponible, veuillez choisir parmi les 4 propositions\n")
		MenuShop()
	}
}

func RetourMenu() {
	Continuer()
	MenuGeneral()
}

func DisplayAtackPhysqiue() {
	if Joueur.Vieactuelle > 0 && CurrentAdversery[0].Vieactuelle > 0 {
		var OptionDisplay = []string{}
		var Index int = 1
		for _, Attaque := range Joueur.Attaques {
			OptionDisplay = append(OptionDisplay, Attaque.Name)
		}
		fmt.Print("\nQue voulez-vous faire :\n")
		for index, OptionsName := range OptionDisplay {
			fmt.Printf("%d. %v\n", index+1, OptionsName)
			index++
			Index++
		}
		fmt.Print(Index, ". Retour\n")
	}
}

func DisplayAtackArmes() {
	if Joueur.Vieactuelle > 0 && CurrentAdversery[0].Vieactuelle > 0 {
		var OptionDisplay = []string{}
		var Index int
		for _, Attaque := range Joueur.EquipedWeapon.Attaques {
			OptionDisplay = append(OptionDisplay, Attaque.Name)
		}
		fmt.Print("\nQue voulez-vous faire :\n")
		for index, OptionsName := range OptionDisplay {
			fmt.Printf("%d. %v\n", index+1, OptionsName)
			index++
			Index++
		}
		fmt.Print(Index+1, ". Retour\n")
	}
}

func IsDead() {
	fmt.Print("Vous êtes mort sale noob\n")
	fmt.Print("vous récussitez avec ", Joueur.Viemax/2, " PV\n")
	Joueur.Vieactuelle = Joueur.Viemax / 2
	RetourMenu()
}

func Continuer() {
	time.Sleep(1 * time.Second)
	fmt.Print("\nAppuyer sur esapce pour continuer ...\n")
	bufio.NewReader(os.Stdin).ReadBytes(' ')
}

func DisplayPotions() {
	var index = 0
	if Joueur.Vieactuelle > 0 && CurrentAdversery[AdverseryChoice].Vieactuelle > 0 {
		var PotionsDisplay = []string{}
		for _, NamePotion := range AllPotions {
			PotionsDisplay = append(PotionsDisplay, NamePotion.Name)
		}
		fmt.Print("quelle potion voulez vous utilisé ?\n")
		for _, PotionsName := range PotionsDisplay {
			fmt.Printf("%d. %v\n", index+1, PotionsName)
			index++
		}
		fmt.Print("3. Retour\n")
	}
}

func Fuite() {
	if Joueur.Bourse-CurrentAdversery[AdverseryChoice].GivenMoney > 0 {
		fmt.Print("Vous vous êtes fais raquètter pour partire, vous perdez : ", CurrentAdversery[AdverseryChoice].GivenMoney, " Techno-Dollars\n")
		Joueur.Bourse -= CurrentAdversery[AdverseryChoice].GivenMoney
	} else {
		fmt.Print("\nVous fuyez et vous avez pas d'argents, vous êtes pitoyable.")
	}
	if Joueur.EXP-CurrentAdversery[AdverseryChoice].GivenExp > 0 {
		Joueur.EXP -= CurrentAdversery[AdverseryChoice].GivenExp
	} else {
		fmt.Print("\nVous fuyez car vous êtes sans experience, vous êtes pitoyable.")
	}
	RetourMenu()
}

func MenuQuetes() {
	var SelectionQuete int
	fmt.Print("quelle quete voulez vous faire ?\n1. Bosses\n2. Enemys Classique\n")
	fmt.Scan(&SelectionQuete)
	if SelectionQuete == 1 {
		fmt.Print("\033[H\033[2J")
		AdverseryChoice = 1
		CombatMode()
	}
	if SelectionQuete == 2 {
		fmt.Print("\033[H\033[2J")
		AdverseryChoice = 2
		CombatMode()
	} else {
		fmt.Print("se n'est pas une option disponible")
	}

}

func CombatMode() {
	fmt.Print("\033[H\033[2J")
	CurrentAdversery[AdverseryChoice].Vieactuelle = CurrentAdversery[AdverseryChoice].Viemax
	var Selection02 int = 0
	for Joueur.Vieactuelle > 0 && CurrentAdversery[AdverseryChoice].Vieactuelle > 0 {
		fmt.Print("\nPV :", Joueur.Vieactuelle, "/", Joueur.Viemax, "                                               ", CurrentAdversery[AdverseryChoice].Nom, " : ", CurrentAdversery[AdverseryChoice].Vieactuelle, "/", CurrentAdversery[AdverseryChoice].Viemax, "\n\n")
		fmt.Print("que voulez vous faire ?\n1. Attaque\n2. Potions\n3. Fuire\n")
		fmt.Scan(&Selection02)
		if Selection02 == 1 {
			fmt.Print("\nque voulez vous faire ?\n1. Attaque Physique\n2. Attaques Arme\n3. Retour\n")
			var Selection03 int
			fmt.Scan(&Selection03)
			if Selection03 == 1 {
				DisplayAtackPhysqiue()
				AtackBasiqueSystème()
				Response()
				PotionDamage()

			}
			if Selection03 == 2 {
				DisplayAtackArmes()
				AtackWeaponSystème()
				Response()
				PotionDamage()
			}
			if Selection03 == 3 {
				CombatMode()
			}
			if Selection03 < 0 || Selection03 > 3 {
				fmt.Print("non disponible")
				CombatMode()
			}
		}
		if Selection02 == 2 {
			DisplayPotions()
			PotionCHoice()
			Response()
		}
		if Selection02 == 3 {
			fmt.Print("\033[H\033[2J")
			Fuite()

		}
		if Joueur.Vieactuelle <= 0 {
			IsDead()
		}
		if CurrentAdversery[AdverseryChoice].Vieactuelle <= 0 {
			Reward()
			DropRate()
			RetourMenu()
		}
		if Selection02 < 0 || Selection02 > 3 {
			fmt.Print("Ce n'est pas un choix disponoble\n")
		}
	}
}

var K int // indice pour le retour

func BlackSmithDisplay() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("                 Forgeron\n")
	var ArmesForgeable = []*Armes{&GantThanos, &StormBreaker}
	var IndexObj int

	for index, Name := range ArmesForgeable {
		fmt.Print(index+1, ". ", Name.Name, "\n")
		K++
		IndexObj = 0
		for IndexObj < len(ArmesForgeable[index].ObjectsCraft) {
			fmt.Print("           - ", ArmesForgeable[index].ObjectsCraft[IndexObj].Name, "\n")
			IndexObj++
		}
	}
	fmt.Print(K, ". Retour")
}

func DisplayInventory() {
	fmt.Print("\033[H\033[2J")
	var ChoiceSection int
	fmt.Print("               Invetaire\n")
	fmt.Print("\n que voulez vous vérifier ?\n\n1. Objets\n2. Armes\n3. Armures\n4. Retour\n")
	fmt.Scan(&ChoiceSection)
	switch ChoiceSection {
	case 1:
		DisplayInventoryObjects()
	case 2:
		DisplayInventoryWeapons()
	case 3:
		DisplayInventoryArmures()
	case 4:
		RetourMenu()
	default:
		fmt.Print("se n'est pas une Option disponible")
		DisplayInventory()
	}
}
