package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     *string
	previous *string
}

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
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
			description: "Display 20 location area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous page 20 location area",
			callback:    commandMapBack,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap(cfg *config) error {
	if cfg.next == nil {
		return fmt.Errorf("Error setting the URL")
	}

	response, err := http.Get(*cfg.next)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	locations := LocationArea{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}

	cfg.next = locations.Next
	cfg.previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	response, err := http.Get(*cfg.previous)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	locations := LocationArea{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}

	cfg.previous = locations.Previous
	cfg.next = locations.Next

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
