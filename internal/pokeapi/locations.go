package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(url *string) (RespLocations, error) {
	fullURL := baseURL + "/location-area"
	if url != nil {
		fullURL = *url
	}

	if val, ok := c.cache.Get(fullURL); ok {
		locationRes := RespLocations{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return RespLocations{}, err
		}

		return locationRes, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, err
	}

	locationsRes := RespLocations{}
	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(fullURL, data)
	return locationsRes, nil
}
