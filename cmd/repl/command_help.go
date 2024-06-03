package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func commandHelp(args ...string) error {
	if len(args) > 0 {
		return errors.New("command does not accept arguments")
	}
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cyan(cmd.name), cmd.description)
	}
	fmt.Println()
	return nil
}
