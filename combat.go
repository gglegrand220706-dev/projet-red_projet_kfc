package personnage

import (
	"fmt"
)

func AtackSysteme() {
    fmt.Println("Que voulez-vous faire :")
    for i, cap := range Joueur.Capacity {
        fmt.Printf("%d. %s\n", i+1, cap)
    }
}

