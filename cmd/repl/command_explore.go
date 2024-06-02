package main

import (
	"errors"
	"fmt"

	"github.com/sp3dr4/pokegodex/internal/pokeapi"
)

func commandExplore(args ...string) error {
	if len(args) != 1 {
		return errors.New("command expects location name argument")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	res, err := pokeapi.ListLocationPokemons(args[0])
	if err != nil {
		return err
	}
	if len(res.PokemonEncounters) == 0 {
		return errors.New("no pokemons found in location")
	}
	fmt.Println("Found Pokemon:")
	for _, p := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}
