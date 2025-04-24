package main

import "github.com/lukemcguire/pokedexcli/internal/pokeapi"

// cliCommand represents a command available in the Pokedex CLI.
// Each command has a name, description, and a callback function.
type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

// Config holds information about next and previous urls for commands that
// require them.
type Config struct {
	Next          *string
	Prev          *string
	PokeClient    *pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
}
