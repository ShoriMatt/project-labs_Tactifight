package main

import (
	"bufio"
	"fmt"
	"strings"
)

func characterTurn(player *Character, monster *Monster, reader *bufio.Reader, turn int) {
	centerText("\n=== MENU DE COMBAT ===")
	centerText("1. Attaquer")
	centerText("2. Inventaire")
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
		centerText(fmt.Sprintf("%s utilise Attaque basique !", player.Name))
		centerText(fmt.Sprintf("Dégâts infligés : %d", damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", monster.Name, monster.HP, monster.MaxHP))

	case "2": // Inventaire
		accessInventory(player, reader)

	default:
		centerText("Choix invalide. Vous perdez votre tour.")
	}
}

func mainMenu(c *Character, reader *bufio.Reader) {
	for {
		centerText(`
 _, _ __, _, _ _,_   __, __, _ _, _  _, _ __,  _, _, 
 |\/| |_  |\ | | |   |_) |_) | |\ | / ` + "`" + ` | |_) /_\ |  
 |  | |   | \| | |   |   | \ | | \| \ , | |   | | | ,
 ~  ~ ~~~ ~  ~ ` + "`~'" + `   ~   ~ ~ ~ ~  ~  ~  ~ ~   ~ ~ ~~~
													 `)
		centerText("1 - Afficher les informations du personnage")
		centerText("2 - Accéder au contenu de l'inventaire")
		centerText("3 - Voir le Forgeron")
		centerText("4 - Combattre le Gobelin d'entraînement")
		centerText("5 - Qui sont-ils ?")
		centerText("6 - Quitter")
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
			centerText("\n--- Début du combat contre le Gobelin d’entraînement ---")

			turn := 1
			for goblin.HP > 0 && c.HP > 0 {
				characterTurn(c, &goblin, reader, turn)
				if goblin.HP > 0 {
					goblinPattern(&goblin, c, turn) // riposte du gobelin
				}
				turn++
			}

			if goblin.HP <= 0 {
				centerText("🎉 Vous avez vaincu le gobelin !")
				c.gainXP(goblin.XPReward)
			} else if c.HP <= 0 {
				centerText("💀 Vous avez été vaincu...")
			}

		case "5":
			centerText("Les artistes cachés sont ABBA et Steven Spielberg !")
		case "6":
			centerText("Au revoir !")
			return
		default:
			centerText("Choix invalide.")
		}
	}
}
