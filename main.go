package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	temp := strings.Split(text, " ")

	var cleanIP []string
	for _,str := range(temp){
		if str == " " {
			continue
		}else if len(str) == 0 {
			continue
		}else{
			cleanIP = append(cleanIP, str)
		}
	}
	return cleanIP
}
