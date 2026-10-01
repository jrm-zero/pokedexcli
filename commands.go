package main

import (
	"fmt"
	"os"
	"github.com/jrm-zero/pokedexcli/internal/pokeapicontrols"
	//"github.com/jrm-zero/pokedexcli/internal/pokecache"
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
	locations, next_url, err := pokeapicontrols.GetLocations(config.next_page, 1, &config.Cache)
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location)
	}
	
	config.previous_page = config.next_page
	config.next_page = next_url
	return nil
}

func commandMapb(config *config) error {
	locations, previous_url, err := pokeapicontrols.GetLocations(config.previous_page, 0, &config.Cache)
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location)
	}
	
	config.next_page = config.previous_page
	config.previous_page = previous_url
	return nil
}