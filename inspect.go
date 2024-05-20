package main

import (
	"fmt"
	"strconv"
	"strings"
)

func inspect(c *config, input []string) error {
	key := input[1]
	var subcommand string
	if len(input) > 2 {
		subcommand = input[2]
	}

	p, exists := c.pokedex[key]
	if !exists {
		fmt.Println("have not caught " + key)
		return nil
	}
	fmt.Println("Inspecting " + key + "...")
	fmt.Println("Name: " + p.Name)
	fmt.Println("Height: " + strconv.Itoa(p.Height))
	fmt.Println("Weight: " + strconv.Itoa(p.Weight))
	fmt.Println("Base XP: " + strconv.Itoa(p.BaseExperience))
	fmt.Println("Stats: ")
	for _, s := range p.Stats {
		fmt.Println(" - " + s.Stat.Name + ": " + strconv.Itoa(s.BaseStat))
	}
	fmt.Println("Types: ")
	for _, s := range p.Types {
		fmt.Println(" - " + strings.ToUpper(s.Type.Name))
	}

	switch subcommand {
	case "":
		return nil
	case "moves":
		return inspectMoves(c, p)
	case "strongest":
		return inspectStrongest(c, p)

	default:
		fmt.Println("Invalid subcommand")
	}
	return nil
}

func inspectMoves(c *config, p Pokemon) error {
	fmt.Println("Moves: ")
	for _, s := range p.Moves {
		reply, err := getPokething[Move](s.Move.URL, c.cache)
		if err != nil {
			fmt.Println(err)
		}
		printMove(reply)
	}
	return nil
}

func inspectStrongest(c *config, p Pokemon) error {
	fmt.Println("Strongest Moves: ")
	moves := []Move{}
	for _, s := range p.Moves {
		reply, err := getPokething[Move](s.Move.URL, c.cache)
		if err != nil {
			fmt.Println(err)
		}
		if reply.Power > 110 && reply.Accuracy == 100 {
			moves = append(moves, reply)
		}
	}
	for _, m := range moves {
		printMove(m)
	}
	return nil
}

func printMove(m Move) {
	fmt.Println(" - " + m.Name)
	fmt.Println("   - Accuracy: " + fmt.Sprintf("%d", m.Accuracy))
	fmt.Println("   - Power: " + fmt.Sprintf("%d", m.Power))
	fmt.Println("   - CritRate: " + fmt.Sprintf("%d", m.Meta.CritRate))
	fmt.Println("   - Attack Type: " + strings.ToUpper(m.Type.Name))
}
