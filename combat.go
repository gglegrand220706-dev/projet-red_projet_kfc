package personnage

import (
	"fmt"
	"math/rand"
)

func AtackSystème() {
    if Joueur.Vieactuelle > 0 && CurrentAdversery[AdverseryChoice].Vieactuelle > 0{
        var AttackChoice int
        var Attaques = []Attaques{}
        Attaques = append(Attaques, Joueur.Attaques...)
        var RandomeRate int
        RandomeRate = rand.Intn(101)
        fmt.Scan(&AttackChoice)
        if RandomeRate <= int(Attaques[AttackChoice - 1].LandingRate) {
            CurrentAdversery[AdverseryChoice].Vieactuelle -= Attaques[AttackChoice -1].Damage
            fmt.Print("attaques réussis, vie restante : ", CurrentAdversery[AdverseryChoice].Vieactuelle, "/", CurrentAdversery[AdverseryChoice].Viemax, " pv\n")
        } else {
            fmt.Print("attaque ratée \n")
        }    
    }
}

func Response() {
    if CurrentAdversery[AdverseryChoice].Vieactuelle > 0 && Joueur.Vieactuelle > 0 {
        var ChoiceResponse int
        var Attaques = []Attaques{}
        Attaques = append(Attaques, CurrentAdversery[AdverseryChoice].Attaques...)
        var UsedAttaques string
        var RandomeRate int
        ChoiceResponse = rand.Intn(4)
        RandomeRate = rand.Intn(101)
        if RandomeRate <= int(Attaques[ChoiceResponse].LandingRate) {
            Joueur.Vieactuelle -= Attaques[ChoiceResponse].Damage
            UsedAttaques = Attaques[ChoiceResponse].Name
            fmt.Print("vous avez été touché par : ", UsedAttaques, "\n")
        }else {
            fmt.Print("son attaque a été râté petit chanceux")
        }
    }
}

var CurrentAdversery = []Adversery{Adversery01, AdverseryLoki}
var AdverseryChoice int