package personnage

import (
	"fmt"
	"math/rand"
    "math"
)

func AtackBasiqueSystème() {
    if Joueur.Vieactuelle > 0 && CurrentAdversery[AdverseryChoice].Vieactuelle > 0 {
        var AttackChoice int
        var Multiplicateur float64
        var Attaques = []Attaques{}
        Attaques = append(Attaques, Joueur.Attaques...)
        var RandomeRate int
        RandomeRate = rand.Intn(101)
        fmt.Scan(&AttackChoice)
        if AttackChoice-1 > len(Attaques)+1 {
            fmt.Print("\033[31m❌ Ce n'est pas une proposition espèce de demeuré\033[0m\n")
            AtackBasiqueSystème()
            return
        }
          if AttackChoice == 5 {
            CombatMode()
            return
        }
        if AttackChoice == len(Attaques)+2 {
            PotionsVie()
            return
        }
      
        if AttackChoice-1 < len(Attaques) {
            if Attaques[AttackChoice-1].Name == "Renvoie Sinetqiue" {
                if Joueur.AbsorbedDamage >= 30 {
                    CurrentAdversery[AdverseryChoice].Vieactuelle -= 30
                    Joueur.AbsorbedDamage = 0 
                    fmt.Printf("\033[33mPV adversaire : \033[31m%d\033[0m/\033[32m%d\033[0m\n",
                    CurrentAdversery[AdverseryChoice].Vieactuelle,
                    CurrentAdversery[AdverseryChoice].Viemax)
                    return
                } else {
                    fmt.Print("vous ne pouvez pas faire cette attaque\n")
                    return
                }
            }
            if RandomeRate <= int(Attaques[AttackChoice-1].LandingRate) {
                Multiplicateur = (1 + Joueur.Puissance/100 + float64(Joueur.EquipedWeapon.Multiplicateur)/100)
                CurrentAdversery[AdverseryChoice].Vieactuelle -= Attaques[AttackChoice-1].Damage * Arrondir(Multiplicateur)
                if CurrentAdversery[AdverseryChoice].Vieactuelle < 0 {
                    CurrentAdversery[AdverseryChoice].Vieactuelle = 0
                }
                fmt.Print("\033[32m✅ Votre attaque a réussi !\033[0m\n")
                fmt.Printf("\033[33mPV adversaire : \033[31m%d\033[0m/\033[32m%d\033[0m\n",
                    CurrentAdversery[AdverseryChoice].Vieactuelle,
                    CurrentAdversery[AdverseryChoice].Viemax)
            } else {
                fmt.Print("\033[31m❌ Votre attaque a raté !\033[0m\n")
            }
        }
    }
}

var BeatenBosses int

func Response() {
    var Count int = 2
    if CurrentAdversery[AdverseryChoice].Vieactuelle > 0 && Joueur.Vieactuelle > 0 {
        if IsPoison {
            if Count != 0 {
                Count--
                CurrentAdversery[AdverseryChoice].Vieactuelle -= 5
                fmt.Printf("\033[35m☠️ %v est affecté par le poison\033[0m\n", CurrentAdversery[AdverseryChoice].Nom)
                fmt.Printf("\033[33mPV adversaire : \033[31m%d\033[0m/\033[32m%d\033[0m\n",
                    CurrentAdversery[AdverseryChoice].Vieactuelle,
                    CurrentAdversery[AdverseryChoice].Viemax)

                if CurrentAdversery[AdverseryChoice].Vieactuelle <= 0 {
                    fmt.Printf("\033[32m🏆 Vous avez vaincu %v !\033[0m\n", CurrentAdversery[AdverseryChoice].Nom)
                    if CurrentAdversery[AdverseryChoice].Bosse {
                        CurrentAdversery[AdverseryChoice].Bosse = false
                        BeatenBosses ++
                    }
                    Reward()
                }
            } else {
                IsPoison = false
            }
        }
    }
    if PourUnSeulMec.Name[0] == CurrentAdversery[AdverseryChoice].Nom {
        var Attaque_CapaityChoice int
        var UsedAttaques_Capacity string
        Attaque_CapaityChoice = rand.Intn(4)
        if Attaque_CapaityChoice + 1 > len(PourUnSeulMec.Att) && !PourUnSeulMec.Cap[0].Used {
            PourUnSeulMec.Cap[0].Used = true
            var RandomeRate int
                RandomeRate = int(rand.Intn(101*(1-Joueur.Agilite/100)))
                if RandomeRate <= int(PourUnSeulMec.Att[Attaque_CapaityChoice].LandingRate){
                    Joueur.Vieactuelle -= int(PourUnSeulMec.Att[Attaque_CapaityChoice].Damage * (1 + CurrentAdversery[AdverseryChoice].Puissance/100))
                    Joueur.AbsorbedDamage += int(PourUnSeulMec.Att[Attaque_CapaityChoice].Damage * (1 + CurrentAdversery[AdverseryChoice].Puissance/100))    
                    UsedAttaques_Capacity = PourUnSeulMec.Att[Attaque_CapaityChoice].Name
                    fmt.Printf("\n\033[31m❌ Vous avez été touché par : %v\033[0m\n", UsedAttaques_Capacity)
                    fmt.Printf("\033[33mPV joueur : \033[31m%d\033[0m/\033[32m%d\033[0m\n",
                    Joueur.Vieactuelle, Joueur.Viemax)
                }
        } else {
             if CurrentAdversery[AdverseryChoice].Vieactuelle > 0 && Joueur.Vieactuelle > 0 {
                var RandomeRate int
                RandomeRate = rand.Intn(101)
                if RandomeRate <= int(PourUnSeulMec.Att[Attaque_CapaityChoice].LandingRate){
                    Joueur.Vieactuelle -= int(PourUnSeulMec.Att[Attaque_CapaityChoice].Damage * (1 + CurrentAdversery[AdverseryChoice].Puissance/100))
                    Joueur.AbsorbedDamage += int(PourUnSeulMec.Att[Attaque_CapaityChoice].Damage * (1 + CurrentAdversery[AdverseryChoice].Puissance/100))
                    UsedAttaques_Capacity = PourUnSeulMec.Att[Attaque_CapaityChoice].Name
                    fmt.Printf("\n\033[31m❌ Vous avez été touché par : %v\033[0m\n", UsedAttaques_Capacity)
                    fmt.Printf("\033[33mPV joueur : \033[31m%d\033[0m/\033[32m%d\033[0m\n",
                    Joueur.Vieactuelle, Joueur.Viemax)
                }
             }
        }
    }
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
            Joueur.AbsorbedDamage +=int(PourUnSeulMec.Att[AdverseryChoice].Damage * (1 + CurrentAdversery[AdverseryChoice].Puissance/100))
            UsedAttaques = Attaques[ChoiceResponse].Name
            fmt.Printf("\033[31m❌ Vous avez été touché par : %v\033[0m\n", UsedAttaques)
            fmt.Printf("\033[33mPV joueur : \033[31m%d\033[0m/\033[32m%d\033[0m\n",
                Joueur.Vieactuelle, Joueur.Viemax)
        } else {
            fmt.Printf("\033[32m✅ L'attaque de %v a raté, petit chanceux !\033[0m\n", CurrentAdversery[AdverseryChoice].Nom)
        }
    }
}
    

var CurrentAdversery = []*Adversery{&Adversery01, &AdverseryLoki, &AdverseryBarry, &AdverseryRicooo, &AdverseryWanda, &AdverseryJoker, &Dormammu, &Ronan, &Darkseid, &ChimereTK}
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
        if RandomeDrop <= CurrentAdversery[AdverseryChoice].Drop[RandomeObj].DropRate &&
            CurrentAdversery[AdverseryChoice].Drop[RandomeObj].MaxAmount > CurrentAdversery[AdverseryChoice].Drop[RandomeObj].Nb {

            CurrentAdversery[AdverseryChoice].Drop[RandomeObj].Nb++
            fmt.Printf("\033[32m🎁 Vous avez reçu \033[31m%v\033[0m\n",
                CurrentAdversery[AdverseryChoice].Drop[RandomeObj].Name)

        } else {
            fmt.Printf("\033[31m❌ Vous avez déjà le maximum de \033[0m\033[31m%v\033[0m\n",
                CurrentAdversery[AdverseryChoice].Drop[RandomeObj].Name)
        }
    }
}
func AtackWeaponSystème() {
    if Joueur.Vieactuelle > 0 && CurrentAdversery[AdverseryChoice].Vieactuelle > 0 {
        var AttackChoice int
        var Attaques = []Attaques{}
        Attaques = append(Attaques, Joueur.EquipedWeapon.Attaques...)
        var RandomeRate int
        RandomeRate = rand.Intn(101)
        fmt.Scan(&AttackChoice)
        if AttackChoice-1 > len(Attaques)+2 {
            fmt.Print("\033[31m❌ Ce n'est pas une proposition espèce de demeuré\033[0m\n")
            AtackBasiqueSystème()
            return
        }
        if AttackChoice -1 ==  len(Attaques)+2 && !PourUnSeulMec.Cap[0].Used && Joueur.EquipedWeapon.Name == PourUnSeulMec.Name[1] {
            fmt.Print("⏳ Le temps se déforme... 🌀 Vous allez rejouer ! 🔁\n")
            PourUnSeulMec.Cap[0].Used = true
            DisplayAtackArmes()
            AtackWeaponSystème()
        }

        if AttackChoice == len(Attaques)+1 {
            CombatMode()
            return
        }
        if AttackChoice-1 < len(Attaques) {
            if RandomeRate <= int(Attaques[AttackChoice-1].LandingRate) {
                if Joueur.EquipedWeapon.Mastery >= Attaques[AttackChoice-1].NeededMastery {
                    CurrentAdversery[AdverseryChoice].Vieactuelle -= Attaques[AttackChoice-1].Damage * Arrondir(float64(1+Joueur.Intelligence/100))
                    if CurrentAdversery[AdverseryChoice].Vieactuelle < 0 {
                        CurrentAdversery[AdverseryChoice].Vieactuelle = 0
                    }
                    fmt.Print("\033[32m✅ Votre attaque a réussi !\033[0m\n")
                    fmt.Printf("\033[33mPV adversaire : \033[31m%d\033[0m/\033[32m%d\033[0m\n",
                        CurrentAdversery[AdverseryChoice].Vieactuelle,
                        CurrentAdversery[AdverseryChoice].Viemax)
                } else {
                    fmt.Print("\033[31m❌ Vous n'avez pas le niveau de maîtrise nécessaire\033[0m\n")
                    DisplayAtackArmes()
                    AtackWeaponSystème()
                }
            } else {
                fmt.Print("\033[31m❌ Votre attaque a raté !\033[0m\n")
            }
        }
    }
}
