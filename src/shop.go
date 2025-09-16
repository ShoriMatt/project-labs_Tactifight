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
		fmt.Println(` _, _  _, __,  _, _,_  _, _, _ __,
 |\/| /_\ |_) / '` + ` |_| /_\ |\ | | \
 |  | | | | \ \ , | | | | | \| |_/
 ~  ~ ~ ~ ~ ~  ~  ~ ~ ~ ~ ~  ~ ~  
								  `)
		fmt.Printf("Vous avez %d or.\n", c.Gold)
		for i, item := range inventaire {
			fmt.Printf("%d. %s (%d or)\n", i+1, item, prix[i])
		}
		fmt.Println("Tapez le numéro de l’objet à acheter, ou 'q' pour quitter.")
		fmt.Print("Choix : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if strings.ToLower(choice) == "q" {
			fmt.Println("Vous quittez le marchand.")
			return
		}

		idx, err := strconv.Atoi(choice)
		if err != nil || idx < 1 || idx > len(inventaire) {
			fmt.Println("Numéro invalide.")
			continue
		}

		item := inventaire[idx-1]
		prixItem := prix[idx-1]

		if len(c.Inventory) >= c.InventoryCapacity &&
			item != "amélioration d'inventaire" &&
			item != "livre de sort : boule de feu" {
			fmt.Printf("Inventaire plein ! Vous ne pouvez pas acheter cet objet (max %d).\n", c.InventoryCapacity)
			continue
		}

		if c.Gold < prixItem {
			fmt.Printf("Pas assez d’or ! (%s coûte %d or, il vous manque %d)\n",
				item, prixItem, prixItem-c.Gold)
			continue
		}

		c.Gold -= prixItem

		switch item {
		case "livre de sort : boule de feu":
			spellBook(c)
			fmt.Printf("Vous avez acheté et appris directement : %s (-%d or)\n", item, prixItem)

		case "amélioration d'inventaire":
			c.upgradeInventorySlot()
			fmt.Printf("Vous avez acheté et utilisé : %s (-%d or)\n", item, prixItem)

		default:
			addInventory(c, item)
			fmt.Printf("Vous avez acheté : %s (-%d or)\n", item, prixItem)
		}

		fmt.Printf("Or restant : %d\n", c.Gold)
	}
}

func forgeron(c *Character, reader *bufio.Reader) {
	inventaire := []string{"Chapeau de l’aventurier", "Tunique de l’aventurier", "Bottes de l’aventurier"}
	MatériauxChapeau := []string{"Plume de Corbeau", "Cuir de Sanglier"}
	MatériauxTunique := []string{"Fourrure de loup", "Fourrure de loup", "Peau de Troll"}
	MatériauxBottes := []string{"Fourrure de loup", "Cuir de Sanglier"}
	prix := []int{10, 10, 10}

	fmt.Println(`
 __,  _, __,  _, __, __,  _, _, _
 |_  / \ |_) / _ |_  |_) / \ |\ |
 |   \ / | \ \ / |   | \ \ / | \|
 ~    ~  ~ ~  ~  ~~~ ~ ~  ~  ~  ~
								 `)
	for i, item := range inventaire {
		fmt.Printf("%d. %s (%d or)\n", i+1, item, prix[i])
	}
	fmt.Println("Voulez-vous fabriquer un item ? (o/n)")
	fmt.Print("Choix : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "o" {
		fmt.Print("Numéro de l'objet à fabriquer : ")
		numStr, _ := reader.ReadString('\n')
		numStr = strings.TrimSpace(numStr)
		idx, err := strconv.Atoi(numStr)
		if err != nil || idx < 1 || idx > len(inventaire) {
			fmt.Println("Numéro invalide.")
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
		fmt.Println("Matériaux requis :", mat)

		for _, m := range mat {
			found := false
			for _, inv := range c.Inventory {
				if strings.EqualFold(inv, m) {
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("Il vous manque le matériau : %s\n", m)
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
			fmt.Println("Inventaire plein ! Vous ne pouvez pas fabriquer cet objet.")
			return
		}

		if c.Gold < prixItem {
			fmt.Println("Pas assez d'or !")
			return
		}

		c.Gold -= prixItem
		addInventory(c, item)
		fmt.Printf("Vous avez fabriqué : %s (-%d or)\n", item, prixItem)
	}
}
