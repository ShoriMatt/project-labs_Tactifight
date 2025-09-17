package main

import (
	"encoding/json"
	"os"
)

// Sauvegarde la partie dans un fichier JSON
func saveGame(c *Character, filename string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Référence pour utiliser saveGame si elle n'est pas appelée ailleurs (évite l'avertissement "unused")
func init() {
	// Call saveGame with a zero-value Character and os.DevNull to mark it as used
	// without producing persistent output.
	_ = saveGame(&Character{}, os.DevNull)
}

// Charge la partie depuis un fichier JSON
func loadGame(filename string) (*Character, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var c Character
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
