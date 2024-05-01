package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/marsteg/pokedex/internal/pokecache"
)

type areaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func explore(c *config, input []string) error {
	area := input[1]
	var url string
	fmt.Printf("Exploring " + area + "...\n")
	url = "https://pokeapi.co/api/v2/location-area/" + area + "/"
	reply, err := getArea(url, c.cache)
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

func getArea(url string, c pokecache.Cache) (areaResponse, error) {
	var reply areaResponse

	//If you already have the data for a given URL (which is our cache key)
	//in the cache, you should use that instead of making a new request.
	//Whenever you do make a request, you should add the response to the cache.
	val, exists := c.Get(url)
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return reply, err
		}
		res.Body.Close()
		c.Add(url, body)
		err = json.Unmarshal(body, &reply)
		if err != nil {
			return reply, err
		}
		return reply, nil
	} else {
		err := json.Unmarshal(val, &reply)
		if err != nil {
			return reply, err
		}
		return reply, nil
	}
}
