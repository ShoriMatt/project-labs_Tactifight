package main

import (
	"fmt"
	"math/rand/v2"
)

type Monster struct {
<<<<<<< HEAD
	Name      string
	MaxHP     int
	HP        int
	AttackPts int
	XPReward  int
=======
	Name       string
	MaxHP      int
	HP         int
	AttackPts  int
	Initiative int
>>>>>>> ff37773 (add initiative)
}

func initGoblin() Monster {
	return Monster{
<<<<<<< HEAD
		Name:      "Gobelin d'entraînement",
		MaxHP:     40,
		HP:        40,
		AttackPts: 5,
		XPReward:  20,
=======
		Name:       "Gobelin d'entraînement",
		MaxHP:      40,
		HP:         40,
		AttackPts:  5,
		Initiative: rand.IntN(100),
>>>>>>> ff37773 (add initiative)
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
			return
		}
	}
	fmt.Printf("\n--- %s est vaincu ! ---\n", goblin.Name)
	player.gainXP(goblin.XPReward)
}
