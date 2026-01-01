package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/flogit2161/pokedexcli/internal/pokeapi"
)

func repl() {
	apiClient := pokeapi.NewClient(5 * time.Second)

	cfg := &config{
		pokeapiClient: apiClient,
	}

	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		userInput.Scan()

		cleanedInput := cleanInput(userInput.Text())
		if len(cleanedInput) == 0 {
			fmt.Println("Please enter a valid command ('help' will display all cmds)")
			continue
		}

		//Only the first word of the slice is checked
		command, ok := getCommands()[cleanedInput[0]]
		if ok == false {
			fmt.Println("Command not found, please type a valid command (example : 'help')")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			return
		}

	}
}

// CLEANS USER INPUT FUNCTION
func cleanInput(text string) []string {
	strippedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(strippedText)
	return strings.Fields(loweredText)
}
