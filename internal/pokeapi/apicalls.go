package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/ArtemisNyx3/pokedexcli/internal/pokecache"
)

const location_area_url = "https://pokeapi.co/api/v2/location-area/"

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
type Explore struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocations(url string, cache *pokecache.Cache) (Location, error) {

	var apiurl string
	if len(url) == 0 {
		apiurl = location_area_url
	} else {
		apiurl = url
	}

	// check cache
	data, ok := cache.Get(apiurl)
	if !ok {
		res, err := http.Get(apiurl)
		if err != nil {
			log.Fatal(err)
		}
		data, err = io.ReadAll(res.Body)
		cache.Add(apiurl, data)
		res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	var loc Location
	if err := json.Unmarshal(data, &loc); err != nil {
		return loc, err
	}

	return loc, nil
}

func ExploreLocation(areaName string, cache *pokecache.Cache) (Explore, error) {

	var exploreData Explore
	if areaName == "" {
		return exploreData, errors.New("Area name is empty")
	}

	apiurl := location_area_url + areaName

	data, ok := cache.Get(apiurl)
	if !ok {
		res, err := http.Get(apiurl)
		if err != nil {
			log.Fatal(err)
		}
		data, err = io.ReadAll(res.Body)
		cache.Add(apiurl, data)
		res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := json.Unmarshal(data, &exploreData); err != nil {
		return exploreData, err
	}

	return exploreData, nil

}
