package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func catch(c *config, input []string) error {
	pokemon := input[1]

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon + "/"
	reply, err := getPokething[Pokemon](url, c.cache)
	if err != nil {
		if err.Error() == "invalid character 'N' looking for beginning of value" {
			fmt.Printf("failed throwing a Pokeball at " + pokemon + "!!\n")
			return errors.New("the Pokemon vanished instantly... or Pokemon does not exist on PokeAPI.co :'-( ")
		}
		fmt.Println(err)
	}
	fmt.Printf("Throwing a Pokeball at " + pokemon + "...\n")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chance := r.Intn(380)
	schance := fmt.Sprintf("%v", chance)
	fmt.Printf("diced chance: " + schance + "\n")
	if chance > reply.BaseExperience {
		fmt.Printf(pokemon + " was caught!\n")
		// Add pokemon to pokedex
		c.pokedex[pokemon] = reply
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf(pokemon + " escaped!\n")
	}

	return nil
}

//Once the Pokemon is caught, add it to the user's Pokedex. I used a map[string]Pokemon
// to keep track of caught Pokemon.
//You'll want to store the Pokemon's data so that in the next step we can use it. (caching)
