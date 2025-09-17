package main

import (
	"fmt"
	"math/rand"
)

type Monster struct {
	Name       string
	HP         int
	MaxHP      int
	XPReward   int
	GoldReward int
	Initiative int
}

func (m *Monster) IsDead() bool {
	return m.HP <= 0
}

func initChameauToxiqueLunaire() Monster {
	return Monster{
		Name:       "Chameau toxique lunaire",
		HP:         30,
		MaxHP:      30,
		XPReward:   10,
		GoldReward: 5,
		Initiative: rand.Intn(100),
	}
}

func ChameauToxiqueLunairePattern(cha *Monster, c *Character, turn int) {
	random := rand.Intn(100)
	if random < 20 && c.PoisonTurns == 0 {
		c.PoisonTurns = 3
		centerText(fmt.Sprintf("%s vous empoisonne pour 3 tours ! Avec l'attaque Crachat stellaire", cha.Name))
		return
	}
	damage := 3
	c.HP -= damage
	if c.HP < 0 {
		c.HP = 0
	}
	centerText(fmt.Sprintf("%s attaque et inflige %d dégâts !", cha.Name, damage))
	centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
}

func initCloneBancalDeMacron() Monster {
	return Monster{
		Name:       "Clone Bancal de Macron",
		HP:         25,
		MaxHP:      25,
		XPReward:   15,
		GoldReward: 7,
		Initiative: rand.Intn(100),
	}
}

func CloneBancalDeMacronPattern(clo *Monster, c *Character, turn int) {
	Mob := *clo
	random := rand.Intn(100)
	if random < 50 {
		damage := 1
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s utilise discour infinie ! et vous fait sauter votre tour en vous infligeant %d dégâts", clo.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		CloneBancalDeMacronPattern(&Mob, c, turn)
	}
	damage := 4
	c.HP -= damage
	if c.HP < 0 {
		c.HP = 0
	}
	centerText(fmt.Sprintf("%s attaque et inflige %d dégâts !", clo.Name, damage))
	centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
}

func initGardePrésidentielSpatial() Monster {
	return Monster{
		Name:       "Garde Présidentiel Spatial",
		HP:         40,
		MaxHP:      40,
		XPReward:   20,
		GoldReward: 10,
		Initiative: rand.Intn(100),
	}
}

func GardePrésidentielSpatialPattern(gar *Monster, c *Character, turn int) {
	random := rand.Intn(100)
	if random < 30 {
		damage := 20
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s vous attaque avec son sceptre-laser ! et inflige %d dégâts !", gar.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	}
	damage := 5
	c.HP -= damage
	if c.HP < 0 {
		c.HP = 0
	}
	centerText(fmt.Sprintf("%s attaque et inflige %d dégâts !", gar.Name, damage))
	centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
}

func initEmmanuelMacroniusIV() Monster {
	return Monster{
		Name:       "Emmanuel Macronius IV",
		HP:         100,
		MaxHP:      100,
		XPReward:   100,
		GoldReward: 50,
		Initiative: rand.Intn(100),
	}
}

func EmmanuelMacroniusIVPattern(boss *Monster, c *Character, turn int) {
	damage := 8
	c.HP -= damage
	if c.HP < 0 {
		c.HP = 0
	}
	centerText(fmt.Sprintf("%s attaque et inflige %d dégâts !", boss.Name, damage))
	centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
}
