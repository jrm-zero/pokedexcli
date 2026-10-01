package main

import (
	"github.com/jrm-zero/pokedexcli/internal/pokecache"
	"time"
)

func main() {
	myConfigPtr := &config{
		commands: getcommands(),
		next_page: "https://pokeapi.co/api/v2/location-area",
		previous_page: "",
		Cache: pokecache.NewCache(5 * time.Second),
	}
	startRepl(myConfigPtr)
}