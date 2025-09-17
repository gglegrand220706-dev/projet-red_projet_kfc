package personnage

import "fmt"

func Reward() {
    GainedExp := CurrentAdversery[AdverseryChoice].GivenExp
    GainedMoney := CurrentAdversery[AdverseryChoice].GivenMoney
    if Joueur.Classe == "Bat family" {
        GainedExp *= 2
    }
    Joueur.EXP += GainedExp
    Joueur.Bourse += GainedMoney
    fmt.Printf("\nVous avez gagné : %d EXP et : %d TeKno-Dollars\n", GainedExp, GainedMoney)
    for Joueur.EXP >= Joueur.ExpNextLevel {
        Joueur.EXP -= Joueur.ExpNextLevel
        NextLevel()
        Joueur.ExpNextLevel = Joueur.Niveau*10 + 20
        fmt.Printf("\nVous passez niveau %d\n", Joueur.Niveau)
        if Joueur.EquipedWeapon.MaxMastery > Joueur.EquipedWeapon.Mastery {
            Joueur.EquipedWeapon.Mastery ++
        } else {
            fmt.Print("votre mastery est au mex sur cette arme")
        }
        
    }
}

func NextLevel() {
	Joueur.Niveau ++
	Joueur.Intelligence += 5
	Joueur.Puissance += 5
	Joueur.Viemax += 5
	Joueur.Vieactuelle = Joueur.Viemax
	Joueur.Agilite += 5
}