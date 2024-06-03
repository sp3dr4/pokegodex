package main

import (
	"errors"
	"fmt"
)

func commandInspect(args ...string) error {
	if len(args) != 1 {
		return errors.New("command expects pokemon name argument")
	}
	name := args[0]
	pokemon, caught := pokedex[name]
	if !caught {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.name)
	fmt.Printf("Height: %d\n", pokemon.height)
	fmt.Printf("Weight: %d\n", pokemon.weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.stats {
		fmt.Printf("  - %s: %d\n", s.name, s.value)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.types {
		fmt.Printf("  - %s\n", t)
	}

	return nil
}
