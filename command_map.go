package main

import (
	"errors"
	"fmt"

	"github.com/sp3dr4/pokegodex/internal/location"
)

var next *string
var prev *string

func commandMap() error {
	res, err := location.ListLocations(next)
	if err != nil {
		return err
	}
	if res.Next != "" {
		next = &res.Next
	}
	if res.Previous != "" {
		prev = &res.Previous
	}
	for _, i := range res.Results {
		fmt.Printf("%s\n", i.Name)
	}
	return nil
}

func commandMapb() error {
	if prev == nil {
		return errors.New("no locations history found")
	}
	res, err := location.ListLocations(prev)
	if err != nil {
		return err
	}
	if res.Next != "" {
		next = &res.Next
	}
	if res.Previous != "" {
		prev = &res.Previous
	}
	for _, i := range res.Results {
		fmt.Printf("%s\n", i.Name)
	}
	return nil
}
