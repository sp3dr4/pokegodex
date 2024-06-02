package main

import (
	"errors"
	"fmt"
)

func commandHelp(args ...string) error {
	if len(args) > 0 {
		return errors.New("command does not accept arguments")
	}
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
