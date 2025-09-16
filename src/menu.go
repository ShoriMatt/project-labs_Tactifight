package main

import (
	"bufio"
	"fmt"
	"strings"
)

func characterTurn(player *Character, monster *Monster, reader *bufio.Reader, turn int) {
	fmt.Println("\n=== MENU DE COMBAT ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Print("Choix > ")

	action, _ := reader.ReadString('\n')
	action = strings.TrimSpace(action)

	switch action {
	case "1": // Attaque basique
		damage := 5
		monster.HP -= damage
		if monster.HP < 0 {
			monster.HP = 0
		}
		fmt.Printf("%s utilise Attaque basique !\n", player.Name)
		fmt.Printf("Dégâts infligés : %d\n", damage)
		fmt.Printf("%s : %d/%d PV\n", monster.Name, monster.HP, monster.MaxHP)

	case "2": // Inventaire
		if len(player.Inventory) == 0 {
			fmt.Println("Votre inventaire est vide.")
		} else {
			fmt.Println("Inventaire :")
			for i, it := range player.Inventory {
				fmt.Printf("  %d. %s\n", i+1, it)
			}
			fmt.Print("Choisissez un objet > ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)

			index := -1
			for i, it := range player.Inventory {
				if strings.EqualFold(it, choice) {
					index = i
					break
				}
			}
			if index != -1 {
				item := player.Inventory[index]
				fmt.Printf("Vous utilisez %s\n", item)
				// Exemple effet : potion de vie
				if strings.Contains(strings.ToLower(item), "potion") {
					player.HP += 20
					if player.HP > player.MaxHP {
						player.HP = player.MaxHP
					}
					fmt.Printf("%s récupère de la vie (%d/%d PV)\n", player.Name, player.HP, player.MaxHP)
				}
				// Supprimer l’objet utilisé
				player.Inventory = append(player.Inventory[:index], player.Inventory[index+1:]...)
			} else {
				fmt.Println("Objet introuvable.")
			}
		}

	default:
		fmt.Println("Choix invalide. Vous perdez votre tour.")
	}
}

func mainMenu(c *Character, reader *bufio.Reader) {
	for {
		fmt.Println(`
 _, _ __, _, _ _,_   __, __, _ _, _  _, _ __,  _, _, 
 |\/| |_  |\ | | |   |_) |_) | |\ | / ` + "`" + ` | |_) /_\ |  
 |  | |   | \| | |   |   | \ | | \| \ , | |   | | | ,
 ~  ~ ~~~ ~  ~ ` + "`~'" + `   ~   ~ ~ ~ ~  ~  ~  ~ ~   ~ ~ ~~~
													 `)
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

			turn := 1
			for goblin.HP > 0 && c.HP > 0 {
				characterTurn(c, &goblin, reader, turn)
				if goblin.HP > 0 {
					goblinPattern(&goblin, c, turn) // riposte du gobelin
				}
				turn++
			}

			if goblin.HP <= 0 {
				fmt.Println("🎉 Vous avez vaincu le gobelin !")
				c.gainXP(goblin.XPReward)
			} else if c.HP <= 0 {
				fmt.Println("💀 Vous avez été vaincu...")
			}

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
