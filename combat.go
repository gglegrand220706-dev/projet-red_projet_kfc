package personnage

import (
	"fmt"
	"math/rand"
)

func AtackSystème() {
    if Joueur.Vieactuelle > 0 && Adversery01.Vieactuelle > 0{
        var AttackChoice int
        var Attaques = []Attaques{}
        Attaques = append(Attaques, Joueur.Attaques...)
        var RandomeRate int
        RandomeRate = rand.Intn(101)
        fmt.Scan(&AttackChoice)
        if RandomeRate <= int(Attaques[AttackChoice - 1].LandingRate) {
            Adversery01.Vieactuelle -= Attaques[AttackChoice -1].Damage
            fmt.Print("attaques réussis, vie restante : ", Adversery01.Vieactuelle, "/", Adversery01.Viemax, " pv\n")
        } else {
            fmt.Print("attaque ratée \n")
        }    
    }
}

func Response() {
    if Adversery01.Vieactuelle > 0 && Joueur.Vieactuelle > 0 {
        var ChoiceResponse int
        var Attaques = []Attaques{AttaqueBasique01, AttaqueBasique02, AttaqueBasique03, AttaqueBasique04}
        var UsedAttaques string
        var RandomeRate int
        ChoiceResponse = rand.Intn(4)
        RandomeRate = rand.Intn(101)
        if RandomeRate <= int(Attaques[ChoiceResponse].LandingRate) {
            Joueur.Vieactuelle -= Attaques[ChoiceResponse].Damage
            UsedAttaques = Attaques[ChoiceResponse].Name
            fmt.Print("Attaque utilisée : ", UsedAttaques, "\n")
            fmt.Print("vous avez été touché, il vous reste : ", Joueur.Vieactuelle, "/", Joueur.Viemax, " pv\n")
        }else {
            fmt.Print("son attaque a été râté petit chanceux")
        }
    }
}

