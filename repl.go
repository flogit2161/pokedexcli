package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {
	url := "https://pokeapi.co/api/v2/location-area"

	cfg := config{
		next:     &url,
		previous: nil,
	}

	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if userInput.Scan() == false {
			fmt.Println("Error Scanning User Input, exiting")
			return
		}
		command, ok := getCommands()[userInput.Text()]
		if ok == false {
			fmt.Println("Command not found, please type a valid command (example : 'help')")
			continue
		}
		command.callback(&cfg)

	}
}
