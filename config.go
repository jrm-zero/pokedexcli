package main

type config struct {
	commands	map[string]cliCommand
	next_page	string
	previous_page	string
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
	}
}