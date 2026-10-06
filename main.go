package main

import (
	"github.com/jrm-zero/pokedexcli/internal/pokecache"
	"github.com/jrm-zero/pokedexcli/internal/pokeapicontrols"
	"time"
)

func main() {
	myConfigPtr := &config{
		commands: getcommands(),
		next_page: "https://pokeapi.co/api/v2/location-area",
		previous_page: "",
		Cache: pokecache.NewCache(4 * time.Millisecond),
		Pokedex: make(map[string]pokeapicontrols.Pokemon),
	}
	startRepl(myConfigPtr)
}