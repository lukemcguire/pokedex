package pokeapi

import (
	"encoding/json"
	"fmt"
)

// GetPokemon fetches information about a pokemon
func (c *Client) GetPokemon(name *string) (Pokemon, error) {
	if name == nil || *name == "" {
		return Pokemon{}, fmt.Errorf("pokemon name or id not provided")
	}

	endpoint := baseURL + "/pokemon/" + *name

	data, err := c.get(endpoint)
	if err != nil {
		return Pokemon{}, err
	}

	var res Pokemon
	err = json.Unmarshal(data, &res)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshaling pokemon data: %w", err)
	}

	return res, nil

}
