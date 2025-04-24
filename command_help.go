package main

import "fmt"

// commandHelp displays all available commands and their descriptions.
func commandHelp(_ *Config, _ ...string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, command := range supportedCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
