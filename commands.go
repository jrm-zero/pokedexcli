package main

import (
	"fmt"
	"os"
	"github.com/jrm-zero/pokedexcli/internal/pokeapicontrols"
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
	retreiveLocations(string(config.locationsOffset))
	config.locationsOffset += 20
	return nil
}