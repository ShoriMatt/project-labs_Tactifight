package main

import (
	"bufio"
	"fmt"
	"strings"
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
	Initiative        int
	MaxMana           int
	Mana              int
	Inventory         []string
	Skills            []string
	Gold              int
	Equipment         Equipment
	InventoryCapacity int
	UpgradeCount      int
	XP                int
	XPToNext          int
	PoisonTurns       int 
}

func initCharacter(name string, class string, level int, maxHP int, currentHP int, Initiative int, MaxMana int, Mana int, inventory []string) Character {
	return Character{
		Name:              name,
		Class:             class,
		Level:             level,
		MaxHP:             maxHP,
		HP:                currentHP,
		MaxMana:           MaxMana,
		Mana:              Mana,
		Initiative:        Initiative,
		Inventory:         inventory,
		Skills:            []string{"Coup de poing"},
		Gold:              100,
		Equipment:         Equipment{Head: "", Torso: "", Feet: ""},
		InventoryCapacity: 10,
		UpgradeCount:      0,
		XP:                0,
		XPToNext:          50,
	}
}

func (c *Character) gainXP(amount int) {
	fmt.Printf("%s a gagné %d points d’expérience !\n", c.Name, amount)
	c.XP += amount

	for c.XP >= c.XPToNext {
		c.XP -= c.XPToNext
		c.Level++
		c.MaxHP += 10
		c.HP = c.MaxHP
		c.XPToNext = int(float64(c.XPToNext) * 1.5)
		fmt.Printf("✨ %s passe au niveau %d ! PV max +10 (%d PV)\n", c.Name, c.Level, c.MaxHP)
	}
}

func (c *Character) recalcMaxHP() {
	baseHP := 0
	switch c.Class {
	case "Elfe":
		baseHP = 80
	case "Humain":
		baseHP = 100
	case "Nain":
		baseHP = 120
	}

	bonus := 0
	if c.Equipment.Head == "Chapeau de l’aventurier" {
		bonus += 10
	}
	if c.Equipment.Torso == "Tunique de l’aventurier" {
		bonus += 25
	}
	if c.Equipment.Feet == "Bottes de l’aventurier" {
		bonus += 15
	}

	c.MaxHP = baseHP + bonus
	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
}

func (c *Character) equip(item string) {
	itemLower := strings.ToLower(item)
	slot := ""
	switch itemLower {
	case "chapeau de l’aventurier":
		slot = "Head"
	case "tunique de l’aventurier":
		slot = "Torso"
	case "bottes de l’aventurier":
		slot = "Feet"
	default:
		fmt.Printf("Impossible d'équiper %s.\n", item)
		return
	}

	if !removeInventory(c, itemLower) {
		fmt.Println("Objet non trouvé dans l'inventaire.")
		return
	}

	switch slot {
	case "Head":
		if c.Equipment.Head != "" {
			addInventory(c, c.Equipment.Head)
		}
		c.Equipment.Head = item
	case "Torso":
		if c.Equipment.Torso != "" {
			addInventory(c, c.Equipment.Torso)
		}
		c.Equipment.Torso = item
	case "Feet":
		if c.Equipment.Feet != "" {
			addInventory(c, c.Equipment.Feet)
		}
		c.Equipment.Feet = item
	}

	fmt.Printf("%s a été équipé.\n", item)
	c.recalcMaxHP()
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

func (c *Character) IsDead() bool {
	if c.HP <= 0 {
		fmt.Printf("%s est mort...\n", c.Name)
		c.HP = c.MaxHP / 2
		fmt.Printf("%s est ressuscité avec %d/%d PV !\n", c.Name, c.HP, c.MaxHP)
		return true
	}
	return false
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
	var MaxMana, Mana int
	var Initiative int

	switch classe {
	case "1":
		class = "Elfe"
		MaxHP = 80
		HP = 40
		MaxMana = 100
		Mana = 50
		Initiative = 15
	case "2":
		class = "Humain"
		MaxHP = 100
		HP = 50
		MaxMana = 80
		Mana = 40
		Initiative = 10
	case "3":
		class = "Nain"
		MaxHP = 120
		HP = 60
		MaxMana = 70
		Mana = 35
		Initiative = 5
	default:
		fmt.Println("Choix invalide.")
		fmt.Println("Classe par défaut : Humain")
		class = "Humain"
		MaxHP = 100
		HP = 50
		Initiative = 10
	}

	initialInventory := []string{"potion de vie", "potion de vie", "potion de vie"}
	return initCharacter(name, class, 1, MaxHP, HP, MaxMana, Mana, Initiative, initialInventory)
}

func displayInfo(c *Character) {
	fmt.Println("\nInformations du personnage")
	fmt.Printf("Nom        : %s\n", c.Name)
	fmt.Printf("Classe     : %s\n", c.Class)
	fmt.Printf("Niveau     : %d\n", c.Level)
	fmt.Printf("PV         : %d / %d\n", c.HP, c.MaxHP)
	fmt.Printf("PV         : %d / %d\n", c.Mana, c.MaxMana)
	fmt.Printf("XP         : %d / %d\n", c.XP, c.XPToNext)
	fmt.Printf("Initiative : %d\n", c.Initiative)
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
