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
