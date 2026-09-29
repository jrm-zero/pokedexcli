package main

import (
	"fmt"
	"os"
	"github.com/jrm-zero/pokedexcli/internal/pokeapicontrols"
	"github.com/jrm-zero/pokedexcli/internal/pokecache"
)

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	
	for _, command := range config.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(config *config) error {
	res, exists := config.cache[config.next_page]
	if !exists {
		err := pokeapicontrols.GetLocations(config.next_page)
		if err != nil {
			return err
		}
		config.cache[config.next_page]
	}

	locations, err := decodeJSONResponse(res)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	
	config.previous_page = config.next_page
	config.next_page = locations.Next
	return nil
}

func commandMapb(config *config) error {
	res, exists := config.cache[config.previous_page]
	if !exists {
		previous_url, err := pokeapicontrols.GetLocations(config.previous_page)
		if err != nil {
			return err
		}
		config.cache[config.previous_page]
	}

	locations, err := decodeJSONResponse(res)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	
	config.next_page = config.previous_page
	config.previous_page = locations.Previous
	return nil
}