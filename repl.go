package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		command, exists := getcommands()[cleanedInput[0]]

		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(strings.ToLower(text))
	return cleanedInput
}

type cliCommand struct {
	name	string
	description	string
	callback	func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	
	for _, command := range getcommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
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