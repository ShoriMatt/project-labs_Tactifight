package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Structure d'un objet
type Item struct {
	Name   string
	Type   string // ex: "heal", "poison"
	Effect int    // quantité de PV rendus ou dégâts
}

// Base de données des objets disponibles
var ItemsDB = map[string]Item{
	"Potion de soin":   {Name: "Potion de soin", Type: "heal", Effect: 50},
	"Potion de poison": {Name: "Potion de poison", Type: "poison", Effect: 30},
}

func patern(p int, c *Character, turn int, Mob Monster) {
	// Exemple de fonction pour un pattern d'attaque
	switch p {
	case 1:
		ChameauToxiqueLunairePattern(&Mob, c, turn)
	case 2:
		CloneBancalDeMacronPattern(&Mob, c, turn)
	case 3:
		GardePrésidentielSpatialPattern(&Mob, c, turn)
	case 0:
		EmmanuelMacroniusIVPattern(&Mob, c, turn)
	default:
		fmt.Println("Pattern inconnu")
	}
}
func combat(c *Character, m *Character) {
	Mob := initChameauToxiqueLunaire()
	p := 1
	fmt.Println("Un Chameau toxique lunaire apparaît !")
	if c.etage == 20 {
		Mob = initEmmanuelMacroniusIV()
		p = 0
		fmt.Println("Vous affrontez le boss final : Emmanuel Macronius IV !")
	} else {
		mobRand := rand.Intn(100)
		if mobRand <= 45 {
			Mob = initChameauToxiqueLunaire()
			p = 1
			fmt.Println("Un Chameau toxique lunaire apparaît !")
		} else if mobRand <= 80 && mobRand > 45 {
			Mob = initCloneBancalDeMacron()
			p = 2
			fmt.Println("Un Clone bancal de Macron apparaît !")
		} else {
			Mob = initGardePrésidentielSpatial()
			p = 3
			fmt.Println("Un Garde Présidentiel Spatial apparaît !")
		}
	}
	turn := 1
	fuite := false
	for Mob.HP > 0 && c.HP > 0 {
		fuite = characterTurn(c, &Mob, reader, turn)
		if fuite {
			break // sortir du combat si fuite
		}
		if Mob.HP > 0 {
			patern(p, c, turn, Mob) // riposte du gobelin
		}
		turn++
	}
	if fuite {
		centerText("Vous avez quitté le combat.")
	} else if Mob.HP <= 0 {
		centerText("🎉 Vous avez vaincu le gobelin !")
		c.gainXP(Mob.XPReward)
		c.Gold += Mob.GoldReward
		centerText(fmt.Sprintf("💰 Vous obtenez %d or !", Mob.GoldReward))
	} else if c.HP <= 0 {
		centerText("💀 Vous avez été vaincu...")
	}

}

// Applique les dégâts du poison au début du tour si empoisonné
func ApplyPoisonEffect(c *Character) {
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

// Tour du joueur (à appeler depuis main ou autre)
func PlayerTurn(player *Character, enemy *Character) {
	reader := bufio.NewReader(os.Stdin)

	ApplyPoisonEffect(player)
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
				MonsterTurn(enemy, player)
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
			for i, itemName := range player.Inventory {
				fmt.Printf("%d. %s\n", i+1, itemName)
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

			chosenName := player.Inventory[index-1]
			chosenItem, exists := ItemsDB[chosenName]
			if !exists {
				fmt.Println("Objet inconnu.")
				continue
			}

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

			if enemy.HP > 0 {
				MonsterTurn(enemy, player)
			}
			return

		default:
			fmt.Println("Choix invalide. Réessayez.")
		}
	}
}

// Tour du monstre (à appeler depuis main ou autre)
func MonsterTurn(monster *Character, player *Character) {
	ApplyPoisonEffect(monster)
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
