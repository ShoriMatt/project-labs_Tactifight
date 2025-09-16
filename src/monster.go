package main

import (
	"fmt"
)

type Monster struct {
	Name     string
	HP       int
	MaxHP    int
	XPReward int
}

func (m *Monster) IsDead() bool {
	return m.HP <= 0
}

func initGoblin() Monster {
	return Monster{
		Name:     "Gobelin d'entraînement",
		HP:       30,
		MaxHP:    30,
		XPReward: 10,
	}
}

func goblinPattern(g *Monster, c *Character, turn int) {
	damage := 3
	c.HP -= damage
	if c.HP < 0 {
		c.HP = 0
	}

	centerText(fmt.Sprintf("%s attaque et inflige %d dégâts !", g.Name, damage))
	centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
}
