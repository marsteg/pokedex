package main

import (
	"fmt"
)

func pokedex(c *config, input []string) error {
	fmt.Println("Your Pokedex:")
	if len(c.pokedex) == 0 {
		fmt.Println("No Pokemon caught, yet... go and catch 'em all!")
	}
	for pokename := range c.pokedex {
		fmt.Printf("- %s\n", pokename)
	}
	return nil
}
