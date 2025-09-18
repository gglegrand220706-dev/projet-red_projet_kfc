package personnage

import (
    "fmt"
    "time"
    "bufio"
    "os"
)

func WelcomeRecords() {
    fmt.Print("\033[H\033[2J")
    fmt.Println("\033[31m")
    fmt.Println("$$\\      $$\\           $$\\                                                   $$$$$$$\\                                                $$\\ ")
    fmt.Println("$$ | $\\  $$ |          $$ |                                                  $$  __$$\\                                               $$ |")
    fmt.Println("$$ |$$$\\ $$ | $$$$$$\\  $$ | $$$$$$$\\  $$$$$$\\  $$$$$$\\$$$$\\   $$$$$$\\        $$ |  $$ | $$$$$$\\   $$$$$$$\\  $$$$$$\\   $$$$$$\\   $$$$$$$ |")
    fmt.Println("$$ $$ $$\\$$ |$$  __$$\\ $$ |$$  _____|$$  __$$\\ $$  _$$  _$$\\ $$  __$$\\       $$$$$$$  |$$  __$$\\ $$  _____|$$  __$$\\ $$  __$$\\ $$  __$$ |")
    fmt.Println("$$$$  _$$$$ |$$$$$$$$ |$$ |$$ /      $$ /  $$ |$$ / $$ / $$ |$$$$$$$$ |      $$  __$$< $$$$$$$$ |$$ /      $$ /  $$ |$$ |  \\__|$$ /  $$ |")
    fmt.Println("$$$  / \\$$$ |$$   ____|$$ |$$ |      $$ |  $$ |$$ | $$ | $$ |$$   ____|      $$ |  $$ |$$   ____|$$ |      $$ |  $$ |$$ |      $$ |  $$ |")
    fmt.Println("$$  /   \\$$ |\\$$$$$$$\\ $$ |\\$$$$$$$\\ \\$$$$$$  |$$ | $$ | $$ |\\$$$$$$$\\       $$ |  $$ |\\$$$$$$$\\ \\$$$$$$$\\ \\$$$$$$  |$$ |      \\$$$$$$$ |")
    fmt.Println("\\__/     \\__| \\_______|\\__| \\_______| \\______/ \\__| \\__| \\__| \\_______|      \\__|  \\__| \\_______| \\_______| \\______/ \\__|       \\_______|")
    fmt.Println("\033[0m")

    fmt.Println("\033[33m=== LORE ===\033[0m\n")

    lignes := []string{
        "Il y a bien longtemps, aux confins du multivers, une entité mécanique nommée TeKnologia a brisé l’équilibre cosmique.",
        "Deux réalités légendaires se sont effondrées, et de leurs fragments est né un monde unique, instable et imprévisible : le Battle World.",
        "Dans ce royaume façonné par la fusion de mondes entiers, les lois de la nature et du temps ne sont plus les mêmes.",
        "Les héros et les légendes d’antan ne sont plus que des échos... mais leurs pouvoirs, eux, subsistent.",
        "Vous êtes l’un des rares capables d’en hériter.",
        "Votre destinée : explorer ce monde, percer ses secrets, et rassembler les artefacts nécessaires pour affronter TeKnologia elle-même.",
        "La victoire pourrait restaurer l’ordre... ou plonger à jamais le multivers dans le chaos.",
        "Votre légende commence maintenant.",
    }

    for _, ligne := range lignes {
        for _, char := range ligne {
            fmt.Printf("\033[33m%c\033[0m", char)
            time.Sleep(50 * time.Millisecond) 
        }
        fmt.Print("\n\n") 
        time.Sleep(1 * time.Second) 
    }

    fmt.Println("\n\033[33mÊTES-VOUS PRÊT ? Appuyez sur ENTRER pour commencer votre aventure...\033[0m")
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}