package main

import (
	"fmt"
)

func commandMap(cfg *Config, _ ...string) error {
	res, err := cfg.PokeClient.ListLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	// update the next and previous pointers in the config
	cfg.Next = res.Next
	cfg.Prev = res.Previous

	// Print the location areas
	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapB(cfg *Config, _ ...string) error {
	cfg.Next = cfg.Prev
	return commandMap(cfg)
}
