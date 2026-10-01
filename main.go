package main

func main() {
	myConfigPtr := &config{
		commands: getcommands(),
		next_page: "https://pokeapi.co/api/v2/location-area",
		previous_page: "",
	}
	startRepl(myConfigPtr)
}