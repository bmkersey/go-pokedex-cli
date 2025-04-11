package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) SingleLocation(location string) (Location, error) {
	fullURL := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(fullURL); ok {
		locactionRes := Location{}
		err := json.Unmarshal(val, &locactionRes)
		if err != nil {
			return Location{}, err
		}

		return locactionRes, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Location{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationRes := Location{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(fullURL, data)
	return locationRes, nil
}