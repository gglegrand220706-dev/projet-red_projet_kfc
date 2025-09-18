package personnage

import (
	"fmt"
)

func PotionsVie() {
    var Healed bool
    if PotionsSoins.Nb > 0 {
        if Joueur.Vieactuelle == Joueur.Viemax && !Healed {
            fmt.Print("\033[31mVous ne pouvez pas utiliser cette potion\033[0m\n")
            Healed = true
        }
        if Joueur.Viemax-Joueur.Vieactuelle >= PotionsSoins.EffectLife && !Healed {
            Joueur.Vieactuelle += PotionsSoins.EffectLife
            PotionsSoins.Nb--
            fmt.Printf("\033[32mPotion de soins utilisée, +%d ❤️ PV\033[0m\n", PotionsSoins.EffectLife)
            Healed = true
        }
        if Joueur.Viemax-Joueur.Vieactuelle < PotionsSoins.EffectLife && !Healed {
            fmt.Printf("\033[32mPotion de soins utilisée, +%d ❤️ PV\033[0m\n", Joueur.Viemax-Joueur.Vieactuelle)
            Joueur.Vieactuelle = Joueur.Viemax
            PotionsSoins.Nb--
            Healed = true
        }
    } else {
        fmt.Print("\033[31mVous n'avez plus de potions\033[0m\n")
    }
    Healed = false
}

var IsPoison bool = false

func PotionCHoice() {
    var PotionsChoice int
    fmt.Scan(&PotionsChoice)
    switch PotionsChoice {
    case 1:
        PotionsVie()
    case 2:
        if AllPotions[1].Nb > 0 {
            IsPoison = true
        } else {
            fmt.Print("\033[31mVous n'avez pas cette potion\033[0m\n")
            DisplayPotions()
        }
    case 3:
        CombatMode()
    default:
        fmt.Print("\033[31mCe n'est pas une option de potion\033[0m\n")
        PotionCHoice()
    }
}

