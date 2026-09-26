package pokeapicontrols

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func RetreiveLocations(url string, page_direction int) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	locations, err := decodeJSONResponse(res)
	if err != nil {
		return "", err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	
	if page_direction == 1 {
		return locations.Next, nil
	}
	
	return locations.Previous, nil
}

func decodeJSONResponse(res *http.Response) (LocationAreaBatch, error) {
	var locations LocationAreaBatch
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return LocationAreaBatch{}, err
	}
	
	return locations, nil
}