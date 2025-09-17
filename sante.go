package personnage

import (
	"fmt"
)

func PotionsVie() {
	var Healed bool
	if PotionsSoins.Nb > PotionsSoins.EffectLife {
		if Joueur.Vieactuelle == Joueur.Viemax && !Healed{
			fmt.Print("vous ne pouvez pas utilisé cette potions")
			Healed =  true
		}		
		if Joueur.Viemax - Joueur.Vieactuelle >= PotionsSoins.EffectLife && !Healed{
			Joueur.Vieactuelle += PotionsSoins.EffectLife
			PotionsSoins.Nb --
			fmt.Print("Potion de soins utiliser, + ", PotionsSoins.EffectLife, " PV\n")
			Healed = true
		}
		if Joueur.Viemax - Joueur.Vieactuelle < PotionsSoins.EffectLife && !Healed{
			Joueur.Vieactuelle = Joueur.Viemax
			PotionsSoins.Nb --
			fmt.Print("Potions de soins utilisée, + ", Joueur.Viemax - Joueur.Vieactuelle, " PV\n" )
			Healed = true
		}
	} else {
		fmt.Print("vous n'avez plus de potions")
	}
	Healed = false
}

var IsPoison bool = false

func PotionDamage(){
	IsPoison = true
}

func PotionCHoice() {
	var PotionsChoice int
	fmt.Scan(&PotionsChoice)
	switch PotionsChoice {
	case 1 :
		PotionsVie()
	case 2 :
		PotionDamage()
	case 3 :
		CombatMode()
	default :
		fmt.Print("\nSe n'est pas une option de potion")
		PotionCHoice()
	}
}
