package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client) GetSinglePokemon(name string) (Pokemon, error){
	fullURL := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(fullURL); ok {
		pokemonRes := Pokemon{}
		err := json.Unmarshal(val, &pokemonRes)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonRes, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonRes := Pokemon{}

	err = json.Unmarshal(data, &pokemonRes)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)
	
	return pokemonRes, nil
}