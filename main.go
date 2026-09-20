package main

import (
	"fmt"
	"bufio"
	"repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for ;; {
		fmt.Print("Pokedex >")
		input := Scanner.Text(Scanner.Scan(scanner))
		cleanedInput := repl.cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanedInput[0])
	}
}

