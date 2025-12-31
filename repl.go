package main 

import (
	"fmt"
	"os"
	"bufio"
)

func repl(){
	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if userInput.Scan() == false{
			fmt.Println("Error Scannning User Input, exiting")
			return
		}
		userInputString := cleanInput(userInput.Text())
		if len(userInputString) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", userInputString[0])

	}
}