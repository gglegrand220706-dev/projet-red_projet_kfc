package personnage

import "fmt"



type Character struct {
	Nom         string
	Classe      string
	vieactuelle int
	Niveau      int
	Viemax       int
	Inventaire   int
	Puissance    int
	Faiblaisse   []string
	Agilite      int
	Intelligence int
	Capacity     []string
}

func CharacterCreation() {
	var joueur Character
	joueur.RecupInfoName()
	var classe Character
	classe.RecupInfoClass()
}

func (u Character) DisplayName() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("votre pseudo est donc: " + u.Nom)
}

func (u *Character) RecupInfoName() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("choisisir un pseudo : ")
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
	for Confirme != 1 {
		ClasseSelection = 0
		fmt.Print("\033[H\033[2J")
		fmt.Println("\n choisisir une classe :\n 1.Kryptonien \n 2.Bat family \n 3.Hulk \n 4.Wakanda")
		fmt.Scan(&ClasseSelection)
		if ClasseSelection == 1 {
			u.Classe = "Kryptonien"
			fmt.Print("\033[H\033[2J")
			fmt.Print("voulez vous choisire ", u.Classe, " comme classe ? \n")
			fmt.Print("Vie max : ", Kryptonien1.Viemax, "\n Inventaire : ", Kryptonien1.Inventaire, "\n Puissance : ", Kryptonien1.Puissance, "\n Faiblaisse : ", Kryptonien1.Faiblaisse, "\n Agilite : ", Kryptonien1.Agilite, "\n Intelligence : ", Kryptonien1.Intelligence, "\n Capacity : ", Kryptonien1.Capacity, "\n")
			fmt.Print("êtes-vous sur ?\n 1. oui \n 2. non \n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = Kryptonien1.Viemax
				u.Inventaire = Kryptonien1.Inventaire
				u.Puissance = Kryptonien1.Puissance
				u.Faiblaisse = Kryptonien1.Faiblaisse
				u.Agilite = Kryptonien1.Agilite
				u.Intelligence = Kryptonien1.Intelligence
				u.Capacity = Kryptonien1.Capacity
				u.DisplayPlayerClass()
			}	
			if	Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}
		
			}

		if ClasseSelection == 2 {
			u.Classe = "Bat family"
			fmt.Print("\033[H\033[2J")
			fmt.Print("voulez vous choisire %v, comme classe ? \n", u.Classe)
			fmt.Print("Vie max : ", BatFamily1.Viemax, "\n Inventaire : ", BatFamily1.Inventaire, "\n Puissance : ", BatFamily1.Puissance, "\n Agilite : ", BatFamily1.Agilite, "\n Intelligence : ", BatFamily1.Intelligence, "\n Capacity : ", BatFamily1.Capacity, "\n")
			fmt.Print("êtes-vous sur ?\n 1. oui \n 2. non\n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = BatFamily1.Viemax
				u.Inventaire = BatFamily1.Inventaire
				u.Puissance = BatFamily1.Puissance
				u.Agilite = BatFamily1.Agilite
				u.Intelligence = BatFamily1.Intelligence
				u.Capacity = BatFamily1.Capacity
				u.DisplayPlayerClass()
			}
			if	Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}
				
			
		}

		if ClasseSelection == 3 {
			u.Classe = "Hulk"
			fmt.Print("\033[H\033[2J")
			fmt.Print("voulez vous choisire %v, comme classe ? \n", u.Classe)
			fmt.Print("Vie max : ", Hulk1.Viemax, "\n Inventaire : ", Hulk1.Inventaire, "\n Puissance : ", Hulk1.Puissance, "\n Faiblaisse : ", Hulk1.Faiblaisse, "\n Agilite : ", Hulk1.Agilite, "\n Intelligence : ", Hulk1.Intelligence, "\n Capacity : ",Hulk1.Capacity, "\n")
			fmt.Print("êtes-vous sur ?\n 1. oui \n 2. non\n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = Hulk1.Viemax
				u.Inventaire = Hulk1.Inventaire
				u.Puissance = Hulk1.Puissance
				u.Faiblaisse = Hulk1.Faiblaisse
				u.Agilite = Hulk1.Agilite
				u.Intelligence = Hulk1.Intelligence
				u.Capacity = Hulk1.Capacity
				u.DisplayPlayerClass()
			}
			if Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}
		}

		if ClasseSelection == 4 {
			u.Classe = "Wakanda"
			fmt.Print("\033[H\033[2J")
			fmt.Print("voulez vous choisire %v, comme classe ? \n", u.Classe)
			fmt.Print("Vie max : ", Wakanda1.Viemax, "\n Inventaire : ", Wakanda1.Inventaire, "\n Puissance : ", Wakanda1.Puissance, "\n Faiblaisse : ", Wakanda1.Faiblaisse, "\n Agilite : ", Wakanda1.Agilite, "\n Intelligence : ", Wakanda1.Intelligence, "\n Capacity : ",Wakanda1.Capacity, "\n")
			fmt.Print("êtes-vous sur ?\n 1. oui \n 2. non\n")
			fmt.Scan(&Confirme)
			if Confirme == 1 {
				u.Viemax = Wakanda1.Viemax
				u.Inventaire = Wakanda1.Inventaire
				u.Puissance = Wakanda1.Puissance
				u.Faiblaisse = Wakanda1.Faiblaisse
				u.Agilite = Wakanda1.Agilite
				u.Intelligence = Wakanda1.Intelligence
				u.Capacity = Wakanda1.Capacity
				u.DisplayPlayerClass()
			}
			if Confirme == 2 {
				fmt.Print("\033[H\033[2J")
			}
				
		}
	}
}

func DisplayInfo(u Character) {
	var joueur Character 
	fmt.Print("infos perso : \n")
	fmt.Print("Pseudo : ", joueur.Nom, "\n Casse : ", joueur.Classe, "\n Niveau : ", joueur.Niveau, "\n Vie : ", joueur.vieactuelle, "/", joueur.Viemax, "\n Compétences : ", joueur.Capacity, "\n Faiblaisses : ", joueur.Faiblaisse  )
}
