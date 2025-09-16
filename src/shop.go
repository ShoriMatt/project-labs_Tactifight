package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func marchand(c *Character, reader *bufio.Reader) {
	inventaire := []string{
		"potion de vie",
		"potion de poison",
		"potion de mana",
		"livre de sort : boule de feu",
		"fourrure de loup",
		"peau de troll",
		"cuir de sanglier",
		"plume de corbeau",
		"amélioration d'inventaire",
	}
	prix := []int{3, 6, 5, 25, 4, 7, 3, 1, 30}

	for {
		centerText(` _, _  _, __,  _, _,_  _, _, _ __,
 |\/| /_\ |_) / '` + ` |_| /_\ |\ | | \
 |  | | | | \ \ , | | | | | \| |_/
 ~  ~ ~ ~ ~ ~  ~  ~ ~ ~ ~ ~  ~ ~  
								  `)
		centerText("\n===========================================")
		centerText(fmt.Sprintf("Vous avez %d or.", c.Gold))
		for i, item := range inventaire {
			centerText(fmt.Sprintf("%d. %s (%d or)", i+1, item, prix[i]))
		}
		centerText("Tapez le numéro de l’objet à acheter, ou 'q' pour quitter.")
		centerText("===========================================\n")
		fmt.Print("Choix : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if strings.ToLower(choice) == "q" {
			centerText("Vous quittez le marchand.")
			return
		}

		idx, err := strconv.Atoi(choice)
		if err != nil || idx < 1 || idx > len(inventaire) {
			centerText("Numéro invalide.")
			continue
		}

		item := inventaire[idx-1]
		prixItem := prix[idx-1]

		if len(c.Inventory) >= c.InventoryCapacity &&
			item != "amélioration d'inventaire" &&
			item != "livre de sort : boule de feu" {
			centerText(fmt.Sprintf("Inventaire plein ! Vous ne pouvez pas acheter cet objet (max %d).", c.InventoryCapacity))
			continue
		}

		if c.Gold < prixItem {
			centerText(fmt.Sprintf("Pas assez d’or ! (%s coûte %d or, il vous manque %d)", item, prixItem, prixItem-c.Gold))
			continue
		}

		c.Gold -= prixItem

		switch item {
		case "livre de sort : boule de feu":
			spellBook(c)
			centerText(fmt.Sprintf("Vous avez acheté et appris directement : %s (-%d or)", item, prixItem))
		case "amélioration d'inventaire":
			c.upgradeInventorySlot()
			centerText(fmt.Sprintf("Vous avez acheté et utilisé : %s (-%d or)", item, prixItem))
		default:
			addInventory(c, item)
			centerText(fmt.Sprintf("Vous avez acheté : %s (-%d or)", item, prixItem))
		}

		centerText(fmt.Sprintf("Or restant : %d", c.Gold))
	}
}

func forgeron(c *Character, reader *bufio.Reader) {
	inventaire := []string{"Chapeau de l’aventurier", "Tunique de l’aventurier", "Bottes de l’aventurier"}
	MatériauxChapeau := []string{"Plume de Corbeau", "Cuir de Sanglier"}
	MatériauxTunique := []string{"Fourrure de loup", "Fourrure de loup", "Peau de Troll"}
	MatériauxBottes := []string{"Fourrure de loup", "Cuir de Sanglier"}
	prix := []int{10, 10, 10}

	centerText(`
 __,  _, __,  _, __, __,  _, _, _
 |_  / \ |_) / _ |_  |_) / \ |\ |
 |   \ / | \ \ / |   | \ \ / | \|
 ~    ~  ~ ~  ~  ~~~ ~ ~  ~  ~  ~
								 `)
	centerText("\n===========================================")
	for i, item := range inventaire {
		centerText(fmt.Sprintf("%d. %s (%d or)", i+1, item, prix[i]))
	}
	centerText("===========================================\n")
	centerText("Voulez-vous fabriquer un item ? (o/n)")
	fmt.Print("Choix : ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "o" {
		fmt.Print("Numéro de l'objet à fabriquer : ")
		numStr, _ := reader.ReadString('\n')
		numStr = strings.TrimSpace(numStr)
		idx, err := strconv.Atoi(numStr)
		if err != nil || idx < 1 || idx > len(inventaire) {
			centerText("Numéro invalide.")
			return
		}
		item := inventaire[idx-1]
		prixItem := prix[idx-1]
		var mat []string
		if idx-1 == 0 {
			mat = MatériauxChapeau
		} else if idx-1 == 1 {
			mat = MatériauxTunique
		} else {
			mat = MatériauxBottes
		}
		centerText(fmt.Sprintf("Matériaux requis : %v", mat))

		for _, m := range mat {
			found := false
			for _, inv := range c.Inventory {
				if strings.EqualFold(inv, m) {
					found = true
					break
				}
			}
			if !found {
				centerText(fmt.Sprintf("Il vous manque le matériau : %s", m))
				return
			}
		}
		for _, m := range mat {
			for i, inv := range c.Inventory {
				if strings.EqualFold(inv, m) {
					c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
					break
				}
			}
		}
		if len(c.Inventory) >= c.InventoryCapacity {
			centerText("Inventaire plein ! Vous ne pouvez pas fabriquer cet objet.")
			return
		}

		if c.Gold < prixItem {
			centerText("Pas assez d'or !")
			return
		}

		c.Gold -= prixItem
		addInventory(c, item)
		centerText(fmt.Sprintf("Vous avez fabriqué : %s (-%d or)", item, prixItem))
	}
}
