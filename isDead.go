package personnage

import (
	"fmt"
)

func isDead() {
	if Joueur.Vieactuelle <= 0 {
		fmt.Print("\033[31m💀 Vous êtes mort, sale noob 😂\033[0m\n")
	}
}