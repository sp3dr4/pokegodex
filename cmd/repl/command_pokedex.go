package main

import (
	"errors"
	"fmt"
)

func commandPokedex(args ...string) error {
	if len(args) > 0 {
		return errors.New("command does not accept arguments")
	}
	if len(pokedex) == 0 {
		fmt.Println("your pokedex is empty!")
	}
	for _, p := range pokedex {
		fmt.Printf(" - %s\n", p.name)
	}
	return nil
}
