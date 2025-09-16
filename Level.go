package personnage

import "fmt"

func Reward() {
	var RestantExp int
	if Joueur.Classe == "Bat family" {
		Joueur.EXP += 2*CurrentAdversery[AdverseryChoice].GivenExp
	} else {
		Joueur.EXP += CurrentAdversery[AdverseryChoice].GivenExp
	}	
	fmt.Print("\nvous avez gagné : ", CurrentAdversery[AdverseryChoice].GivenExp, " EXP et : ", CurrentAdversery[AdverseryChoice].GivenMoney, " techno-Dollars\n")
	Joueur.Bourse += CurrentAdversery[AdverseryChoice].GivenMoney
	if Joueur.ExpNextLevel - Joueur.EXP < CurrentAdversery[AdverseryChoice].GivenExp {
		RestantExp = CurrentAdversery[AdverseryChoice].GivenExp - Joueur.ExpNextLevel - Joueur.EXP
		Joueur.Niveau ++
		Joueur.EXP = RestantExp
	}
	if Joueur.EXP == Joueur.ExpNextLevel {
		Joueur.Niveau += 1
		Joueur.ExpNextLevel = Joueur.Niveau*10 +20
		fmt.Print("\nVous passez niveau ", Joueur.Niveau)
		Joueur.EXP = 0
	}
}
