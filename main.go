package main

import (
	"time"

	"github.com/bmkersey/go-pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokemonCaught: map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
