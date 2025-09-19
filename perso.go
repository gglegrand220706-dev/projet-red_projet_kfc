package personnage

import "fmt"

type Character struct {
	Nom                   string
	Classe                string
	Vieactuelle           int
	Bourse                int
	Niveau                int
	EXP                   int
	ExpNextLevel          int
	Viemax                int
	Inventaire            int
	Puissance             float64
	Agilite               int
	Intelligence          int
	CapacityClasseDisplay []string
	CapacityDisplay       []string
	Attaques              []Attaques
	Capacity			  []Capacity
	EquipedWeapon		  Armes
	EquipedArmure		  Armures
	AbsorbedDamage		  int
}

var Joueur Character

func CharacterCreation() {
	Joueur.RecupInfoName()
	Joueur.RecupInfoClass()
}

func (u Character) DisplayName() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[33mvotre prénom est donc -->\033[0m" + u.Nom)
}

func (u *Character) RecupInfoName() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[33mChoissisez votre prénom jeune avanturier -->\033[0m")
	fmt.Scan(&u.Nom)
	u.DisplayName()
}

func (u Character) DisplayPlayerClass() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[33mvotre classe est donc --> " + u.Classe + "\033[0m")
}

func (u *Character) RecupInfoClass() {
	var Confirme int = 0
	var ClasseSelection int
	u.Niveau = 1
	u.ExpNextLevel = 10
	u.Attaques = []Attaques{HighKick, GutPunch, CoupDePoing}
	u.CapacityDisplay = []string{CoupDePoing.Name, HighKick.Name, GutPunch.Name}
	for Confirme != 1 {
		ClasseSelection = 0
		fmt.Print("\033[H\033[2J")
		fmt.Println("\033[33m\033[33mChoissiser votre classe :\033[0m\033[0m\n\033[31m\033[0m\033[33m1-->\033[0m\033[34m Kryptonien\033[0m\n\033[31m\033[0m\033[33m2-->\033[0m\033[90m Bat family\033[0m\n\033[31m\033[0m\033[33m3-->\033[0m\033[92m Hulk\n\033[31m\033[0m\033[33m4-->\033[0m\033[95m Wakanda\033[0m")
		fmt.Scan(&ClasseSelection)
		if ClasseSelection == 1 {
            u.Classe = "\033[34m Kryptonien\033[0m"
			fmt.Print("\033[H\033[2J")
			fmt.Print("\033[33mvoulez vous choisir -->", u.Classe, " \033[0m\033[33mcomme classe ?\033[0m\n")
			fmt.Print("\033[33m Vie max -->\0330\033[32m ", Kryptonien1.Viemax, "\033[0m\n \033[33mInventaire -->\033[0m \033[32m ", Kryptonien1.Inventaire, "\033[0m\n \033[33mPuissance -->\033[0m \033[32m ", Kryptonien1.Puissance, "\033[0m\033[0m \n \033[33mAgilite -->\033[0m\033[32m ", Kryptonien1.Agilite, "\033[0m\n\033[33m Intelligence -->\033[0m \033[32m ", Kryptonien1.Intelligence, "\033[0m\n\033[33m Capacity -->\033[0m \033[32m ", Kryptonien1.CapacityDisplay, "\033[0m \n")
			fmt.Print("\033[34mÊtes-vous sur ?\033[0m\n \033[33m1 -->\033[0m \033[32moui\033[0m \n \033[33m2 --> \033[31mnon\033[0m \n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = Kryptonien1.Viemax
				u.Vieactuelle = u.Viemax
				u.Inventaire = Kryptonien1.Inventaire
				u.Puissance = Kryptonien1.Puissance
				u.Agilite = Kryptonien1.Agilite
				u.Intelligence = Kryptonien1.Intelligence
				u.CapacityClasseDisplay = Kryptonien1.CapacityDisplay
				u.Capacity = Kryptonien1.Capacity 
				u.Attaques = append(u.Attaques, Kryptonien1.AttaquesSpé...)
				u.DisplayPlayerClass()
			}
			if Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}

		}

		if ClasseSelection == 2 {
			u.Classe = "\033[90m Bat Family \033[0m"
			fmt.Print("\033[H\033[2J")
			fmt.Print("\033[33mvoulez vous choisir -->", u.Classe, " \033[0m\033[33mcomme classe ?\033[0m\n")
			fmt.Print("\033[33m Vie max -->\0330\033[32m ", BatFamily1.Viemax, "\033[0m\n \033[33mInventaire -->\033[0m \033[32m ", BatFamily1.Inventaire, "\033[0m\n \033[33mPuissance -->\033[0m \033[32m ", BatFamily1.Puissance, "\033[0m \n \033[33mAgilite -->\033[0m\033[32m ", BatFamily1.Agilite, "\033[0m\n\033[33m Intelligence -->\033[0m \033[32m ", BatFamily1.Intelligence, "\033[0m\n\033[33m Capacity -->\033[0m \033[32m ", BatFamily1.CapacityDisplay, "\033[0m \n")
			fmt.Print("\033[90mÊtes-vous sur ?\033[0m\n \033[33m1 -->\033[0m \033[32moui\033[0m \n \033[33m2 --> \033[31mnon\033[0m \n")
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
			if Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}

		}

		if ClasseSelection == 3 {
			u.Classe = "\033[32m Hulk\033[0m"
			fmt.Print("\033[H\033[2J")
			fmt.Print("\033[33mvoulez vous choisir -->", u.Classe, " \033[0m\033[33mcomme classe ?\033[0m\n")
			fmt.Print("\033[33m Vie max -->\0330\033[32m ", Hulk1.Viemax, "\033[0m\n \033[33mInventaire -->\033[0m \033[32m ", Hulk1.Inventaire, "\033[0m\n \033[33mPuissance -->\033[0m \033[32m ", Hulk1.Puissance, "\033[0m \033[0m \n \033[33mAgilite -->\033[0m\033[32m ", Hulk1.Agilite, "\033[0m\n\033[33m Intelligence -->\033[0m \033[32m ", Hulk1.Intelligence, "\033[0m\n\033[33m Capacity -->\033[0m \033[32m ", Hulk1.CapacityDisplay, "\033[0m \n")
			fmt.Print("\033[32mÊtes-vous sur ?\033[0m\n \033[33m1 -->\033[0m \033[32moui\033[0m \n \033[33m2 --> \033[31mnon\033[0m \n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = Hulk1.Viemax
				u.Vieactuelle = u.Viemax
				u.Inventaire = Hulk1.Inventaire
				u.Puissance = Hulk1.Puissance
				u.Agilite = Hulk1.Agilite
				u.Intelligence = Hulk1.Intelligence
				u.CapacityClasseDisplay = Hulk1.CapacityDisplay
				u.Attaques = append(u.Attaques, Hulk1.AttaquesSpé...)
				u.DisplayPlayerClass()
			}
			if Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}
		}

		if ClasseSelection == 4 {
			u.Classe = "\033[95m Wakenda\033[0m"
			fmt.Print("\033[H\033[2J")
			fmt.Print("\033[33mvoulez vous choisir -->", u.Classe, " \033[0m\033[33mcomme classe ?\033[0m\n")
			fmt.Print("\033[33m Vie max -->\0330\033[32m ", Wakanda1.Viemax, "\033[0m\n \033[33mInventaire -->\033[0m \033[32m ", Wakanda1.Inventaire, "\033[0m\n \033[33mPuissance -->\033[0m \033[32m ", Wakanda1.Puissance, "\033[0m\033[0m \n \033[33mAgilite -->\033[0m\033[32m ", Wakanda1.Agilite, "\033[0m\n\033[33m Intelligence -->\033[0m \033[32m ", Wakanda1.Intelligence, "\033[0m\n\033[33m Capacity -->\033[0m \033[32m ", Wakanda1.CapacityDisplay, "\033[0m \n")
			fmt.Print("\033[95mÊtes-vous sur ?\033[0m\n \033[33m1 -->\033[0m \033[32moui\033[0m \n \033[33m2 --> \033[31mnon\033[0m \n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = Wakanda1.Viemax
				u.Vieactuelle = u.Viemax
				u.Inventaire = Wakanda1.Inventaire
				u.Puissance = Wakanda1.Puissance
				u.Agilite = Wakanda1.Agilite
				u.Intelligence = Wakanda1.Intelligence
				u.CapacityClasseDisplay = Wakanda1.CapacityDisplay
				u.Attaques = append(u.Attaques, Wakanda1.AttaquesSpé...)
				u.DisplayPlayerClass()
			}
			if Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}

		}
	}
}

func DisplayInfo() {
	fmt.Print("\033[33minfos perso : \n\033[0m")
fmt.Println("\033[33mPseudo -->\033[32m", Joueur.Nom,
    "\n\033[33mClasse -->\033[32m", Joueur.Classe,
    "\n\033[33mNiveau -->\033[32m", Joueur.Niveau,
    "\n\033[33mExp -->\033[32m", Joueur.EXP, "/", Joueur.ExpNextLevel,
    "\n\033[33mVie -->\033[32m", Joueur.Vieactuelle, "/", Joueur.Viemax,
    "\n\033[33mCompétences -->\033[32m", Joueur.CapacityClasseDisplay,
	)
}

