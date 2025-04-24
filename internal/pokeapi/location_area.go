package pokeapi

import (
	"encoding/json"
	"fmt"
)

// ListLocationAreas fetches location areas, using a URL if provided or the default endpoint
func (c *Client) ListLocationAreas(url *string) (LocationAreasRes, error) {
	endpoint := baseURL + "/location-area"
	if url != nil && *url != "" {
		endpoint = *url
	}

	data, err := c.get(endpoint)
	if err != nil {
		return LocationAreasRes{}, err
	}

	var res LocationAreasRes
	err = json.Unmarshal(data, &res)
	if err != nil {
		return LocationAreasRes{}, fmt.Errorf("error unmarshaling location areas: %w", err)
	}

	return res, nil
}

// ListLocation fetches all pokemon available at a given area
func (c *Client) ListLocation(area *string) (LocationRes, error) {
	if area == nil || *area == "" {
		return LocationRes{}, fmt.Errorf("area not provided for search")
	}

	endpoint := baseURL + "/location-area/" + *area

	data, err := c.get(endpoint)
	if err != nil {
		return LocationRes{}, err
	}

	var res LocationRes
	err = json.Unmarshal(data, &res)
	if err != nil {
		return LocationRes{}, fmt.Errorf("error unmarshaling location areas: %w", err)
	}

	return res, nil

}
