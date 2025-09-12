package personnage

import (
	"fmt"
	"math/rand"
)

func AtackSystème() {
    if Joueur.Vieactuelle > 0 && Adversery01.Vieactuelle > 0{
        var AttackChoice int
        var Damges = []int{CoupDePoing.Damage, HighKick.Damage, GutPunch.Damage, AttaqueLasser.Damage}
        fmt.Scan(&AttackChoice)
        Adversery01.Vieactuelle -= Damges[AttackChoice -1]
        fmt.Print("attaques réussis, vie restante : ", Adversery01.Vieactuelle, "/", Adversery01.Viemax, " pv\n")
    }
}

func Response() {
    if Adversery01.Vieactuelle > 0 && Joueur.Vieactuelle > 0 {
        var ChoiceResponse int
        var Damages = []int{AttaqueBasique01.Damage, AttaqueBasique02.Damage, AttaqueBasique03.Damage}
        ChoiceResponse = rand.Intn(3)
        Joueur.Vieactuelle -= Damages[ChoiceResponse]
        fmt.Print("vous avez été touché, il vous reste : ", Joueur.Vieactuelle, "/", Joueur.Viemax, " pv\n")
    }
}

