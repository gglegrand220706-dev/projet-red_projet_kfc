package personnage

import "fmt"



type Character struct {
	Nom         		string
	Classe      		string
	Vieactuelle 		int
	Bourse 				int
	Niveau      		int
	EXP					int
	ExpNextLevel 		int	
	Viemax       		int
	Inventaire   		int
	Puissance    		int
	Faiblaisse   		[]string
	Agilite      		int
	Intelligence        int
	CapacityClasseDisplay     []string
	CapacityDisplay     []string
	Capacity            []Attaques
}

var Joueur Character

func CharacterCreation() {
	Joueur.RecupInfoName()
	Joueur.RecupInfoClass()
}

func (u Character) DisplayName() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[33mvotre pseudo est donc -->\033[0m" + u.Nom)
}

func (u *Character) RecupInfoName() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[33mChoissiser votre pseudo -->\033[0m")
	fmt.Scan(&u.Nom)
	u.DisplayName()
}

func (u Character) DisplayPlayerClass() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("votre classe est donc : " + u.Classe)
}

func (u *Character) RecupInfoClass() {
	var Confirme int = 0
	var ClasseSelection int
	u.Capacity = []Attaques{HighKick, GutPunch, CoupDePoing}
	u.CapacityDisplay = []string{CoupDePoing.Name, HighKick.Name, GutPunch.Name}
	for Confirme != 1 {
		ClasseSelection = 0
		fmt.Print("\033[H\033[2J")
		fmt.Println("\033[33m\033[33mChoissiser votre classe :\033[0m\033[0m\n\033[31m\033[0m\033[33m1-->\033[0m\033[34m Kryptonien\033[0m\n\033[31m\033[0m\033[33m2-->\033[0m\033[90m Bat family\033[0m\n\033[31m\033[0m\033[33m3-->\033[0m\033[92m Hulk\n\033[31m\033[0m\033[33m4-->\033[0m\033[95m Wakanda\033[0m")
		fmt.Scan(&ClasseSelection)
		if ClasseSelection == 1 {
			u.Classe = "\033[34m Kryptonien\033[0m"
			fmt.Print("\033[H\033[2J")
			fmt.Print("\033[33mvoulez vous choisir --> ",u.Classe," \033[0m comme classe ? \n")
			fmt.Print("\033[33m Vie max -->\0330\033[32m ",Kryptonien1.Viemax,"\033[0m\n \033[33mInventaire -->\033[0m \033[32m ",Kryptonien1.Inventaire,"\033[0m\n \033[33mPuissance -->\033[0m \033[32m ",Kryptonien1.Puissance,"\033[0m \n \033[33mFaiblaisse -->\033[0m\033[32m ",Kryptonien1.Faiblaisse,"\033[0m \n \033[33mAgilite -->\033[0m\033[32m ",Kryptonien1.Agilite,"\033[0m\n\033[33m Intelligence -->\033[0m \033[32m ",Kryptonien1.Intelligence,"\033[0m\n\033[33m Capacity -->\033[0m \033[32m ",Kryptonien1.CapacityDisplay,"\033[0m \n")
			fmt.Print("êtes-vous sur ?\n 1. oui \n 2. non \n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = Kryptonien1.Viemax
				u.Vieactuelle = u.Viemax
				u.Inventaire = Kryptonien1.Inventaire
				u.Puissance = Kryptonien1.Puissance
				u.Faiblaisse = Kryptonien1.Faiblaisse
				u.Agilite = Kryptonien1.Agilite
				u.Intelligence = Kryptonien1.Intelligence
				u.CapacityClasseDisplay = Kryptonien1.CapacityDisplay
				u.DisplayPlayerClass()
			}	
			if	Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}
		
			}

		if ClasseSelection == 2 {
			u.Classe = "Bat family"
			fmt.Print("\033[H\033[2J")
			fmt.Print("voulez vous choisire ", u.Classe, " comme classe ? \n")
			fmt.Print("Vie max : ", BatFamily1.Viemax, "\n Inventaire : ", BatFamily1.Inventaire, "\n Puissance : ", BatFamily1.Puissance, "\n Agilite : ", BatFamily1.Agilite, "\n Intelligence : ", BatFamily1.Intelligence, "\n Capacity : ", BatFamily1.CapacityDisplay, "\n")
			fmt.Print("êtes-vous sur ?\n 1. oui \n 2. non\n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = BatFamily1.Viemax
				u.Vieactuelle = u.Viemax
				u.Bourse = 2000
				u.Inventaire = BatFamily1.Inventaire
				u.Puissance = BatFamily1.Puissance
				u.Agilite = BatFamily1.Agilite
				u.Intelligence = BatFamily1.Intelligence
				u.CapacityClasseDisplay = BatFamily1.CapacityDisplay
				u.DisplayPlayerClass()
			}
			if	Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}
				
			
		}

		if ClasseSelection == 3 {
			u.Classe = "Hulk"
			fmt.Print("\033[H\033[2J")
			fmt.Print("voulez vous choisire ", u.Classe, " comme classe ? \n", u.Classe)
			fmt.Print("Vie max : ", Hulk1.Viemax, "\n Inventaire : ", Hulk1.Inventaire, "\n Puissance : ", Hulk1.Puissance, "\n Faiblaisse : ", Hulk1.Faiblaisse, "\n Agilite : ", Hulk1.Agilite, "\n Intelligence : ", Hulk1.Intelligence, "\n Capacity : ",Hulk1.CapacityDisplay, "\n")
			fmt.Print("êtes-vous sur ?\n 1. oui \n 2. non\n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = Hulk1.Viemax
				u.Vieactuelle = u.Viemax
				u.Inventaire = Hulk1.Inventaire
				u.Puissance = Hulk1.Puissance
				u.Faiblaisse = Hulk1.Faiblaisse
				u.Agilite = Hulk1.Agilite
				u.Intelligence = Hulk1.Intelligence
				u.CapacityClasseDisplay = Hulk1.CapacityDisplay
				u.DisplayPlayerClass()
			}
			if Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}
		}

		if ClasseSelection == 4 {
			u.Classe = "Wakanda"
			fmt.Print("\033[H\033[2J")
			fmt.Print("voulez vous choisire ", u.Classe, " comme classe ? \n", u.Classe)
			fmt.Print("Vie max : ", Wakanda1.Viemax, "\n Inventaire : ", Wakanda1.Inventaire, "\n Puissance : ", Wakanda1.Puissance, "\n Faiblaisse : ", Wakanda1.Faiblaisse, "\n Agilite : ", Wakanda1.Agilite, "\n Intelligence : ", Wakanda1.Intelligence, "\n Capacity : ",Wakanda1.CapacityDisplay, "\n")
			fmt.Print("êtes-vous sur ?\n 1. oui \n 2. non\n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = Wakanda1.Viemax
				u.Vieactuelle = u.Viemax
				u.Inventaire = Wakanda1.Inventaire
				u.Puissance = Wakanda1.Puissance
				u.Faiblaisse = Wakanda1.Faiblaisse
				u.Agilite = Wakanda1.Agilite
				u.Intelligence = Wakanda1.Intelligence
				u.CapacityClasseDisplay = Wakanda1.CapacityDisplay
				u.DisplayPlayerClass()
			}
			if Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}
				
		}
	}
}

func DisplayInfo() {
	fmt.Print("infos perso : \n")
	fmt.Print("Pseudo : ", Joueur.Nom, "\n Casse : ", Joueur.Classe, "\n Niveau : ", Joueur.Niveau, "\n Exp : ", Joueur.EXP, "/", Joueur.ExpNextLevel, "\n Vie : ", Joueur.Vieactuelle, "/", Joueur.Viemax, "\n Compétences : ", Joueur.CapacityClasseDisplay, "\n Faiblaisses : ", Joueur.Faiblaisse)
}
