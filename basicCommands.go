package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/marsteg/pokedex/internal/pokecache"
)

type Pokething interface {
	Pokemon
	LocationArea
	LocationAreas
}

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

func getPokething[T Pokemon | LocationArea | LocationAreas | Move](url string, c pokecache.Cache) (T, error) {
	var pokemon T

	val, exists := c.Get(url)
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return pokemon, err
		}
		res.Body.Close()
		c.Add(url, body)
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			return pokemon, err
		}
		return pokemon, nil
	} else {
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}
}
