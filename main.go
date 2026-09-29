package main

import (
	"time"
	"github.com/jrm-zero/pokedexcli/internal/pokecache"
)
func main() {
	myConfigPtr := &config{
		commands: getcommands(),
		next_page: "https://pokeapi.co/api/v2/location-area",
		previous_page: "",
		cache: newCache(5 * time.Second),
	}
	startRepl(myConfigPtr)
}