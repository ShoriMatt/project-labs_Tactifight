package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func addInventory(c *Character, item string) {
	if len(c.Inventory) >= c.InventoryCapacity {
		fmt.Printf("Inventaire plein ! (max %d)\n", c.InventoryCapacity)
		return
	}
	item = strings.ToLower(item)
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("%s a été ajouté à l'inventaire.\n", formatNom(item))
}

func removeInventory(c *Character, item string) bool {
	for i, v := range c.Inventory {
		if v == item {
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
		fmt.Println("Aucune potion de vie disponible.")
		return
	}

	potion := c.Inventory[found]
	c.Inventory = append(c.Inventory[:found], c.Inventory[found+1:]...)

	oldHP := c.HP
	c.HP += 50
	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
	fmt.Printf("Tu as utilisé %s : PV %d -> %d / %d\n", potion, oldHP, c.HP, c.MaxHP)
}

func poisonPot(c *Character) {
	fmt.Println("Tu as utilisé une potion de poison !")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		c.HP -= 10
		if c.HP < 0 {
			c.HP = 0
		}
		fmt.Printf("Dégâts de poison (%ds) : %d/%d PV\n", i, c.HP, c.MaxHP)
		if c.IsDead() {
			return
		}
	}
}

func accessInventory(c *Character, reader *bufio.Reader) {
	for {
		fmt.Println("\nInventaire")
		if len(c.Inventory) == 0 {
			fmt.Println("(vide)")
		} else {
			for i, item := range c.Inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}
		}
		fmt.Println("\nOptions :")
		fmt.Println("u - Utiliser un objet")
		fmt.Println("e - Équiper un objet")
		fmt.Println("m - Aller voir le marchand")
		fmt.Println("b - Retour")
		fmt.Print("Choix > ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "b" {
			return
		}
		if choice == "u" {
			if len(c.Inventory) == 0 {
				fmt.Println("Aucun objet à utiliser.")
				continue
			}
			fmt.Print("Numéro de l'objet à utiliser : ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			idx, err := strconv.Atoi(numStr)
			if err != nil || idx < 1 || idx > len(c.Inventory) {
				fmt.Println("Numéro invalide.")
				continue
			}
			item := c.Inventory[idx-1]

			if strings.Contains(strings.ToLower(item), "potion de vie") {
				takePotion(c)
			} else if strings.Contains(strings.ToLower(item), "potion de poison") {
				removeInventory(c, item)
				poisonPot(c)
			} else if strings.Contains(strings.ToLower(item), "livre de sort : boule de feu") {
				removeInventory(c, item)
				spellBook(c)
			} else if strings.Contains(strings.ToLower(item), "amélioration d'inventaire") {
				removeInventory(c, item)
				c.upgradeInventorySlot()
			} else {
				fmt.Printf("L'utilisation de %s n'est pas encore implémentée.\n", item)
			}
		}
		if choice == "e" {
			if len(c.Inventory) == 0 {
				fmt.Println("Aucun objet à équiper.")
				continue
			}
			fmt.Print("Numéro de l'objet à équiper : ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			idx, err := strconv.Atoi(numStr)
			if err != nil || idx < 1 || idx > len(c.Inventory) {
				fmt.Println("Numéro invalide.")
				continue
			}
			item := c.Inventory[idx-1]
			c.equip(item)
		}
		if choice == "m" {
			marchand(c, reader)
		}
	}
}
