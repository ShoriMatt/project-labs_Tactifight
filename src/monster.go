package main

import (
	"fmt"
	"math/rand/v2"
)

type Monster struct {
	Name       string
	MaxHP      int
	HP         int
	AttackPts  int
	XPReward   int
	Initiative int
}

func initGoblin() Monster {
	return Monster{
		Name:       "Gobelin d'entraînement",
		MaxHP:      40,
		HP:         40,
		AttackPts:  5,
		XPReward:   20,
		Initiative: rand.IntN(100),
	}
}

func goblinPattern(goblin *Monster, player *Character, turn int) {
	damage := goblin.AttackPts
	if turn%3 == 0 {
		damage *= 2
	}
	player.HP -= damage
	if player.HP < 0 {
		player.HP = 0
	}

	centerText(fmt.Sprintf("%s inflige à %s %d de dégâts", goblin.Name, player.Name, damage))
	centerText(fmt.Sprintf("%s : %d/%d PV", player.Name, player.HP, player.MaxHP))

	player.IsDead()
}
