package location

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sp3dr4/pokegodex/internal/pokecache"
)

var cache pokecache.Cache = *pokecache.NewCache(2 * time.Minute)

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
	url := "https://pokeapi.co/api/v2/location"
	if pageUrl != nil {
		url = *pageUrl
	}

	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("err fetching list of locations: %v", err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
		}

		if err != nil {
			return nil, fmt.Errorf("err reading response: %v", err)
		}
		cache.Add(url, body)
	}

	locations := LocationsResponse{}
	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, fmt.Errorf("err decoding list of locations: %v", err)
	}
	return &locations, nil
}
