package main

import (
	//"time"
	"github.com/jrm-zero/pokedexcli/internal/pokecache"
)

type config struct {
	commands	map[string]cliCommand
	next_page	string
	previous_page	string
	Cache pokecache.Cache
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
	}
}