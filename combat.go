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
        if AttackChoice - 1 < len(Attaques){
            fmt.Print("se n'est pas une proposition espèce de demeuré")
            AtackSystème()
        }        
        if AttackChoice == len(Attaques) + 1 {
            PotionsVie()
        }
        if AttackChoice - 1 < len(Attaques){
            if RandomeRate <= int(Attaques[AttackChoice - 1].LandingRate) {
               CurrentAdversery[AdverseryChoice].Vieactuelle -= int(Attaques[AttackChoice -1].Damage * (1 + Joueur.Puissance/100))
               if CurrentAdversery[AdverseryChoice].Vieactuelle < 0 {
                CurrentAdversery[AttackChoice].Vieactuelle = 0
                fmt.Print("attaques réussis, vie restante : ", CurrentAdversery[AdverseryChoice].Vieactuelle, "/", CurrentAdversery[AdverseryChoice].Viemax, " PV\n")
               } else {
                    fmt.Print("attaques réussis, vie restante : ", CurrentAdversery[AdverseryChoice].Vieactuelle, "/", CurrentAdversery[AdverseryChoice].Viemax, " PV\n")
               } 
            } else {
               fmt.Print("attaque ratée \n")
            }
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
        ChoiceResponse = rand.Intn(len(CurrentAdversery[AdverseryChoice].Attaques))
        RandomeRate = rand.Intn(101)
        if RandomeRate <= int(Attaques[ChoiceResponse].LandingRate) {
            Joueur.Vieactuelle -= int(Attaques[ChoiceResponse].Damage * (1 + CurrentAdversery[AdverseryChoice].Puissance/100))
            UsedAttaques = Attaques[ChoiceResponse].Name
            fmt.Print("vous avez été touché par : ", UsedAttaques, "\n")
        }else {
            fmt.Print("son attaque a été râté petit chanceux")
        }
    }
}

var CurrentAdversery = []*Adversery{&Adversery01, &AdverseryLoki, &AdverseryBarry}
var AdverseryChoice int
