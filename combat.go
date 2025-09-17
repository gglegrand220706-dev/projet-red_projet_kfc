package personnage

import (
	"fmt"
	"math/rand"
    "math"
)

func AtackSystème() {
    if Joueur.Vieactuelle > 0 && CurrentAdversery[AdverseryChoice].Vieactuelle > 0{
        var AttackChoice int
        var Multiplicateur float64 
        var Attaques = []Attaques{}
        Attaques = append(Attaques, Joueur.Attaques...)
        var RandomeRate int
        RandomeRate = rand.Intn(101)
        fmt.Scan(&AttackChoice)
        if AttackChoice - 1 > len(Attaques) +1 {
            fmt.Print("ce n'est pas une proposition espèce de demeuré")
            AtackSystème()
        }        
        if AttackChoice == len(Attaques) + 1 {
            PotionsVie()
        }
        if AttackChoice == 5{
            CombatMode()
        }
        if AttackChoice - 1 < len(Attaques) {
            if RandomeRate <= int(Attaques[AttackChoice - 1].LandingRate) {       
                Multiplicateur = (1 + Joueur.Puissance/100)
               CurrentAdversery[AdverseryChoice].Vieactuelle -= Attaques[AttackChoice -1].Damage * Arrondir(Multiplicateur) 
               if CurrentAdversery[AdverseryChoice].Vieactuelle < 0 {
                CurrentAdversery[AdverseryChoice].Vieactuelle = 0
               } 
                fmt.Print("votre attaque a réussi, votre adversaire a toujours ", CurrentAdversery[AdverseryChoice].Vieactuelle, "/", CurrentAdversery[AdverseryChoice].Viemax, " PV\n")
            } else {
               fmt.Print("votre attaque a raté \n")
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
            fmt.Print("l'attaque de", CurrentAdversery[AdverseryChoice].Nom, " a raté, petit chanceux")
        }
    }
}

var CurrentAdversery = []*Adversery{&Adversery01, &AdverseryLoki, &AdverseryBarry}
var AdverseryChoice int

func Arrondir(Multiplicateur float64) int {
    entier := math.Floor(Multiplicateur)
    decimal := Multiplicateur - entier
    if decimal > 0.5 {
        return int(entier) + 1
    } else if decimal < 0.5 {
        return int(entier)
    } else { 
        return int(entier) + 1
    }
}

func DropRate() {
    var RandomeDrop int
    var RandomeObj int
    if CurrentAdversery[AdverseryChoice].Vieactuelle == 0 {
        RandomeDrop = rand.Intn(101)
        RandomeObj = rand.Intn(len(CurrentAdversery[AdverseryChoice].Drop))
        if RandomeDrop <= CurrentAdversery[AdverseryChoice].Drop[RandomeObj].DropRate && CurrentAdversery[AdverseryChoice].Drop[RandomeObj].MaxAmount > CurrentAdversery[AdverseryChoice].Drop[RandomeObj].Nb {
            CurrentAdversery[AdverseryChoice].Drop[RandomeObj].Nb ++
            fmt.Print("vous avez reçu ", CurrentAdversery[AdverseryChoice].Drop[RandomeObj].Name, "\n")
        } else {
            fmt.Print("vous avez déjà le max de ", CurrentAdversery[AdverseryChoice].Drop[RandomeObj].Name, "\n")
        }
    }
    if P == 3{
        for P > 0{
            CurrentAdversery[AdverseryChoice].Vieactuelle-=5
            fmt.Print("", CurrentAdversery[AdverseryChoice].Nom, " subit des dégâts de poison, il lui reste ", CurrentAdversery[AdverseryChoice].Vieactuelle, "/", CurrentAdversery[AdverseryChoice].Vieactuelle, " PV.\n")
}
    }
        
}
    