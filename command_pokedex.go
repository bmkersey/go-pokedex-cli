package main

import "fmt"


func commandPokedex(cfg *config, args ...string) error {
	currentPokemon := cfg.pokemonCaught
	if len(currentPokemon) < 1 {
		fmt.Println("You have not caught any pokemon yet...")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range currentPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}