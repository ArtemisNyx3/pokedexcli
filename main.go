package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// REPL Loop
	// var userInput string
	// scanner := bufio.NewScanner(strings.NewReader(userInput))
	// fmt.Println("Scanned this --- %v", scanner)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		key := scanner.Text()
		userInput := cleanInput(key)
		fmt.Printf("Your command was: %s\n", userInput[0])
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
