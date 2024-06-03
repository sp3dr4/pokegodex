package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/sp3dr4/pokegodex/internal/pokeapi"
)

const maxExp int = 1000
const decreaseScalingFactor float64 = 1 // [0, 1]
const catchThreshold float64 = .5

type Pokemon struct {
	baseExperience int
	height         int
	weight         int
	id             int
	name           string
	types          []string
	stats          []struct {
		name  string
		value int
	}
}

func (p *Pokemon) doesBallCatch() bool {
	normalizedExp := float64(p.baseExperience) / float64(maxExp)
	chance := 1 - decreaseScalingFactor*normalizedExp
	if chance < 0 {
		chance = 0
	} else if chance > 1 {
		chance = 1
	}
	final := chance * rand.Float64()
	return final >= catchThreshold
}

func fromResponse(resp *pokeapi.PokemonDetailsResponse) *Pokemon {
	types := make([]string, len(resp.Types))
	for i, rt := range resp.Types {
		types[i] = rt.Type.Name
	}
	stats := make([]struct {
		name  string
		value int
	}, len(resp.Stats))
	for i, rs := range resp.Stats {
		stats[i] = struct {
			name  string
			value int
		}{name: rs.Stat.Name, value: rs.BaseStat}
	}
	return &Pokemon{
		baseExperience: resp.BaseExperience,
		height:         resp.Height,
		weight:         resp.Weight,
		id:             resp.ID,
		name:           resp.Name,
		types:          types,
		stats:          stats,
	}
}

var seedVal string = os.Getenv("POKEGODEX_SEED")

var pokedex map[string]*Pokemon = map[string]*Pokemon{}

func commandCatch(args ...string) error {
	if seedVal != "" {
		seedInt, err := strconv.ParseInt(seedVal, 10, 64)
		if err == nil {
			fmt.Printf("seeding with %d\n", seedInt)
			rand.New(rand.NewSource(seedInt))
		}
	}

	if len(args) != 1 {
		return errors.New("command expects pokemon name argument")
	}
	name := args[0]

	res, err := pokeapi.GetPokemon(name)
	if err != nil {
		return err
	}
	if _, alreadyCaught := pokedex[name]; alreadyCaught {
		return errors.New("pokemon already caught")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon := fromResponse(res)
	if pokemon.doesBallCatch() {
		pokedex[name] = pokemon
		color.Green("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
