package main

type config struct {
	commands	map[string]cliCommand
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
	}
}