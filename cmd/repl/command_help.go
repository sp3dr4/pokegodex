package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

var yellow func(a ...interface{}) string = color.New(color.FgYellow, color.Bold).SprintFunc()

func commandHelp(args ...string) error {
	if len(args) > 0 {
		return errors.New("command does not accept arguments")
	}
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println()
	fmt.Printf("Welcome to the %s\n", yellow("Pokedex!"))
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cyan(cmd.name), cmd.description)
	}
	fmt.Println()
	return nil
}
