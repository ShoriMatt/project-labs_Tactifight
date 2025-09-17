package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Structure d'un objet
type Item struct {
	Name   string
	Type   string // ex: "heal", "poison"
	Effect int    // quantité de PV rendus ou dégâts
}

// Base de données des objets disponibles
var ItemsDB = map[string]Item{
	"potion de vie":                               {"potion de vie", "heal", 50},
	"potion de mana":                              {"potion de mana", "mana", 30},
	"potion de poison":                            {"potion de poison", "poison", 10},
	"potion de poison cosmique":                   {"potion de poison cosmique", "poison", 15},
	"livre de sort : explosion de sable cosmique": {"livre de sort : explosion de sable cosmique", "spell", 0},
	"sceptre-laser doré":                          {"sceptre-laser doré", "equip", 5},
	"trône gravitationnel":                        {"trône gravitationnel", "artefact", 20},
}

var SpellsDB = map[string]Item{
	"balle perçante":              {"balle perçante", "spell", 6},
	"explosion de sable cosmique": {"explosion de sable cosmique", "spell", 20},
}

// MonsterPoisonTurns tracks remaining poison turns per Monster pointer.
// This avoids accessing a non-existent PoisonTurns field on the Monster struct.
var MonsterPoisonTurns = make(map[*Monster]int)

// AttackBonuses tracks permanent attack bonuses per character pointer.
// Some Character implementations may not have an AttackBonus field, so keep bonuses here.
var AttackBonuses = make(map[*Character]int)

func useItem(c *Character, mob *Monster, reader *bufio.Reader) {
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	fmt.Println("Inventaire :")
	for i, item := range c.Inventory {
		fmt.Printf("%d - %s\n", i+1, item)
	}
	fmt.Print("Choix > ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Action annulée.")
		return
	}

	idx := int(input[0]-'0') - 1
	if idx < 0 || idx >= len(c.Inventory) {
		fmt.Println("Choix invalide.")
		return
	}

	itemName := strings.ToLower(c.Inventory[idx])
	item, exists := ItemsDB[itemName]
	if !exists {
		fmt.Println("Objet inconnu.")
		return
	}

	switch item.Type {
	case "heal":
		old := c.HP
		c.HP += item.Effect
		if c.HP > c.MaxHP {
			c.HP = c.MaxHP
		}
		fmt.Printf("Vous utilisez %s et récupérez %d PV (%d/%d).\n", item.Name, c.HP-old, c.HP, c.MaxHP)

	case "mana":
		old := c.Mana
		c.Mana += item.Effect
		if c.Mana > c.MaxMana {
			c.Mana = c.MaxMana
		}
		fmt.Printf("Vous utilisez %s et récupérez %d Mana (%d/%d).\n", item.Name, c.Mana-old, c.Mana, c.MaxMana)

	case "poison":
		if mob != nil {
			MonsterPoisonTurns[mob] = item.Effect
			fmt.Printf("☠️ Vous utilisez %s ! %s est empoisonné (%d tours).\n", item.Name, mob.Name, item.Effect)
		}

	case "spell":
		if item.Name == "livre de sort : explosion de sable cosmique" {
			if !contains(c.Skills, "Explosion de sable cosmique") {
				c.Skills = append(c.Skills, "Explosion de sable cosmique")
				centerText("Nouvelle compétence apprise : Explosion de sable cosmique !")
			}
		}

	case "equip":
		if item.Name == "sceptre-laser doré" {
			AttackBonuses[c] += item.Effect
			fmt.Printf("+%d dégâts permanents grâce au Sceptre-laser doré (bonus total: +%d).\n", item.Effect, AttackBonuses[c])
		}

	case "artefact":
		if item.Name == "trône gravitationnel" {
			c.MaxHP += item.Effect
			fmt.Printf("Votre PV max augmente de %d grâce au Trône gravitationnel !\n", item.Effect)
		}
	}

	removeInventory(c, c.Inventory[idx])
}

func addInventory(c *Character, item string) {
	if len(c.Inventory) >= c.InventoryCapacity {
		centerText(fmt.Sprintf("Inventaire plein ! (max %d)", c.InventoryCapacity))
		return
	}
	item = strings.ToLower(item)
	c.Inventory = append(c.Inventory, item)
	centerText(fmt.Sprintf("%s a été ajouté à l'inventaire.", formatNom(item)))
}

func removeInventory(c *Character, item string) bool {
	for i, v := range c.Inventory {
		if strings.EqualFold(v, item) {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func takePotion(c *Character) {
	found := -1
	for i, v := range c.Inventory {
		if strings.Contains(strings.ToLower(v), "potion de vie") {
			found = i
			break
		}
	}
	if found == -1 {
		centerText("Aucune potion de vie disponible.")
		return
	}

	potion := c.Inventory[found]

	oldHP := c.HP
	c.HP += 50
	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
	centerText(fmt.Sprintf("Tu as utilisé %s : PV %d -> %d / %d", potion, oldHP, c.HP, c.MaxHP))
}

// Potion de poison : utilisable seulement en combat, inflige des dégâts au monstre
func poisonPot(c *Character, m *Monster, enCombat bool) {
	if !enCombat {
		centerText("Tu ne peux utiliser la potion de poison qu'en combat !")
		return
	}
	if m == nil {
		centerText("Aucun monstre à empoisonner.")
		return
	}

	centerText("Tu as utilisé une potion de poison sur le monstre !")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		m.HP -= 10
		if m.HP < 0 {
			m.HP = 0
		}
		centerText(fmt.Sprintf("Le monstre subit des dégâts de poison (%ds) : %d/%d PV", i, m.HP, m.MaxHP))
		if m.IsDead() {
			centerText("Le monstre est mort à cause du poison !")
			return
		}
	}
}

func accessInventory(c *Character, m *Monster, enCombat bool, reader *bufio.Reader) {
	for {
		centerText("\n===========================================")
		centerText("Inventaire")
		if len(c.Inventory) == 0 {
			centerText("(vide)")
		} else {
			for i, item := range c.Inventory {
				centerText(fmt.Sprintf("%d. %s", i+1, formatNom(item)))
			}
		}
		centerText("===========================================\n")
		centerText("Options :")
		centerText("u - Utiliser un objet")
		centerText("e - Équiper un objet")
		centerText("m - Aller voir le marchand")
		centerText("b - Retour")
		fmt.Print("Choix > ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "b" {
			return
		}
		if choice == "u" {
			if len(c.Inventory) == 0 {
				centerText("Aucun objet à utiliser.")
				continue
			}
			fmt.Print("Numéro de l'objet à utiliser : ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			idx, err := strconv.Atoi(numStr)
			if err != nil || idx < 1 || idx > len(c.Inventory) {
				centerText("Numéro invalide.")
				continue
			}
			item := c.Inventory[idx-1]

			if strings.Contains(strings.ToLower(item), "potion de vie") {
				takePotion(c)
				removeInventory(c, item)
			} else if strings.Contains(strings.ToLower(item), "potion de poison") {
				removeInventory(c, item)
				poisonPot(c, m, enCombat)
			} else if strings.Contains(strings.ToLower(item), "balle perçante") {
				removeInventory(c, item)
				spellBook(c)
			} else if strings.Contains(strings.ToLower(item), "amélioration d'inventaire") {
				removeInventory(c, item)
				c.upgradeInventorySlot()
			} else {
				centerText(fmt.Sprintf("L'utilisation de %s n'est pas encore implémentée.", item))
			}
		}
		if choice == "e" {
			if len(c.Inventory) == 0 {
				centerText("Aucun objet à équiper.")
				continue
			}
			fmt.Print("Numéro de l'objet à équiper : ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			idx, err := strconv.Atoi(numStr)
			if err != nil || idx < 1 || idx > len(c.Inventory) {
				centerText("Numéro invalide.")
				continue
			}
			item := c.Inventory[idx-1]
			c.equip(item)
		}
		if choice == "m" {
			if enCombat {
				centerText("Impossible d'aller voir le marchand en plein combat !")
				continue
			}
			marchand(c, reader)
		}
	}
}
