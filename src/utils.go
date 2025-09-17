package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/term"
)

func IsAlpha(s string) bool {
	for _, r := range s {
		if !(unicode.IsLetter(r) || r == ' ' || r == '-' || r == '\'') {
			return false
		}
	}
	return true
}

func formatNom(nom string) string {
	if len(nom) == 0 {
		return ""
	}
	mots := strings.Fields(strings.ToLower(nom))
	for i, m := range mots {
		r := []rune(m)
		if len(r) > 0 {
			r[0] = unicode.ToUpper(r[0])
		}
		mots[i] = string(r)
	}
	return strings.Join(mots, " ")
}

func spellBook(c *Character) {
	spell := "balle perçante"
	for _, s := range c.Skills {
		if s == spell {
			centerText("Vous connaissez déjà ces munition " + spell + " !")
			return
		}
	}
	c.Skills = append(c.Skills, spell)
	centerText("Nouvelle munition acquise : " + spell)
}

// Affiche du texte centré dans le terminal
func centerText(text string) {
	width := 80
	if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
		width = w
	}
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		l := len([]rune(line)) // rune-safe pour accents
		if l >= width {
			fmt.Println(line)
			continue
		}
		padding := (width - l) / 2
		fmt.Printf("%s%s\n", strings.Repeat(" ", padding), line)
	}
}
