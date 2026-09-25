package pokeapicontrols

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Name struct {
	name	string
	language	NamedAPIResource
}

type NamedAPIResource struct {
	name	string
	url		string
}

type EncounterMethodRate struct {
	encounter_method	NamedAPIResource
	version_details		EncounterVersionDetails
}

type EncounterVersionDetails struct {
	rate	int
	version		NamedAPIResource
}

type VersionEncounterDetail struct {
	version	NamedAPIResource
	max_chance	int
	encounter_details	Encounter
}

type Encounter struct {
	min_level	int
	max_level	int
	condition_values	NamedAPIResource
	chance	int
	method	NamedAPIResource
}

type PokemonEncounter struct {
	pokemon	NamedAPIResource
	version_details	VersionEncounterDetail
}
type Location struct {
	id	int `json:"id"`
	name	string	`json:"name"`
	game_index	int	`json:"game_index"`
	encounter_method_rates	EncounterMethodRate	`json:"encounter_method_rates"`
	location	NamedAPIResource	`"json:"location"`
	names	Name	`json:"names"`
	pokemon_encounters	PokemonEncounter	`json:"pokemon_encounters"`
}

func retreiveLocations(locationsOffset string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationsOffset)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	locations, err := decodeJSONResponse(res)
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location.name)
	}

	return nil
}

func decodeJSONResponse(res *http.Response) ([]Location, error) {
	var locations []Location
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return nil, err
	}
	return locations, nil
}