package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}


func startRepl() {
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
			err := command.callback()
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}