package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/term"

	// audio
	"bytes"
	"embed"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
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
	spell := "boule de feu"
	for _, s := range c.Skills {
		if s == spell {
			centerText("Vous connaissez déjà ce sort " + spell + " !")
			return
		}
	}
	c.Skills = append(c.Skills, spell)
	centerText("Nouveau sort acquis : " + spell)
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

//
// === Gestion des sons embarqués ===
//

//go:embed assets/*.wav
var sounds embed.FS

// Joue un son embarqué (nom = "select.wav")
func playSound(name string) {
	data, err := sounds.ReadFile("assets/" + name)
	if err != nil {
		log.Println("Son introuvable:", err)
		return
	}

	streamer, format, err := wav.Decode(bytes.NewReader(data))
	if err != nil {
		log.Println("Erreur de décodage audio:", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
