package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sp3dr4/pokegodex/internal/cache"
)

var c cache.Cache = *cache.NewCache(2 * time.Minute)

func get(url string) ([]byte, error) {
	body, ok := c.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("err sending http request: %v", err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
		}

		if err != nil {
			return nil, fmt.Errorf("err reading response: %v", err)
		}
		c.Add(url, body)
	}
	return body, nil
}
