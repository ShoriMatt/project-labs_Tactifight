package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	HP        int
	Inventory []string
	Skills    []string
	Gold      int
	Equipment Equipment
}

func initCharacter(name string, class string, level int, maxHP int, currentHP int, inventory []string) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		HP:        currentHP,
		Inventory: inventory,
		Skills:    []string{"Coup de poing"},
		Gold:      100,
		Equipment: Equipment{"", "", ""},
	}
}

func displayInfo(c *Character) {
	fmt.Println("\nInformations du personnage")
	fmt.Printf("Nom        : %s\n", c.Name)
	fmt.Printf("Classe     : %s\n", c.Class)
	fmt.Printf("Niveau     : %d\n", c.Level)
	fmt.Printf("PV         : %d / %d\n", c.HP, c.MaxHP)
	fmt.Printf("Or         : %d\n", c.Gold)
	fmt.Printf("Inventaire : %d item(s)\n", len(c.Inventory))
	if len(c.Inventory) > 0 {
		fmt.Println("  Items :", strings.Join(c.Inventory, ", "))
	}
	if len(c.Skills) > 0 {
		fmt.Println("Sorts      :", strings.Join(c.Skills, ", "))
	}
	fmt.Println("=================================")
}

func addInventory(c *Character, item string) {
	c.Inventory = append(c.Inventory, item)
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
		if strings.Contains(strings.ToLower(v), "potion") {
			found = i
			break
		}
	}
	if found == -1 {
		fmt.Println("Aucune potion disponible.")
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
		fmt.Println("m - Marchand")
		fmt.Println("b - Retour")
		fmt.Print("Choix > ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "b":
			return

		case "u":
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
			if strings.Contains(strings.ToLower(item), "potion") {
				takePotion(c)
			} else {
				fmt.Printf("L'utilisation de %s n'est pas encore implémentée.\n", item)
			}
		case "m":
			Marchand(c, reader)
		}
	}
}

func Marchand(c *Character, reader *bufio.Reader) {
	inventaire := []string{"potion de vie"}
	if len(inventaire) == 0 {
		fmt.Println("Le marchand n'a rien a vendre")
		return
	} else {
		for i, item := range inventaire {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}
	fmt.Println("voulez vous achetez un item")
	fmt.Println("o / n")
	fmt.Print("choix :")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "n":
		return
	case "o":
		fmt.Print("Numéro de l'objet à acheter : ")
		numStr, _ := reader.ReadString('\n')
		numStr = strings.TrimSpace(numStr)
		idx, err := strconv.Atoi(numStr)
		if err != nil || idx < 1 || idx > len(inventaire) {
			fmt.Println("Numéro invalide.")
			return
		}
		item := inventaire[idx-1]
		if strings.Contains(strings.ToLower(item), "potion de vie") {
			addInventory(c, item)
			fmt.Printf("Vous avez acheté : %s\n", item)
		}
	}
}

func mainMenu(c *Character, reader *bufio.Reader) {
	for {
		fmt.Println("\nMenu Principal")
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder au contenu de l'inventaire")
		fmt.Println("3 - Quitter")
		fmt.Print("Choix > ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			displayInfo(c)
		case "2":
			accessInventory(c, reader)
		case "3":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Projet RED - Test")
	fmt.Print("Entre ton nom (ou Entrée pour 'Joueur') : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Joueur"
	}

	initialInventory := []string{"Potion de vie", "Potion de vie", "Potion de vie"}
	c1 := initCharacter(name, "Elfe", 1, 100, 40, initialInventory)

	fmt.Printf("Personnage créé : %s (%s) - PV %d/%d - %d potions\n", c1.Name, c1.Class, c1.HP, c1.MaxHP, len(c1.Inventory))
	mainMenu(&c1, reader)
}
