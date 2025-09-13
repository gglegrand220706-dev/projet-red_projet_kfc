package personnage

import (
	"fmt"
	"math/rand"
)

func AtackSystème() {
    if Joueur.Vieactuelle > 0 && Adversery01.Vieactuelle > 0{
        var AttackChoice int
        var Damges = []int{CoupDePoing.Damage, HighKick.Damage, GutPunch.Damage, AttaqueLasser.Damage}
        var LandingRate =[]int{int(CoupDePoing.LandingRate), int(HighKick.LandingRate), int(GutPunch.LandingRate), int(AttaqueLasser.LandingRate)}
        var RandomeRate int
        RandomeRate = rand.Intn(101)
        fmt.Scan(&AttackChoice)
        if RandomeRate <= LandingRate[AttackChoice - 1] {
            Adversery01.Vieactuelle -= Damges[AttackChoice -1]
            fmt.Print("attaques réussis, vie restante : ", Adversery01.Vieactuelle, "/", Adversery01.Viemax, " pv\n")
        } else {
            fmt.Print("attaque ratée \n")
        }
            
    }
}

func Response() {
    if Adversery01.Vieactuelle > 0 && Joueur.Vieactuelle > 0 {
        var ChoiceResponse int
        var Damages = []int{AttaqueBasique01.Damage, AttaqueBasique02.Damage, AttaqueBasique03.Damage, AttaqueBasique04.Damage}
        var NameAttaques = []string{AttaqueBasique01.Name, AttaqueBasique02.Name, AttaqueBasique03.Name, AttaqueBasique04.Name}
        var UsedAttaques string
        var LandingRate = []int{int(AttaqueBasique01.LandingRate), int(AttaqueBasique02.LandingRate), int(AttaqueBasique03.LandingRate), int(AttaqueBasique04.LandingRate)}
        var RandomeRate int
        ChoiceResponse = rand.Intn(4)
        RandomeRate = rand.Intn(101)
        if RandomeRate <= LandingRate[ChoiceResponse] {
            Joueur.Vieactuelle -= Damages[ChoiceResponse]
            UsedAttaques = NameAttaques[ChoiceResponse]
            fmt.Print("Attaque utilisée : ", UsedAttaques, "\n")
            fmt.Print("vous avez été touché, il vous reste : ", Joueur.Vieactuelle, "/", Joueur.Viemax, " pv\n")
        }else {
            fmt.Print("son attaque a été râté petit chanceux")
        }
    }
}

