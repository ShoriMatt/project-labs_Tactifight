package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Character struct {
	Name              string
	Class             string
	Level             int
	MaxHP             int
	HP                int
	Inventory         []string
	Skills            []string
	Gold              int
	Equipment         Equipment
	InventoryCapacity int
	UpgradeCount      int
}

type Monster struct {
	Name      string
	MaxHP     int
	HP        int
	AttackPts int
}

func initGoblin() Monster {
	return Monster{
		Name:      "Gobelin d'entraînement",
		MaxHP:     40,
		HP:        40,
		AttackPts: 5,
	}
}

func initCharacter(name string, class string, level int, maxHP int, currentHP int, inventory []string) Character {
	return Character{
		Name:              name,
		Class:             class,
		Level:             level,
		MaxHP:             maxHP,
		HP:                currentHP,
		Inventory:         inventory,
		Skills:            []string{"Coup de poing"},
		Gold:              100,
		Equipment:         Equipment{Head: "", Torso: "", Feet: ""},
		InventoryCapacity: 10,
		UpgradeCount:      0,
	}
}

func (c *Character) upgradeInventorySlot() {
	if c.UpgradeCount < 3 {
		c.InventoryCapacity += 10
		c.UpgradeCount++
		fmt.Printf("Inventaire augmenté ! Nouvelle capacité : %d (Améliorations utilisées : %d/3)\n",
			c.InventoryCapacity, c.UpgradeCount)
	} else {
		fmt.Println("Vous avez déjà utilisé toutes vos améliorations d'inventaire (3/3) !")
	}
}

func IsAlpha(s string) bool {
	for _, r := range s {
		if !(unicode.IsLetter(r) || r == ' ' || r == '-') {
			return false
		}
	}
	return true
}

func formatNom(nom string) string {
	if len(nom) == 0 {
		return ""
	}
	nom = strings.ToLower(nom)
	r := []rune(nom)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func characterCreation(reader *bufio.Reader) Character {
	var name string

	for {
		fmt.Print("Entre ton nom (ou Entrée pour 'Joueur') : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			name = "Joueur"
			break
		}

		if IsAlpha(input) {
			name = formatNom(input)
			break
		} else {
			fmt.Println("Erreur : le nom ne doit contenir que des lettres, espaces ou tirets.")
		}
	}

	fmt.Println("\nChoisis ta classe :")
	fmt.Println("1 - Elfe (80 PV max)")
	fmt.Println("2 - Humain (100 PV max)")
	fmt.Println("3 - Nain (120 PV max)")
	fmt.Print("Choix > ")
	classe, _ := reader.ReadString('\n')
	classe = strings.TrimSpace(classe)

	var class string
	var MaxHP, HP int

	switch classe {
	case "1":
		class = "Elfe"
		MaxHP = 80
		HP = 40
	case "2":
		class = "Humain"
		MaxHP = 100
		HP = 50
	case "3":
		class = "Nain"
		MaxHP = 120
		HP = 60
	default:
		fmt.Println("Choix invalide.")
		fmt.Println("Classe par défaut : Humain")
		class = "Humain"
		MaxHP = 100
		HP = 50
	}

	initialInventory := []string{"potion de vie", "potion de vie", "potion de vie"}
	return initCharacter(name, class, 1, MaxHP, HP, initialInventory)
}

func displayInfo(c *Character) {
	fmt.Println("\nInformations du personnage")
	fmt.Printf("Nom        : %s\n", c.Name)
	fmt.Printf("Classe     : %s\n", c.Class)
	fmt.Printf("Niveau     : %d\n", c.Level)
	fmt.Printf("PV         : %d / %d\n", c.HP, c.MaxHP)
	fmt.Printf("Or         : %d\n", c.Gold)
	fmt.Printf("Inventaire : %d/%d item(s)\n", len(c.Inventory), c.InventoryCapacity)

	if len(c.Inventory) > 0 {
		for i, it := range c.Inventory {
			fmt.Printf("  %d. %s\n", i+1, formatNom(it))
		}
	}

	if len(c.Skills) > 0 {
		fmt.Println("Sorts      :", strings.Join(c.Skills, ", "))
	}
	fmt.Printf("Équipement : Tête [%s], Torse [%s], Pieds [%s]\n",
		c.Equipment.Head, c.Equipment.Torso, c.Equipment.Feet)
	fmt.Println("=================================")
}

func (c *Character) IsDead() bool {
	if c.HP <= 0 {
		fmt.Printf("%s est mort...\n", c.Name)
		c.HP = c.MaxHP / 2
		fmt.Printf("%s est ressuscité avec %d/%d PV !\n", c.Name, c.HP, c.MaxHP)
		return true
	}
	return false
}

func spellBook(c *Character) {
	spell := "Boule de feu"

	for _, s := range c.Skills {
		if s == spell {
			fmt.Println("Vous connaissez déjà le sort Boule de feu !")
			return
		}
	}

	c.Skills = append(c.Skills, spell)
	fmt.Println("Nouveau sort appris :", spell)
}

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
		if choice == "m" {
			marchand(c, reader)
		}
	}
}

func marchand(c *Character, reader *bufio.Reader) {
	inventaire := []string{
		"potion de vie",
		"potion de poison",
		"livre de sort : boule de feu",
		"fourrure de loup",
		"peau de troll",
		"cuir de sanglier",
		"plume de corbeau",
		"amélioration d'inventaire",
	}
	prix := []int{3, 6, 25, 4, 7, 3, 1, 30}

	for {
		fmt.Println("\n--- Marchand ---")
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

	fmt.Println("\n--- Forgeron ---")
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

func goblinPattern(goblin *Monster, player *Character, turns int) {
	for turn := 1; turn <= turns; turn++ {
		damage := goblin.AttackPts
		if turn%3 == 0 {
			damage *= 2
		}
		player.HP -= damage
		if player.HP < 0 {
			player.HP = 0
		}
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, damage)
		fmt.Printf("%s : %d/%d PV\n\n", player.Name, player.HP, player.MaxHP)
		if player.IsDead() {
			break
		}
	}
}

func mainMenu(c *Character, reader *bufio.Reader) {
	for {
		fmt.Println("\nMenu Principal")
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder au contenu de l'inventaire")
		fmt.Println("3 - Voir le Forgeron")
		fmt.Println("4 - Combattre le Gobelin d'entraînement")
		fmt.Println("5 - Quitter")
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
			goblinPattern(&goblin, c, 6)
		case "5":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Projet RED - Fusion")

	c1 := characterCreation(reader)

	fmt.Printf("Personnage créé : %s (%s) - PV %d/%d - %d potions\n",
		c1.Name, c1.Class, c1.HP, c1.MaxHP, len(c1.Inventory))

	mainMenu(&c1, reader)
}
