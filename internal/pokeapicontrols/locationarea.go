package pokeapicontrols

import (
	"encoding/json"
	"net/http"
	//"fmt"
	"time"
	"io"
	"github.com/jrm-zero/pokedexcli/internal/pokecache"
	"errors"
)

type LocationArea struct {
	ID                   int    `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	GameIndex            int    `json:"game_index,omitempty"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"encounter_method,omitempty"`
		VersionDetails []struct {
			Rate    int `json:"rate,omitempty"`
			Version struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version,omitempty"`
		} `json:"version_details,omitempty"`
	} `json:"encounter_method_rates,omitempty"`
	Location struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"location,omitempty"`
	Names []struct {
		Name     string `json:"name,omitempty"`
		Language struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"language,omitempty"`
	} `json:"names,omitempty"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"pokemon,omitempty"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version,omitempty"`
			MaxChance        int `json:"max_chance,omitempty"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level,omitempty"`
				MaxLevel int `json:"max_level,omitempty"`
				Chance   int `json:"chance,omitempty"`
				Method   struct {
					Name string `json:"name,omitempty"`
					URL  string `json:"url,omitempty"`
				} `json:"method,omitempty"`
				ConditionValues []any `json:"condition_values,omitempty"`
				PokemonDetails  any   `json:"pokemon_details,omitempty"`
			} `json:"encounter_details,omitempty"`
		} `json:"version_details,omitempty"`
	} `json:"pokemon_encounters,omitempty"`
}

func ExploreArea(url string, c *pokecache.Cache) ([]string, error) {
	_, exists := c.Cache[url]
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode == 404 {
			return nil, errors.New("Location does not exist.")
		}

		newVal, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		c.Cache[url] = pokecache.CacheEntry {
			CreatedAt: time.Now(),
			Val: newVal,
		}
	}
	body := c.Cache[url]

	res_decoded, err := decodeJSONResponseLocationArea(body.Val)
	if err != nil {
		return nil, err
	}

	var pokemon []string
	for _, p := range res_decoded.PokemonEncounters {
		pokemon = append(pokemon, p.Pokemon.Name)
	}

	return pokemon, nil
}

func decodeJSONResponseLocationArea(jsondata []byte) (LocationArea, error) {
	var locationDetails LocationArea
	if err := json.Unmarshal(jsondata, &locationDetails); err != nil {
		return LocationArea{}, err
	}
	
	return locationDetails, nil
}