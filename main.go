package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ArtemisNyx3/pokedexcli/internal/pokeapi"
)

var cliDirectory map[string]cliCommand

type configuration struct {
	next     string
	previous string
}

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
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Displays the names of the previous 20 locations",
			callback:    commandMapBack,
		},
	}

	config := configuration{
		next:     "",
		previous: "",
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		key := scanner.Text()
		userInput := cleanInput(key)
		// fmt.Printf("Your command was: %s\n", userInput[0])
		command, err := cliDirectory[userInput[0]]
		if !err {
			fmt.Println("Invalid Command")
		} else {

			command.callback(&config)
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

type cliCommand struct {
	name        string
	description string
	callback    func(c *configuration) error
}

func commandExit(c *configuration) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *configuration) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for name, cmd := range cliDirectory {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}

func commandMap(c *configuration) error {
	locations,err := pokeapi.GetLocations(c.next)
	if err != nil {
		fmt.Println("Got error --- ",err)
		return err
	}
	for _,result := range locations.Results{
		fmt.Println(result.Name)
	}
	c.next = locations.Next
	c.previous = locations.Previous
	return nil
}

func commandMapBack(c *configuration) error {
	locations,err := pokeapi.GetLocations(c.previous)
	if err != nil {
		fmt.Println("Got error --- ",err)
		return err
	}
	for _,result := range locations.Results{
		fmt.Println(result.Name)
	}
	c.next = locations.Next
	c.previous = locations.Previous
	return nil
}