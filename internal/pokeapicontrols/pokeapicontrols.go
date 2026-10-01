package pokeapicontrols

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"io"
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

func GetLocations(url string, page_direction int, c *pokecache.Cache) ([]string, string, error) {
	body, exists := c.Cache[url]
	if !exists {
		fmt.Println(c.Cache[url])
		res, err := http.Get(url)
		if err != nil {
			return nil, "", err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, "", err
		}

		c.Cache[url] = pokecache.CacheEntry {
			CreatedAt: time.Now(),
			Val: body,
		}
	}
	
	res_decoded, err := decodeJSONResponse(body.Val)
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


func decodeJSONResponse(jsondata []byte) (LocationAreaBatch, error) {
	var locations LocationAreaBatch
	if err := json.Unmarshal(jsondata, &locations); err != nil {
		return LocationAreaBatch{}, err
	}
	
	return locations, nil
}