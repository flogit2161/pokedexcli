package main

import (
	"fmt"
	"os"

	"github.com/flogit2161/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
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
		"map": {
			name:        "map",
			description: "Display next page 20 location area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous page 20 location area",
			callback:    commandMapBack,
		},
	}
}

// EXIT FUNCTION
func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// HELP FUNCTION
func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

// MAP 20 LOCATIONS FUNCTION
func commandMap(cfg *config) error {
	locResponse, err := cfg.pokeapiClient.ClientRequest(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = locResponse.Next
	cfg.previous = locResponse.Previous

	for _, loc := range locResponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

// MAP 20 LOCATIONS PREVIOUS PAGE FUNCTION
func commandMapBack(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	locResponse, err := cfg.pokeapiClient.ClientRequest(cfg.previous)
	if err != nil {
		return err
	}

	cfg.previous = locResponse.Previous
	cfg.next = locResponse.Next

	for _, loc := range locResponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
