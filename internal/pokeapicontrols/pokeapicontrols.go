package pokeapicontrols

import (
	"encoding/json"
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

func GetLocations(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}

func decodeJSONResponse(res *http.Response) (LocationAreaBatch, error) {
	var locations LocationAreaBatch
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return LocationAreaBatch{}, err
	}
	
	return locations, nil
}