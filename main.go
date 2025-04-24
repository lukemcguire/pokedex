// Package main implements a simple Pokedex command-line interface
// with various commands for interacting with Pokemon data.
package main

import (
	"time"

	"github.com/lukemcguire/pokedexcli/internal/pokeapi"
)

// main entry point for the Pokedex CLI
func main() {
	pokeClient := pokeapi.NewClient(pokeapi.WithHTTPTimeout(5 * time.Second))
	cfg := &Config{
		PokeClient:    pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
