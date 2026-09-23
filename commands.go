package main

import (
	"fmt"
	"os"
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
	
	for _, command := range getcommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}