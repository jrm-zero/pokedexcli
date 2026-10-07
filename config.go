package main

import (
	//"time"
	"github.com/jrm-zero/pokedexcli/internal/pokecache"
	"github.com/jrm-zero/pokedexcli/internal/pokeapicontrols"
)

type config struct {
	commands	map[string]cliCommand
	next_page	string
	previous_page	string
	Cache pokecache.Cache
	Pokedex map[string]pokeapicontrols.Pokemon
}

func getcommands() map[string]cliCommand {
		return map[string]cliCommand{
		"exit": {
			name:	"exit",
			description:	"Exit the Pokedex",
			callback:	commandExit,
		},
		"help": {
			name:	"help",
			description:	"Catalog of commands",
			callback:	commandHelp,
		},
		"map": {
			name:	"map",
			description:	"Displays next 20 pages",
			callback:	commandMap,
		},
		"mapb": {
			name:	"mapb",
			description:	"Displays previous 20 pages",
			callback:	commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Shows pokemon within a specified area location 'explore <location name>'",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempts to catch pokemon specified 'catch <pokemon name>'",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Prints basic stats of pokemon that have been caught and are in the pokedex 'inspect <pokemon name>'",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists all pokemon in the users pokedex",
			callback: commandPokedex,
		},
	}
}