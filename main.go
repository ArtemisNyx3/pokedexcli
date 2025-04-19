package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ArtemisNyx3/pokedexcli/internal/pokeapi"
	"github.com/ArtemisNyx3/pokedexcli/internal/pokecache"
)

var cliDirectory = make(map[string]cliCommand)

type configuration struct {
	next     string
	previous string
	cache    *pokecache.Cache
	pokedex  map[string]pokeapi.Pokemon
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
		},"inspect": {
			name:        "inspect",
			description: "inspect <pokemon-name> : Pokedex data for pokemon",
			callback: commandInspect,
		},
	}

	config := configuration{
		next:     "",
		previous: "",
		cache:    pokecache.NewCache(10 * time.Second),
		pokedex:  make(map[string]pokeapi.Pokemon),
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
			if len(userInput) == 1 {
				command.callback(&config, "")
			} else {
				command.callback(&config, userInput[1])
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
	callback    func(config *configuration, arg string) error
}

func commandExit(c *configuration, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *configuration, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for name, cmd := range cliDirectory {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}

func commandMap(config *configuration, arg string) error {
	locations, err := pokeapi.GetLocations(config.next, config.cache)
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

func commandMapBack(config *configuration, arg string) error {
	locations, err := pokeapi.GetLocations(config.previous, config.cache)
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

func commandExplore(config *configuration, arg string) error {
	resultJSON, err := pokeapi.ExploreLocation(arg, config.cache)

	if err != nil {
		fmt.Println("Got error --- ", err)
		return err
	}
	for _, encounter := range resultJSON.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *configuration, arg string) error {

	pokemon, err := pokeapi.CatchPokemon(arg, config.cache)
	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	if err != nil {
		fmt.Println("Got error --- ", err)
		return err
	}
	chance := float64(rand.Intn(pokemon.BaseExperience-0)) / float64(pokemon.BaseExperience)
	if chance >= 0.5 {
		fmt.Printf("%s was caught!\n", arg)
		config.pokedex[arg] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", arg)
	}

	return nil
}

func commandInspect(config *configuration, arg string) error {
	pokemon,ok := config.pokedex[arg]
	if !ok{
		fmt.Printf("Haven't caught %s\n",arg)
	}else{
		fmt.Printf("Name: %s\n",pokemon.Name)
		fmt.Printf("Height: %v\n",pokemon.Height)
		fmt.Printf("Weight: %v\n",pokemon.Weight)


	}
	return nil
}
