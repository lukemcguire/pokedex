package main

import "fmt"

// commandExplore lists all pokemon available in the specified area

func commandExplore(cfg *Config, areas ...string) error {
	if len(areas) < 1 {
		return fmt.Errorf("error: must provide at least one area to explore")
	}

	for _, area := range areas {
		fmt.Printf("Exploring %s...\n", area)
		err := listAreaPokemon(cfg, area)
		if err != nil {
			fmt.Printf("Error fetching %s area details: %s\n", area, err)
		}
	}
	return nil
}

func listAreaPokemon(cfg *Config, area string) error {
	locationRes, err := cfg.PokeClient.ListLocation(&area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationRes.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil

}
