package pokeapi

import (
	"encoding/json"
	"fmt"
)

type PokemonsResponse struct {
	ID       int `json:"id"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func ListLocationPokemons(name string) (*PokemonsResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + name

	body, err := Get(url)
	if err != nil {
		return nil, err
	}

	pokemons := PokemonsResponse{}
	if err := json.Unmarshal(body, &pokemons); err != nil {
		return nil, fmt.Errorf("err decoding list of pokemons: %v", err)
	}
	return &pokemons, nil
}
