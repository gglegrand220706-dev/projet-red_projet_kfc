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

    fmt.Printf("\n\033[33mVous avez gagné : \033[32m%d ✨ EXP\033[32m Et \033[32m%d 🪙 TeKno-Dirham\033[0m\n",
        GainedExp, GainedMoney)

    for Joueur.EXP >= Joueur.ExpNextLevel {
        Joueur.EXP -= Joueur.ExpNextLevel
        NextLevel()
        Joueur.ExpNextLevel = Joueur.Niveau*10 + 20
        fmt.Printf("\n\033[33mVous passez niveau \033[32m%d 🆙\033[0m\n", Joueur.Niveau)

        if Joueur.EquipedWeapon.MaxMastery > Joueur.EquipedWeapon.Mastery {
            Joueur.EquipedWeapon.Mastery++
        } else {
            fmt.Print("\033[31m✅ Votre mastery est au max sur cette arme ! 🗡️🔥\033[0m\n")
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