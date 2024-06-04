package main

import (
	"fmt"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/fatih/color"
)

var cyan func(a ...interface{}) string = color.New(color.FgCyan, color.Bold).SprintFunc()
var redErr func(a ...interface{}) string = color.New(color.FgRed).SprintFunc()

func startRepl() {
	input := ""
	fmt.Printf("\n%s > ", cyan("Pokedex"))
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.CtrlC, keys.Escape:
			return true, nil
		case keys.Enter:
			fmt.Println()
			if tokens := cleanInput(input); tokens != nil {
				processCommand(*tokens...)
			}
			input = ""
			fmt.Printf("\n%s > ", cyan("Pokedex"))
		case keys.Backspace:
			if len(input) >= 1 {
				input = input[:len(input)-1]
				fmt.Printf("\033[2K\r%s > %s", cyan("Pokedex"), input)
			}
		default:
			input += string(key.Runes)
			fmt.Print(string(key.Runes))
		}
		return false, nil
	})
}

func processCommand(inputs ...string) {
	command, exists := getCommands()[inputs[0]]
	if !exists {
		fmt.Println(redErr("Unknown command"))
		return
	}
	if err := command.callback(inputs[1:]...); err != nil {
		fmt.Println(redErr(err))
	}
}

func cleanInput(text string) *[]string {
	output := strings.TrimSpace(strings.ToLower(text))
	words := strings.Fields(output)
	if len(words) == 0 {
		return nil
	}
	return &words
}

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Shows the names of the next 20 location in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the names of the previous 20 location in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows the names of Pokemon in the given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the specified Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows details about a previously caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows a list of all the caught Pokemons",
			callback:    commandPokedex,
		},
	}
}
