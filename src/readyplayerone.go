package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Structure d'un objet
type Item struct {
	Name   string
	Type   string // ex: "heal", "poison"
	Effect int    // quantité de PV rendus ou dégâts
}

// Structure d'un personnage
type Character struct {
	Name       string
	HP         int
	MaxHP      int
	Damage     int
	Inventory  []Item
	PoisonTurns int // Nombre de tours restants sous poison
}

// Applique les dégâts du poison au début du tour si empoisonné
func applyPoisonEffect(c *Character) {
	if c.PoisonTurns > 0 {
		fmt.Printf("%s souffre du poison !\n", c.Name)
		c.HP -= 10
		if c.HP < 0 {
			c.HP = 0
		}
		fmt.Printf("%s - PV : %d / %d\n", c.Name, c.HP, c.MaxHP)
		c.PoisonTurns--

		if c.HP == 0 {
			fmt.Printf("%s est vaincu par le poison !\n", c.Name)
		}
	}
}

// Fonction du tour du joueur
func characterTurn(player *Character, enemy *Character) {
	reader := bufio.NewReader(os.Stdin)

	// Appliquer poison au début du tour du joueur
	applyPoisonEffect(player)
	if player.HP == 0 {
		return
	}

	for {
		fmt.Println("\n=== Menu ===")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Print("Choisissez une option : ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			attackName := "Attaque basique"
			damage := 5
			enemy.HP -= damage
			if enemy.HP < 0 {
				enemy.HP = 0
			}

			fmt.Printf("\n%s inflige %d dégâts à %s avec %s\n", player.Name, damage, enemy.Name, attackName)
			fmt.Printf("%s - PV : %d / %d\n", enemy.Name, enemy.HP, enemy.MaxHP)

			if enemy.HP > 0 {
				monsterTurn(enemy, player)
			} else {
				fmt.Printf("%s est vaincu !\n", enemy.Name)
			}
			return

		case "2":
			if len(player.Inventory) == 0 {
				fmt.Println("Votre inventaire est vide.")
				continue
			}

			fmt.Println("\n=== Inventaire ===")
			for i, item := range player.Inventory {
				fmt.Printf("%d. %s\n", i+1, item.Name)
			}
			fmt.Print("Choisissez un objet à utiliser : ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			var index int
			fmt.Sscanf(input, "%d", &index)
			if index < 1 || index > len(player.Inventory) {
				fmt.Println("Choix invalide.")
				continue
			}

			chosenItem := player.Inventory[index-1]
			fmt.Printf("\nVous utilisez %s\n", chosenItem.Name)

			switch chosenItem.Type {
			case "heal":
				player.HP += chosenItem.Effect
				if player.HP > player.MaxHP {
					player.HP = player.MaxHP
				}
				fmt.Printf("%s récupère %d PV.\n", player.Name, chosenItem.Effect)
				fmt.Printf("%s - PV : %d / %d\n", player.Name, player.HP, player.MaxHP)

			case "poison":
				// Appliquer l'état poison à l'ennemi
				if enemy.PoisonTurns > 0 {
					fmt.Printf("%s est déjà empoisonné.\n", enemy.Name)
				} else {
					enemy.PoisonTurns = 3
					fmt.Printf("%s est empoisonné pour 3 tours !\n", enemy.Name)
				}

			default:
				fmt.Println("Type d'objet inconnu.")
			}

			// Supprimer l'objet utilisé
			player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)

			// Tour du monstre
			if enemy.HP > 0 {
				monsterTurn(enemy, player)
			}
			return

		default:
			fmt.Println("Choix invalide. Réessayez.")
		}
	}
}

// Tour du monstre
func monsterTurn(monster *Character, player *Character) {
	// Appliquer poison au début du tour du monstre
	applyPoisonEffect(monster)
	if monster.HP == 0 {
		return
	}

	attackName := "Coup de griffe"
	damage := 3
	player.HP -= damage
	if player.HP < 0 {
		player.HP = 0
	}

	fmt.Printf("\n%s utilise %s !\n", monster.Name, attackName)
	fmt.Printf("%s inflige %d dégâts à %s.\n", monster.Name, damage, player.Name)
	fmt.Printf("%s - PV : %d / %d\n", player.Name, player.HP, player.MaxHP)

	if player.HP == 0 {
		fmt.Printf("%s est vaincu...\n", player.Name)
	}
}

// Fonction principale
func main() {
	player := &Character{
		Name:   "Héros",
		HP:     20,
		MaxHP:  20,
		Damage: 5,
		Inventory: []Item{
			{Name: "Potion de soin", Type: "heal", Effect: 50},
			{Name: "Potion de poison", Type: "poison", Effect: 0},
		},
	}

	monster := &Character{
		Name:   "Gobelin d'entraînement",
		HP:     40,
		MaxHP:  40,
		Damage: 3,
	}

	for player.HP > 0 && monster.HP > 0 {
		characterTurn(player, monster)
	}

	fmt.Println("\nCombat terminé.")
}
