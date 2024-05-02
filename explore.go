package main

import (
	"errors"
	"fmt"
)

func explore(c *config, input []string) error {
	area := input[1]
	var url string
	fmt.Printf("Exploring " + area + "...\n")
	url = "https://pokeapi.co/api/v2/location-area/" + area + "/"
	reply, err := getPokething[LocationArea](url, c.cache)
	if err != nil {
		if err.Error() == "invalid character 'N' looking for beginning of value" {
			return errors.New("no Pokemon found... or explored Area does not exist on PokeAPI.co :'-( ")
		}
		fmt.Println(err)
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range reply.PokemonEncounters {
		s := fmt.Sprintf(" - %s", encounter.Pokemon.Name)
		fmt.Println(s)
	}
	return nil
}
