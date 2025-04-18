package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)



type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(url string)  (Location, error){
	fmt.Println("In pokeapi package")
	var apiurl string
	if len(url) == 0 {
		apiurl = "https://pokeapi.co/api/v2/location-area/"
	}else{
		apiurl = url
	}

	res, err := http.Get(apiurl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var loc Location
	if err= json.Unmarshal(body, &loc); err != nil {
		return loc,err
	}

	return loc,nil
}