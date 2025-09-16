package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bienvenue sur TACTIFIGHT !")

	c1 := characterCreation(reader)

	fmt.Printf("Personnage créé : %s (%s) - PV %d/%d - %d potions\n",
		c1.Name, c1.Class, c1.HP, c1.MaxHP, len(c1.Inventory))

	mainMenu(&c1, reader)
}
