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
		command, ok := getCommands()[userInput.Text()]
		if ok == false {
			fmt.Println("Command not found, please type a valid command (example : 'help')")
			continue
		}
		command.callback()
		
	}
}