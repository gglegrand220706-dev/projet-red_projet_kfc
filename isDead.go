package personnage

import (
	"fmt"
)

func isDead() {
	if Joueur.Vieactuelle <= 0 {
		fmt.Print("Vous êtes mort sale noob")
	}
}