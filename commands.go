package main

import (
	"fmt"
	"os"
	"github.com/jrm-zero/pokedexcli/internal/pokeapicontrols"
	"errors"
	"math/rand"
	//"github.com/jrm-zero/pokedexcli/internal/pokecache"
)

func commandExit(config *config, arguments ...string) error {
	if len(arguments) != 0 {
		return errors.New("Command exit takes no arguments.")
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, arguments ...string) error {
	if len(arguments) != 0 {
		return errors.New("Command help takes no arguments.")
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	
	for _, command := range config.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(config *config, arguments ...string) error {
	if len(arguments) != 0 {
		return errors.New("Command map takes no arguments.")
	}
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

func commandMapb(config *config, arguments ...string) error {
	if len(arguments) != 0 {
		return errors.New("Command mapb takes no arguments.")
	}
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

func commandExplore(config *config, arguments ...string) error {
	if len(arguments) != 1 {
		return errors.New("Command explore takes one argument 'explore <location area>.")
	}
	fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", arguments[0])
	pokemon, err := pokeapicontrols.ExploreArea(fullURL, &config.Cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", arguments[0])
	fmt.Println("Found Pokemon:")
	for _, p := range pokemon {
		fmt.Println(p)
	}

	return nil
}

func commandCatch(config *config, arguments ...string) error {
	if len(arguments) != 1 {
		return errors.New("Command catch takes one argument 'catch <pokemon name>")
	}
	fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", arguments[0])
	pokemon, err := pokeapicontrols.GetPokemon(fullURL, &config.Cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", arguments[0])
	caught := catchOutcome(pokemon.BaseExperience)
	if caught {
		fmt.Printf("%s was caught!\n", arguments[0])
		config.Pokedex[arguments[0]] = pokemon
		return nil
	} 

	fmt.Printf("%s escaped!\n", arguments[0])
	return nil
}

func catchOutcome(baseExperience int) bool {
	t := rand.Intn(400)
	if t < baseExperience {
		return false
	}

	return true
}

func commandInspect(config *config, arguments ...string) error {
	pokemon, exists := config.Pokedex[arguments[0]]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("	-%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("	-%s\n", t.Type.Name)
	}

	return nil
}