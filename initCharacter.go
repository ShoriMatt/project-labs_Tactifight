package character

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
