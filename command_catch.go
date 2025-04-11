package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must pass a pokemon name")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a pokeball at %s.\n", pokemonName)

	pokemonRes, err := cfg.pokeapiClient.GetSinglePokemon(pokemonName)
	if err != nil {
		return err
	}

	baseXP := pokemonRes.BaseExperience

	catchResult := rand.Intn(baseXP)

	if catchResult > 40 {
		fmt.Printf("%s escaped!!!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!!!\n", pokemonName)

	cfg.pokemonCaught[pokemonRes.Name] = pokemonRes
	return nil
}