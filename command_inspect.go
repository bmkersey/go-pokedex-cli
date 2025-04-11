package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please provide a pokemon to inspect")
	}

	pokemonName := args[0]

	if info, ok := cfg.pokemonCaught[pokemonName]; !ok {
		fmt.Printf("You have not caught %s yet.\n", pokemonName)
	}else {
		fmt.Printf("Name: %s\n", info.Name)
		fmt.Printf("Height: %d\n", info.Height)
		fmt.Printf("Weight: %d\n", info.Weight)
		fmt.Println("Stats:")
		for _, stat := range info.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range info.Types {
			fmt.Println(   "-", typeInfo.Type.Name)
		}
	}
	
	return nil
}