package main

import (
	"fmt"
	"os"
)

// commandExit handles the exit command by printing a goodbye message
// and exiting the program.
func commandExit(_ *Config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
