package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/marsteg/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	next     string
	previous string
	cache    pokecache.Cache
	pokedex  map[string]Pokemon
}

func main() {
	var c config
	c.loadConfig()

	fmt.Println("Pokedex starting!")
	scanner := bufio.NewScanner(os.Stdin)
	var last string
	fmt.Print(last)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		// arrow up should bring back the most recent command

		//input := strings.Fields(scanner.Text())
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		last = strings.Join(input, " ")

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&c, input)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}

}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.TrimSpace(words[i])
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func (c *config) loadConfig() {
	//simple loading empty config as we do not have persistency
	c.cache = pokecache.NewCache(5 * time.Minute)

	c.pokedex = make(map[string]Pokemon)

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world. Each subsequent call to map will display the next 20 locations, and so on.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go Backwards on the Map! Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore an Area and find the pokemons there.",
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon and add it to your Pokedex!",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon with your Pokedex!",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Check which Pokemon you caught already and added to your Pokedex.",
			callback:    pokedex,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    help,
		},
		"lil": {
			name:        "lil",
			description: "it's a tiny command",
			callback:    tiny,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exit,
		},
	}
}
