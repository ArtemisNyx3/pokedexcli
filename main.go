package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliDirectory map[string]cliCommand

func main() {
	cliDirectory = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		key := scanner.Text()
		userInput := cleanInput(key)
		// fmt.Printf("Your command was: %s\n", userInput[0])
		command, err := cliDirectory[userInput[0]]
		if err == false {
			fmt.Println("Invalid Command")
		} else {

			command.callback()
		}

	}

}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	temp := strings.Split(text, " ")

	var cleanIP []string
	for _, str := range temp {
		if str == " " {
			continue
		} else if len(str) == 0 {
			continue
		} else {
			cleanIP = append(cleanIP, str)
		}
	}
	return cleanIP
}
