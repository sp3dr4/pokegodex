package main

import (
	"errors"
	"fmt"

	"github.com/sp3dr4/pokegodex/internal/pokeapi"
)

var next *string
var prev *string

func commandMap(args ...string) error {
	if len(args) > 0 {
		return errors.New("command does not accept arguments")
	}
	res, err := pokeapi.ListLocations(next)
	if err != nil {
		return err
	}
	if res.Next != "" {
		next = &res.Next
	} else {
		next = nil
	}
	if res.Previous != "" {
		prev = &res.Previous
	} else {
		prev = nil
	}
	for _, i := range res.Results {
		fmt.Printf("%s\n", i.Name)
	}
	return nil
}

func commandMapb(args ...string) error {
	if len(args) > 0 {
		return errors.New("command does not accept arguments")
	}
	if prev == nil {
		return errors.New("no locations history found")
	}
	res, err := pokeapi.ListLocations(prev)
	if err != nil {
		return err
	}
	if res.Next != "" {
		next = &res.Next
	} else {
		next = nil
	}
	if res.Previous != "" {
		prev = &res.Previous
	} else {
		prev = nil
	}
	for _, i := range res.Results {
		fmt.Printf("%s\n", i.Name)
	}
	return nil
}
