package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bmkersey/go-pokedex-cli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}


func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		word := cleaned[0]
		if command, exists := getCommands()[word]; exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println("Error:", err)
			} 
		}else {
			fmt.Println("Command not recognized")
		}
	}
}

func cleanInput(text string) []string {
	var cleaned []string
	words := strings.Fields(strings.ToLower(text))
	cleaned = append(cleaned, words...)
	return cleaned
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Retrieves next page of locations",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Retrieves the previous page of locations",
			callback: commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}