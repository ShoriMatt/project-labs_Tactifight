package main

import (
	"fmt"
)

type Character struct {
	Nom                string
	Classe             string
	Niveau             int
	PointsDeVieMax     int
	PointsDeVieActuels int
	Inventaire         []string
}

func main() {
	perso := Character{
		Nom:                "Arthas",
		Classe:             "Soldat spatial",
		Niveau:             10,
		PointsDeVieMax:     120,
		PointsDeVieActuels: 120,
		Inventaire:         []string{"Épée", "Bouclier", "Potion de soin"},
	}

	fmt.Println(perso)
}
