package main

import (
	"bufio"
	"fmt"
	"strings"
)

func mainMenu(c *Character, reader *bufio.Reader) {
	for {
		fmt.Println("\nMenu Principal")
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder au contenu de l'inventaire")
		fmt.Println("3 - Voir le Forgeron")
		fmt.Println("4 - Combattre le Gobelin d'entraînement")
		fmt.Println("5 - Qui sont-ils ?")
		fmt.Println("6 - Quitter")
		fmt.Print("Choix > ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			displayInfo(c)
		case "2":
			accessInventory(c, reader)
		case "3":
			forgeron(c, reader)
		case "4":
			goblin := initGoblin()
			fmt.Println("\n--- Début du combat contre le Gobelin d’entraînement ---")
			goblinPattern(&goblin, c, 6)
		case "5":
			fmt.Println("Les artistes cachés sont ABBA et Steven Spielberg !")
		case "6":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
