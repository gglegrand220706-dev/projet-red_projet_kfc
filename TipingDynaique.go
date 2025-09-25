package personnage

import (
	"fmt"
	"time"
)

func DynamiqueType(Phrases string, Wait int) {
	for _, CurrentCharacter := range Phrases {
		fmt.Printf("%c", CurrentCharacter)
		time.Sleep(time.Duration(Wait) * time.Millisecond)
	}
}
