package pokeapicontrols

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/jrm-zero/pokedexcli/internal/pokecache"
)

type LocationAreaBatch struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string    `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"results,omitempty"`
}

func GetLocations(url string, page_direction int) ([]string, string, error) {
	cache := NewCache(5 * time.Second)
	res, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer res.Body.Close()
	cache.Add(url, res)
	res_decoded, err := decodeJSONResponse(res)
	if err != nil {
		return nil, "", err
	}

	var locations []string
	for _, location_result := range res_decoded.Results {
		locations = append(locations, location_result.Name)
	}

	if page_direction == 1 {
		return locations, res_decoded.Next, nil
	}
	
	return locations, res_decoded.Previous, nil
}


func decodeJSONResponse(res *http.Response) (LocationAreaBatch, error) {
	var locations LocationAreaBatch
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return LocationAreaBatch{}, err
	}
	
	return locations, nil
}