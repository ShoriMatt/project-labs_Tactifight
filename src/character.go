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
		Initiative:        Initiative,
		MaxMana:           MaxMana,
		Mana:              Mana,
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
	centerText(fmt.Sprintf("%s a gagné %d points d’expérience !", c.Name, amount))
	c.XP += amount
	for c.XP >= c.XPToNext {
		c.XP -= c.XPToNext
		c.Level++
		c.MaxHP += 10
		c.HP = c.MaxHP
		c.XPToNext = int(float64(c.XPToNext) * 1.5)
		centerText(fmt.Sprintf("✨ %s passe au niveau %d ! PV max +10 (%d PV)", c.Name, c.Level, c.MaxHP))
	}
}

func (c *Character) recalcMaxHP() {
	baseHP := map[string]int{"Elfe": 80, "Humain": 100, "Nain": 120}[c.Class]
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
		centerText(fmt.Sprintf("Impossible d'équiper %s.", item))
		return
	}
	if !removeInventory(c, itemLower) {
		centerText("Objet non trouvé dans l'inventaire.")
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
	centerText(fmt.Sprintf("%s a été équipé.", item))
	c.recalcMaxHP()
}

func (c *Character) upgradeInventorySlot() {
	if c.UpgradeCount < 3 {
		c.InventoryCapacity += 10
		c.UpgradeCount++
		centerText(fmt.Sprintf("Inventaire augmenté ! Nouvelle capacité : %d (Améliorations utilisées : %d/3)",
			c.InventoryCapacity, c.UpgradeCount))
	} else {
		centerText("Vous avez déjà utilisé toutes vos améliorations d'inventaire (3/3) !")
	}
}

func (c *Character) IsDead() bool {
	if c.HP <= 0 {
		centerText(fmt.Sprintf("%s est mort...", c.Name))
		c.HP = c.MaxHP / 2
		centerText(fmt.Sprintf("%s est ressuscité avec %d/%d PV !", c.Name, c.HP, c.MaxHP))
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
			centerText("Erreur : le nom ne doit contenir que des lettres, espaces ou tirets.")
		}
	}
	centerText("\nChoisis ta classe :")
	centerText("1 - Assasin (80 PV max)")
	centerText("2 - Gladiateur (100 PV max)")
	centerText("3 - Tank (120 PV max)")
	fmt.Print("Choix > ")
	classe, _ := reader.ReadString('\n')
	classe = strings.TrimSpace(classe)

	var class string
	var MaxHP, HP int
	var MaxMana, Mana int
	var Initiative int
	switch classe {
	case "1":
		class, MaxHP, HP, MaxMana, Mana, Initiative = "Assasin", 80, 40, 100, 50, 15
	case "2":
		class, MaxHP, HP, MaxMana, Mana, Initiative = "Gladiateur", 100, 50, 80, 40, 10
	case "3":
		class, MaxHP, HP, MaxMana, Mana, Initiative = "Tank", 120, 60, 70, 35, 5
	default:
		centerText("Choix invalide. Classe par défaut : Gladiateur")
		class, MaxHP, HP, MaxMana, Mana, Initiative = "gladiateur", 100, 50, 80, 40, 10
	}

	initialInventory := []string{"potion de vie", "potion de vie", "potion de vie"}
	return initCharacter(name, class, 1, MaxHP, HP, Initiative, MaxMana, Mana, initialInventory)
}

func displayInfo(c *Character) {
	centerText("\n===========================================")
	centerText("Informations du personnage")
	centerText(fmt.Sprintf("Nom        : %s", c.Name))
	centerText(fmt.Sprintf("Classe     : %s", c.Class))
	centerText(fmt.Sprintf("Niveau     : %d", c.Level))
	centerText(fmt.Sprintf("PV         : %d / %d", c.HP, c.MaxHP))
	centerText(fmt.Sprintf("Mana       : %d / %d", c.Mana, c.MaxMana))
	centerText(fmt.Sprintf("XP         : %d / %d", c.XP, c.XPToNext))
	centerText(fmt.Sprintf("Initiative : %d", c.Initiative))
	centerText(fmt.Sprintf("Or         : %d", c.Gold))
	centerText(fmt.Sprintf("Inventaire : %d/%d item(s)", len(c.Inventory), c.InventoryCapacity))
	if len(c.Inventory) > 0 {
		for i, it := range c.Inventory {
			centerText(fmt.Sprintf("  %d. %s", i+1, formatNom(it)))
		}
	}
	if len(c.Skills) > 0 {
		centerText("Sorts      : " + strings.Join(c.Skills, ", "))
	}
	centerText(fmt.Sprintf("Équipement : Tête [%s], Torse [%s], Pieds [%s]",
		c.Equipment.Head, c.Equipment.Torso, c.Equipment.Feet))
	centerText("===========================================\n")
}
