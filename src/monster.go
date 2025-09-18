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
		HP:         150,
		MaxHP:      150,
		XPReward:   200,
		GoldReward: 50,
		Initiative: rand.Intn(100),
	}
}

func EmmanuelMacroniusIVPattern(boss *Monster, c *Character, turn int) {
	random := rand.Intn(100)
	if random < 30 {
		damage := 30
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s vous attaque avec son armés ! et inflige %d dégâts !", boss.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	} else if random < 60 {
		damage := 50
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s utilise 49.3 ! et inflige %d dégâts !", boss.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	} else if random < 80 {
		damage := 40
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s utilise la réforme des retraites ! et inflige %d dégâts !", boss.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	}
	damage := 8
	c.HP -= damage
	if c.HP < 0 {
		c.HP = 0
	}
	centerText(fmt.Sprintf("%s attaque et inflige %d dégâts !", boss.Name, damage))
	centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
}
func initBrigite() Monster {
	return Monster{
		Name:       "Brigite la Broyeuse",
		HP:         500,
		MaxHP:      500,
		XPReward:   1000,
		GoldReward: 300,
		Initiative: rand.Intn(100),
	}
}

func BrigitePattern(boss_2 *Monster, c *Character, turn int) {
	Mob := *boss_2
	random := rand.Intn(100)

	// Attaque spéciale débloquée si Brigite est à moins de 50% PV
	if boss_2.HP <= boss_2.MaxHP/2 && random < 5 {
		damage := 200
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s utilise : Dynastie Élyséenne → combo spécial : elle appelle Emmanuel Macronius V pour un coup combiné ! Vous subissez %d dégâts !!!", boss_2.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	} else if random < 15 {
		random2 := rand.Intn(100)
		damage := 1
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s utilise Coup de Sac Vuitton ! peut vous faire sauter votre tour en vous infligeant %d dégâts", boss_2.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		if random2 < 50 {
			CloneBancalDeMacronPattern(&Mob, c, turn)
		}
		return
	} else if random < 35 {
		damage := 35
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s utilise Frappe Première Dame ! Vous êtes projeté et subissez %d dégâts !", boss_2.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	} else if random < 50 {
		damage := 50
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s vous attaque avec sa Main de Fer dans un Gant Dior” → attaque lourde qui inflige %d dégâts avec un style raffiné.", boss_2.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	} else if random < 60 {
		heal := 20 + rand.Intn(21) // rend 20 à 30 PV
		boss_2.HP += heal
		if boss_2.HP > boss_2.MaxHP {
			boss_2.HP = boss_2.MaxHP
		}
		centerText(fmt.Sprintf("%s utilise une Retouche Esthétique et récupère %d PV !", boss_2.Name, heal))
		centerText(fmt.Sprintf("%s : %d/%d PV", boss_2.Name, boss_2.HP, boss_2.MaxHP))
		return
	} else if random < 70 {
		damage := 100
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		centerText(fmt.Sprintf("%s utilise gifle stélaire ! Vous mais preque KO et vous ridiculise devant toute l'arrène et subissez %d dégâts !", boss_2.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	} else if random < 80 {
		damage := 70
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}

		centerText(fmt.Sprintf("%s vous fait Un Baiser Assassin sur la joue… vous inflige des dégâts psychologiques irréversible ! et subissez %d dégâts !", boss_2.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	} else if random < 85 {
		damage := 60
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}

		centerText(fmt.Sprintf("%s vous donne un Coup de Talon Cosmique → frappe élégante mais complètement absurde. ! et subissez %d dégâts !", boss_2.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	} else if random < 90 {
		damage := 60
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}

		centerText(fmt.Sprintf("%s utilise Serrage de Main Mortel → un simple shake qui vous casse le bras! et inflige %d dégâts !", boss_2.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	} else if random < 95 {
		damage := 60
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}

		centerText(fmt.Sprintf("%s Pose de Selfie Explosive → elle pose, vous etes traumatisé par la scène ! et inflige %d dégâts !", boss_2.Name, damage))
		centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
		return
	}
	damage := 10
	c.HP -= damage
	if c.HP < 0 {
		c.HP = 0
	}
	centerText(fmt.Sprintf("%s vous frappe et inflige %d dégâts !", boss_2.Name, damage))
	centerText(fmt.Sprintf("%s : %d/%d PV", c.Name, c.HP, c.MaxHP))
}
