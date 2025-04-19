package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ArtemisNyx3/pokedexcli/internal/pokeapi"
	"github.com/ArtemisNyx3/pokedexcli/internal/pokecache"
)

var cliDirectory =  make(map[string]cliCommand )

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
		"explore": {
			name:        "explore",
			description: "explore <area-name> : Displays the name of the pokemon in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch <pokemon-name> : Throw Pokeball at pokemon",
			// TODO: Check pokemon in current area
			callback: commandCatch,
		},
	}

	config := configuration{
		next:     "",
		previous: "",
	}

	cache := pokecache.NewCache(10 * time.Second)

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
			if len(userInput) == 1 {
				command.callback(&config, cache, "")
			} else {
				command.callback(&config, cache, userInput[1])
			}
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
	callback    func(config *configuration, cache *pokecache.Cache, arg string) error
}

func commandExit(c *configuration, cache *pokecache.Cache , arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *configuration, cache *pokecache.Cache,  arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for name, cmd := range cliDirectory {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}

func commandMap(config *configuration, cache *pokecache.Cache,  arg string) error {
	locations, err := pokeapi.GetLocations(config.next, cache)
	if err != nil {
		fmt.Println("Got error --- ", err)
		return err
	}
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	config.next = locations.Next
	config.previous = locations.Previous
	return nil
}

func commandMapBack(config *configuration, cache *pokecache.Cache, arg string) error {
	locations, err := pokeapi.GetLocations(config.previous, cache)
	if err != nil {
		fmt.Println("Got error --- ", err)
		return err
	}
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	config.next = locations.Next
	config.previous = locations.Previous
	return nil
}

func commandExplore(config *configuration, cache *pokecache.Cache, arg string) error {
	resultJSON, err := pokeapi.ExploreLocation(arg, cache)

	if err != nil {
		fmt.Println("Got error --- ", err)
		return err
	}
	for _, encounter := range resultJSON.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *configuration, cache *pokecache.Cache, arg string) error {

	result, err := pokeapi.CatchPokemon(arg, cache)
	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	if err != nil {
		fmt.Println("Got error --- ", err)
		return err
	}
	if result {
		fmt.Printf("%s was caught!\n", arg)
	} else {
		fmt.Printf("%s escaped!\n",arg )
	}

	return nil
}
