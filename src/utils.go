package main

import (
	"fmt"
	"strings"
	"unicode"
)

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
