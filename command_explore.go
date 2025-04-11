package main

import (
	"errors"
	"fmt"
)


func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a location name")
	}
	locName := args[0]
	locationRes, err := cfg.pokeapiClient.SingleLocation(locName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locName)
	fmt.Println("Pokemon found:")
	for _, location := range locationRes.PokemonEncounters {
		fmt.Println(location.Pokemon.Name)

	}

	return nil
}