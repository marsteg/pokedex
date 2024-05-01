package main

import (
	"fmt"
	"os"
)

func help(c *config, _ []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for com := range commands {
		fmt.Printf(commands[com].name + ": " + commands[com].description + "\n")

	}
	return nil
}

func exit(c *config, _ []string) error {
	os.Exit(0)
	return nil
}

func tiny(c *config, input []string) error {
	fmt.Println("Tiny is a little lil, but not a Poke-Quil!")
	fmt.Printf("Command: %v, argument: %v\n", input[0], input[1])
	return nil
}
