package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Spell struct {
	Name   string
	Cost   int
	Damage int
	Effect string // pour gérer des effets spéciaux plus tard (poison, heal, etc.)
}

var SpellDB = map[string]Spell{
	"Explosion de sable cosmique": {
		Name:   "Explosion de sable cosmique",
		Cost:   10,
		Damage: 20,
		Effect: "damage",
	},
	"Éclair gravitationnel": {
		Name:   "Éclair gravitationnel",
		Cost:   15,
		Damage: 30,
		Effect: "damage",
	},
	"Soin stellaire": {
		Name:   "Soin stellaire",
		Cost:   12,
		Damage: -20, // négatif pour soin
		Effect: "heal",
	},
}

// Gère quel pattern d'attaque utiliser selon le monstre
func patern(p int, c *Character, turn int, mob *Monster) {
	switch p {
	case 1:
		ChameauToxiqueLunairePattern(mob, c, turn)
	case 2:
		CloneBancalDeMacronPattern(mob, c, turn)
	case 3:
		GardePrésidentielSpatialPattern(mob, c, turn)
	case 0:
		EmmanuelMacroniusIVPattern(mob, c, turn)
	default:
		fmt.Println("Pattern inconnu")
	}
}

// Combat principal
func combat(c *Character) {
	var Mob Monster
	p := 1

	if c.etage == 20 {
		Mob = initEmmanuelMacroniusIV()
		p = 0
		centerText("Vous affrontez le boss final : Emmanuel Macronius IV !")
	} else {
		mobRand := rand.Intn(100)
		if mobRand <= 45 {
			Mob = initChameauToxiqueLunaire()
			p = 1
			centerText("Un Chameau toxique lunaire apparaît !")
		} else if mobRand <= 80 {
			Mob = initCloneBancalDeMacron()
			p = 2
			centerText("Un Clone bancal de Macron apparaît !")
		} else {
			Mob = initGardePrésidentielSpatial()
			p = 3
			centerText("Un Garde Présidentiel Spatial apparaît !")
		}
	}

	turn := 1
	fuite := false
	for !Mob.IsDead() && c.HP > 0 {
		fuite = PlayerTurn(c, &Mob)
		if fuite {
			break
		}
		if !Mob.IsDead() {
			patern(p, c, turn, &Mob)
		}
		turn++
	}

	if fuite {
		centerText("Vous avez quitté le combat.")
	} else if Mob.IsDead() {
		centerText("🎉 Vous avez vaincu le monstre !")
		c.etage++
		centerText(fmt.Sprintf("Vous montez à l'étage %d", c.etage))
		c.gainXP(Mob.XPReward)
		c.Gold += Mob.GoldReward
		centerText(fmt.Sprintf("💰 Vous obtenez %d or !", Mob.GoldReward))
		switch Mob.Name {
		case "Chameau toxique lunaire":
			addInventory(c, "potion de poison cosmique")
			centerText("🎁 Vous obtenez une Potion de poison cosmique !")

		case "Clone Bancal de Macron":
			spell := "Explosion de sable cosmique"
			if !contains(c.Skills, spell) {
				c.Skills = append(c.Skills, spell)
				centerText("📖 Nouveau sort appris : Explosion de sable cosmique !")
			}

		case "Garde Présidentiel Spatial":
			addInventory(c, "sceptre-laser doré")
			centerText("⚔️ Vous obtenez le Sceptre-laser doré !")

		case "Emmanuel Macronius IV":
			centerText("🎆 Vous avez récupéré votre liberté !")
			addInventory(c, "trône gravitationnel")
			centerText("👑 Artefact obtenu : Trône gravitationnel (PV max augmenté)")
			c.MaxHP += 30
			c.HP = c.MaxHP
		}

	} else if c.HP <= 0 {
		centerText(fmt.Sprintf("\n💀 %s est tombé au combat !\n", c.Name))

		// Résurrection automatique
		c.HP = c.MaxHP / 2
		c.Mana = c.MaxMana / 2

		centerText(fmt.Sprintf("✨ %s est ressuscité par une énergie cosmique !\n", c.Name))
		centerText(fmt.Sprintf("Vous revenez avec %d/%d PV et %d/%d Mana.\n", c.HP, c.MaxHP, c.Mana, c.MaxMana))

		return
	}
}

// Applique les dégâts du poison au début du tour si empoisonné
func ApplyPoisonEffect(c *Character) {
	if c.PoisonTurns > 0 {
		centerText(fmt.Sprintf("%s souffre du poison !\n", c.Name))
		c.HP -= 10
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s - PV : %d / %d\n", c.Name, c.HP, c.MaxHP))
		c.PoisonTurns--

		if c.HP == 0 {
			centerText(fmt.Sprintf("%s est vaincu par le poison !\n", c.Name))
		}
	}
}

// Tour du joueur
func PlayerTurn(player *Character, enemy *Monster) bool {
	reader := bufio.NewReader(os.Stdin)

	ApplyPoisonEffect(player)
	if player.HP == 0 {
		return false
	}

	for {
		centerText("\n=== MENU DE COMBAT ===")
		centerText("1. Attaquer")
		centerText("2. Inventaire")
		centerText("3. Fuir")
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

			centerText(fmt.Sprintf("%s inflige %d dégâts à %s avec %s", player.Name, damage, enemy.Name, attackName))
			centerText(fmt.Sprintf("%s - PV : %d / %d", enemy.Name, enemy.HP, enemy.MaxHP))

			return false

		case "2":
			if len(player.Inventory) == 0 {
				centerText("Votre inventaire est vide.")
				continue
			}

			centerText("\n=== Inventaire ===")
			for i, itemName := range player.Inventory {
				centerText(fmt.Sprintf("%d. %s", i+1, itemName))
			}
			fmt.Print("Choisissez un objet à utiliser : ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			var index int
			fmt.Sscanf(input, "%d", &index)
			if index < 1 || index > len(player.Inventory) {
				centerText("Choix invalide.")
				continue
			}

			chosenName := player.Inventory[index-1]
			chosenItem, exists := ItemsDB[chosenName]
			if !exists {
				centerText("Objet inconnu.")
				continue
			}

			centerText(fmt.Sprintf("Vous utilisez %s", chosenItem.Name))

			switch chosenItem.Type {
			case "heal":
				player.HP += chosenItem.Effect
				if player.HP > player.MaxHP {
					player.HP = player.MaxHP
				}
				centerText(fmt.Sprintf("%s récupère %d PV.", player.Name, chosenItem.Effect))
				centerText(fmt.Sprintf("%s - PV : %d / %d", player.Name, player.HP, player.MaxHP))

			case "poison":
				if enemy.HP > 0 {
					centerText(fmt.Sprintf("%s est empoisonné pour 3 tours !", enemy.Name))
					enemy.HP -= chosenItem.Effect
					if enemy.HP < 0 {
						enemy.HP = 0
					}
				}

			default:
				centerText("Type d'objet inconnu.")
			}

			// Supprimer l'objet utilisé
			player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)

			return false

		case "3":
			// Lancer un sort
			if len(player.Skills) == 0 {
				centerText("Vous ne connaissez aucun sort.")
				continue
			}

			centerText("\n=== Sorts disponibles ===")
			for i, spell := range player.Skills {
				if s, ok := SpellDB[spell]; ok {
					centerText(fmt.Sprintf("%d. %s (Coût: %d mana, Effet: %s)", i+1, s.Name, s.Cost, s.Effect))
				} else {
					centerText(fmt.Sprintf("%d. %s (inconnu)", i+1, spell))
				}
			}

			fmt.Print("Choisissez un sort : ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			var index int
			fmt.Sscanf(input, "%d", &index)
			if index < 1 || index > len(player.Skills) {
				centerText("Choix invalide.")
				continue
			}

			chosenSpell := player.Skills[index-1]
			spell, exists := SpellDB[chosenSpell]
			if !exists {
				centerText("Ce sort est inconnu.")
				continue
			}

			if player.Mana < spell.Cost {
				centerText("Pas assez de mana !")
				continue
			}

			player.Mana -= spell.Cost

			switch spell.Effect {
			case "damage":
				enemy.HP -= spell.Damage
				if enemy.HP < 0 {
					enemy.HP = 0
				}
				centerText(fmt.Sprintf("%s lance %s et inflige %d dégâts à %s !", player.Name, spell.Name, spell.Damage, enemy.Name))
				centerText(fmt.Sprintf("%s - PV : %d / %d", enemy.Name, enemy.HP, enemy.MaxHP))

			case "heal":
				player.HP -= spell.Damage // rappel : négatif → soin
				if player.HP > player.MaxHP {
					player.HP = player.MaxHP
				}
				centerText(fmt.Sprintf("%s lance %s et récupère %d PV !", player.Name, spell.Name, -spell.Damage))
				centerText(fmt.Sprintf("%s - PV : %d / %d", player.Name, player.HP, player.MaxHP))

			default:
				centerText("L'effet du sort n'est pas encore implémenté.")
			}

			centerText(fmt.Sprintf("Mana restant : %d / %d", player.Mana, player.MaxMana))
			return false

		case "4":
			centerText("Vous prenez la fuite !")
			return true

		default:
			centerText("Choix invalide. Réessayez.")
		}
	}
}
