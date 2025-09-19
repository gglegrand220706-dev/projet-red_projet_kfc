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
	fmt.Print("\033[33m1 -->\033[0m \033[31m📚 Base de données\033[0m\n\033[33m2 -->\033[0m \033[31m🛒 Boutique\033[0m\n\033[33m3 -->\033[0m \033[31m🏋️ Entraînement\033[0m\n\033[33m4 -->\033[0m \033[31m🗺️  Quête\033[0m\n\033[33m5 -->\033[0m \033[31m🚪 Quitter\033[0m\n")
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
		Joueur.AbsorbedDamage = 0
		AdverseryChoice = 0
		CombatMode()
	}
	if selection == 4 {
		fmt.Print("\033[H\033[2J")
		RenvoieSinetique.Damage = 0
		MenuQuetes()
	}
	if selection == 5 {
		return
	}
	if selection > 5 || selection <= 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Option indisponible, veuillez choisir parmi les 4 propositions\n")
		MenuGeneral()

	}
}

func MenuDataBase() {
    var selection int
    fmt.Print("\033[H\033[2J")
    fmt.Print("\033[36m          Tekno-Book          \033[0m\n")
    fmt.Print("\033[33mOptions :\033[0m\n")
    fmt.Print("\033[33m1 --> \033[31mFiche du personnage\033[0m\n")
    fmt.Print("\033[33m2 --> \033[31mInventaire\033[0m\n")
    fmt.Print("\033[33m3 --> \033[31mRetour\033[0m\n")
    fmt.Print("\033[33mQue voulez-vous faire ?\033[0m\n")
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
        fmt.Print("\033[31m❌ Option indisponible, veuillez choisir parmi les 3 propositions\033[0m\n")
        MenuDataBase()
    }
}

func MenuShop() {
    var selection int
    fmt.Print("\033[H\033[2J")
    fmt.Print("\033[36m          BOUTIQUE          \033[0m\n")
    fmt.Print("\033[33mOptions :\033[0m\n")
    fmt.Print("\033[33m1 --> \033[31mPotions\033[0m\n")
    fmt.Print("\033[33m2 --> \033[31mArmes\033[0m\n")
    fmt.Print("\033[33m3 --> \033[31mArmures\033[0m\n")
    fmt.Print("\033[33m4 --> \033[31mForgeron\033[0m\n")
    fmt.Print("\033[33m5 --> \033[31mRetour\033[0m\n")
    fmt.Print("\033[33mQue voulez-vous faire ?\033[0m\n")
    fmt.Scan(&selection)

    if selection == 1 {
        fmt.Print("\033[H\033[2J")
        fmt.Print("\033[36m          BOUTIQUE          \033[0m\n")
        fmt.Print("\033[33mQuelle potion voulez-vous acheter ?\033[0m\n")
        ShopPotion()
        Continuer()
        MenuShop()
    }
    if selection == 2 {
        fmt.Print("\033[H\033[2J")
        fmt.Print("\033[36m          BOUTIQUE          \033[0m\n")
        fmt.Print("\033[33mQuelle arme voulez-vous acheter ?\033[0m\n")
        ShopArmes()
        RetourMenu()
    }
    if selection == 3 {
        fmt.Print("\033[H\033[2J")
        fmt.Print("\033[36m          BOUTIQUE          \033[0m\n")
        fmt.Print("\033[33mQuelle armure voulez-vous acheter ?\033[0m\n")
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
        fmt.Print("\033[31m❌ Option indisponible, veuillez choisir parmi les 5 propositions\033[0m\n")
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

        fmt.Print("\n\033[33mQue voulez-vous faire :\033[0m\n")
        for index, OptionsName := range OptionDisplay {
			if OptionsName == "Renvoie Sinetqiue" {
				fmt.Printf("\033[33m%d --> \033[31m%v %c\033[0m\n", index+1, RenvoieSinetique.Name, RenvoieSinetique.Damage)
			} else {
				fmt.Printf("\033[33m%d --> \033[31m%v\033[0m\n", index+1, OptionsName)
            index++
            Index++
			}       
        }
        fmt.Printf("\033[33m%d --> \033[31mRetour\033[0m\n", Index)
    }
}

func DisplayAtackArmes() {
    if Joueur.Vieactuelle > 0 && CurrentAdversery[0].Vieactuelle > 0 {
        var OptionDisplay = []Attaques{}
        var Index int
        for _, Attaque := range Joueur.EquipedWeapon.Attaques {
            OptionDisplay = append(OptionDisplay, Attaque)
        }

        fmt.Print("\n\033[33mQue voulez-vous faire :\033[0m\n")
        for index, OptionsName := range OptionDisplay {
            if Joueur.EquipedWeapon.Mastery < OptionsName.NeededMastery {
                fmt.Printf("\033[33m%d --> \033[31m%v\033[0m \033[31m(maîtrise trop basse)\033[0m\n", index+1, OptionsName.Name)
            } else {
                fmt.Printf("\033[33m%d --> \033[31m%v\033[0m\n", index+1, OptionsName.Name)
            }
            Index++
			if Joueur.EquipedWeapon.Name == PourUnSeulMec.Name[1]{
				fmt.Print(Index+1, PourUnSeulMec.Cap[0].Name)
			}
        }
        fmt.Printf("\033[33m%d --> \033[31mRetour\033[0m\n", Index+2)
    }
}

func IsDead() {
    fmt.Print("\033[31m💀 Vous êtes mort, sale noob 😂\033[0m\n")
    fmt.Printf("\033[33mVous ressuscitez avec \033[31m%d\033[0m PV\n", Joueur.Viemax/2)
    Joueur.Vieactuelle = Joueur.Viemax / 2
    RetourMenu()
}

func Continuer() {
    time.Sleep(1 * time.Second)
    fmt.Print("\033[33m\nAppuyez sur espace puis Entrée pour continuer ...\033[0m\n")
    bufio.NewReader(os.Stdin).ReadBytes(' ')
}

func DisplayPotions() {
    var index = 0
    if Joueur.Vieactuelle > 0 && CurrentAdversery[AdverseryChoice].Vieactuelle > 0 {
        var PotionsDisplay []string
        for _, NamePotion := range AllPotions {
            PotionsDisplay = append(PotionsDisplay, NamePotion.Name)
        }

        fmt.Print("\n\033[33mQuelle potion voulez-vous utiliser ?\033[0m\n")
        for _, PotionsName := range PotionsDisplay {
            fmt.Printf("\033[33m%d --> \033[31m%v\033[0m\n", index+1, PotionsName)
            index++
        }
        fmt.Printf("\033[33m%d --> \033[31mRetour\033[0m\n", index+1)
    }
}

func Fuite() {
    if Joueur.Bourse-CurrentAdversery[AdverseryChoice].GivenMoney > 0 {
        fmt.Printf("\033[31mVous vous êtes fait raquetté pour partir, vous perdez : \033[32m%d 🪙 Techno-Dollars\033[0m\n",
            CurrentAdversery[AdverseryChoice].GivenMoney)
        Joueur.Bourse -= CurrentAdversery[AdverseryChoice].GivenMoney
    } else {
        fmt.Print("\033[31m\nVous fuyez et vous n'avez pas d'argent, vous êtes pitoyable.\033[0m\n")
    }
    if Joueur.EXP-CurrentAdversery[AdverseryChoice].GivenExp > 0 {
        Joueur.EXP -= CurrentAdversery[AdverseryChoice].GivenExp
    } else {
        fmt.Print("\033[31m\nVous fuyez car vous êtes sans expérience, vous êtes pitoyable.\033[0m\n")
    }
    RetourMenu()
}

func MenuQuetes() {
    var SelectionQuete int
    fmt.Print("\033[33mQuelle quête voulez-vous faire ?\033[0m\n")
    fmt.Print("\033[33m1 --> \033[31mBosses\033[0m\n")
    fmt.Print("\033[33m2 --> \033[31mEnnemis Classiques\033[0m\n")
    fmt.Print("\033[33m3 --> Retour :\033[0m\n")
    fmt.Scan(&SelectionQuete)
	switch SelectionQuete {
	case 1 :
		fmt.Print("\033[H\033[2J")
        MenuBosse()
	case 2 :
		fmt.Print("\033[H\033[2J")
        AdverseryChoice = 2
        CombatMode()
	case 3 :
		RetourMenu()
		default : 
			fmt.Print("\033[31mCe n'est pas une option disponible\033[0m\n")
			MenuQuetes()
	}
}

func CombatMode() {
    fmt.Print("\033[H\033[2J")
	PourUnSeulMec.Cap[0].Used = false
    CurrentAdversery[AdverseryChoice].Vieactuelle = CurrentAdversery[AdverseryChoice].Viemax
    var Selection02 int = 0

    for Joueur.Vieactuelle > 0 && CurrentAdversery[AdverseryChoice].Vieactuelle > 0 {
        fmt.Printf("\n\033[33mPV : \033[32m%d\033[0m / \033[32m%d\033[0m    \033[31m%v\033[0m : \033[32m%d\033[0m / \033[32m%d\033[0m\n\n",
            Joueur.Vieactuelle, Joueur.Viemax,
            CurrentAdversery[AdverseryChoice].Nom,
            CurrentAdversery[AdverseryChoice].Vieactuelle,
            CurrentAdversery[AdverseryChoice].Viemax)

        fmt.Print("\033[33mQue voulez-vous faire ?\033[0m\n")
        fmt.Print("\033[33m1 --> \033[31mAttaque\033[0m\n")
        fmt.Print("\033[33m2 --> \033[31mPotions\033[0m\n")
        fmt.Print("\033[33m3 --> \033[31mFuite\033[0m\n")
        fmt.Scan(&Selection02)

        if Selection02 == 1 {
            fmt.Print("\033[33mQuel type d'attaque ?\033[0m\n")
            fmt.Print("\033[33m1 --> \033[31mAttaque Physique\033[0m\n")
            fmt.Print("\033[33m2 --> \033[31mAttaques Arme\033[0m\n")
            fmt.Print("\033[33m3 --> \033[31mRetour\033[0m\n")
            var Selection03 int
            fmt.Scan(&Selection03)

            if Selection03 == 1 {
                DisplayAtackPhysqiue()
                AtackBasiqueSystème()
                Response()
            }
            if Selection03 == 2 {
                DisplayAtackArmes()
                AtackWeaponSystème()
                Response()
            }
            if Selection03 == 3 {
                CombatMode()
            }
            if Selection03 < 0 || Selection03 > 3 {
                fmt.Print("\033[31mOption non disponible\033[0m\n")
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
            fmt.Print("\033[31mCe n'est pas un choix disponible\033[0m\n")
        }
    }
}


var K int // indice pour le retour

func BlackSmithDisplay() {
    fmt.Print("\033[H\033[2J")
    fmt.Print("\033[36m                 Forgeron\033[0m\n")
    var ArmesForgeable = []*Armes{&GantThanos, &StormBreaker}
	var ArmuresForgeable = []*Armures{&ArmureBatman}
	var TeknologiaForgeable = TeknologiaItem
    var IndexObj int
    K = 1

    for index, Name := range ArmesForgeable {
        fmt.Printf("\033[33m%d --> \033[31m%v\033[0m\n", index+1, Name.Name)
        K++
        IndexObj = 0
        for IndexObj < len(ArmesForgeable[index].ObjectsCraft) {
            fmt.Printf("           \033[33m- \033[32m%v\033[0m\n", ArmesForgeable[index].ObjectsCraft[IndexObj].Name)
            IndexObj++
        }
    }
	for index, Name := range ArmuresForgeable {
        fmt.Printf("\033[33m%d --> \033[31m%v\033[0m\n", K, Name.Name)
        K++
        IndexObj = 0
        for IndexObj < len(ArmuresForgeable[index].ObjectCraft) {
            fmt.Printf("           \033[33m- \033[32m%v\033[0m\n", ArmuresForgeable[index].ObjectCraft[IndexObj].Name)
            IndexObj++
        }
    }
	fmt.Print(K, " --> ", TeknologiaForgeable.Name, "\n")
	for IndexObj < len(TeknologiaForgeable.ObjectsCraft) {
            fmt.Printf("           \033[33m- \033[32m%v\033[0m\n", TeknologiaForgeable.ObjectsCraft[IndexObj].Name)
            IndexObj++
        }
    fmt.Printf("\033[33m%d --> \033[31mRetour\033[0m\n", K+1)
}

func DisplayInventory() {
    fmt.Print("\033[H\033[2J")
    var ChoiceSection int
    fmt.Print("\033[36m               Inventaire\033[0m\n")
    fmt.Print("\n\033[33mQue voulez-vous vérifier ?\033[0m\n\n")
    fmt.Print("\033[33m1 --> \033[31mObjets\033[0m\n")
    fmt.Print("\033[33m2 --> \033[31mArmes\033[0m\n")
    fmt.Print("\033[33m3 --> \033[31mArmures\033[0m\n")
    fmt.Print("\033[33m4 --> \033[31mRetour\033[0m\n")
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
        fmt.Print("\033[31mCe n'est pas une option disponible\033[0m\n")
        DisplayInventory()
    }
}

func MenuBosse() {
	var IndexBosses int
	var BosseSelection int
	fmt.Print("\n\nqui voulez vous afronter ?\n\n")
	for indexBosses, Bosses := range AllBosses {
		if Bosses.Nom == "Chimère TeKnologia" && BeatenBosses < 6 {
			fmt.Print(indexBosses+1, " -->", Bosses.Nom, " ", "(battez les autres bosses avant)\n")
			IndexBosses ++
		} else {
			fmt.Print(indexBosses+1, " -->", Bosses.Nom, " ", "(Niveau ", Bosses.Niveau, ")\n")
			IndexBosses ++
		}
		
	}
	fmt.Print(IndexBosses +1, "--> Retour\n")
	fmt.Scan(&BosseSelection)
	switch BosseSelection {
	case 1 :
		AdverseryChoice = 1
		CombatMode()
	case 2 :
		AdverseryChoice = 4
		CombatMode()
	case 3 :
		AdverseryChoice = 5
		CombatMode()
	case 4 :
		AdverseryChoice = 6
		CombatMode()
	case 5 :
		AdverseryChoice = 7
		CombatMode()
	case 6 :
		AdverseryChoice = 8
		CombatMode()
	case 7 :
		if BeatenBosses > 6 {
			AdverseryChoice = 9
			CombatMode()
		} else {
			fmt.Print("\n\nvous ne pouvez pas encore l'afronter\n\n")
			MenuBosse()
		}		
	case 8 :
		fmt.Print("\033[H\033[2J")
		MenuQuetes()
		default : 
		fmt.Print("se n'est pas une option, lâche ça kantin")
		MenuBosse()
	}
}
