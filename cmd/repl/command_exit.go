package main

import (
	"errors"
	"os"
)

func commandExit(args ...string) error {
	if len(args) > 0 {
		return errors.New("command does not accept arguments")
	}
	os.Exit(0)
	return nil
}
