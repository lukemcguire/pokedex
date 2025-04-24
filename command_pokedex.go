package main

import "fmt"

// commandHelp displays all available commands and their descriptions.
func commandPokedex(cfg *Config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
