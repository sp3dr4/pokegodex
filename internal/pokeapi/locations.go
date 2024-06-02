package pokeapi

import (
	"encoding/json"
	"fmt"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationsResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func ListLocations(pageUrl *string) (*LocationsResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	body, err := Get(url)
	if err != nil {
		return nil, err
	}

	locations := LocationsResponse{}
	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, fmt.Errorf("err decoding list of locations: %v", err)
	}
	return &locations, nil
}
