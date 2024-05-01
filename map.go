package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/marsteg/pokedex/internal/pokecache"
)

type response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config, _ []string) error {
	if c.next == "" {
		reply, err := getreply("https://pokeapi.co/api/v2/location-area/", c.cache)
		if err != nil {
			fmt.Println(err)
		}
		c.next = reply.Next
		c.previous = reply.Previous
		for _, r := range reply.Results {
			fmt.Println(r.Name)
		}
		return nil
	} else {
		reply, err := getreply(c.next, c.cache)
		if err != nil {
			fmt.Println(err)
		}
		c.next = reply.Next
		c.previous = reply.Previous
		for _, r := range reply.Results {
			fmt.Println(r.Name)
		}

		return nil
	}
}

func commandMapB(c *config, _ []string) error {
	if c.previous == "" {
		fmt.Println("error, go forward on the map first")
		return nil
	} else {
		reply, err := getreply(c.previous, c.cache)
		if err != nil {
			fmt.Println(err)
		}
		c.next = reply.Next
		c.previous = reply.Previous
		for _, r := range reply.Results {
			fmt.Println(r.Name)
		}

		return nil
	}
}

func getreply(url string, c pokecache.Cache) (response, error) {
	var reply response

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
