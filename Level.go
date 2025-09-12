package personnage

import "fmt"

func AddExp() {
	if Joueur.Classe == "Bat family" {
		Joueur.EXP += 2*Adversery01.GivenExp
	} else {
		Joueur.EXP += Adversery01.GivenExp
	}	
	fmt.Print("\nvous avez gagné : ", Adversery01.GivenExp, " EXP")
	if Joueur.EXP >= Joueur.ExpNextLevel {
		Joueur.Niveau += 1
		Joueur.ExpNextLevel = Joueur.Niveau + (10 - Joueur.Niveau) + 10
		fmt.Print("\nVous passez niveau ", Joueur.Niveau)
		Joueur.EXP = 0
	}
}
