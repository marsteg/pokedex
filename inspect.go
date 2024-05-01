package main

import (
	"fmt"
	"strconv"
)

func inspect(c *config, input []string) error {
	key := input[1]

	p, exists := c.pokedex[key]
	if !exists {
		fmt.Println("have not caught " + key)
		return nil
	}
	fmt.Println("Inspecting " + key + "...")
	fmt.Println("Name: " + p.Name)
	fmt.Println("Height: " + strconv.Itoa(p.Height))
	fmt.Println("Weight: " + strconv.Itoa(p.Weight))
	fmt.Println("Stats: ")
	for _, s := range p.Stats {
		fmt.Println(" - " + s.Stat.Name + ": " + strconv.Itoa(s.BaseStat))
	}
	fmt.Println("Types: ")
	for _, s := range p.Types {
		fmt.Println(" - " + s.Type.Name)
	}
	return nil
}
