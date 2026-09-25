package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()

		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			fmt.Println("Type something next time, huh?")
			continue
		}

		command, exists := config.commands[cleanedInput[0]]

		if exists {
			err := command.callback(config)
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
	callback	func(config *config) error
}

