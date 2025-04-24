package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/lukemcguire/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("must provide a name or id of pokemon to catch")
	}

	name := args[0]

	pokemon, err := cfg.PokeClient.GetPokemon(&name)
	if err != nil {
		return err
	}

	if caught := throwPokeball(pokemon); caught {
		cfg.caughtPokemon[name] = pokemon
	}
	return nil

}

func throwPokeball(pokemon pokeapi.Pokemon) bool {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	var catchRate float64
	switch {
	case pokemon.BaseExperience < 100:
		catchRate = 0.8
	case pokemon.BaseExperience < 200:
		catchRate = 0.5
	default:
		catchRate = 0.2
	}
	if rand.Float64() < catchRate {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return true
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return false
	}
}
