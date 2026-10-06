package pokeapicontrols

import (
	"encoding/json"
	"net/http"
	//"fmt"
	"time"
	"io"
	"github.com/jrm-zero/pokedexcli/internal/pokecache"
)

type Pokemon struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	BaseExperience int    `json:"base_experience,omitempty"`
	Height         int    `json:"height,omitempty"`
	IsDefault      bool   `json:"is_default,omitempty"`
	Order          int    `json:"order,omitempty"`
	Weight         int    `json:"weight,omitempty"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden,omitempty"`
		Slot     int  `json:"slot,omitempty"`
		Ability  struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"ability,omitempty"`
	} `json:"abilities,omitempty"`
	PastAbilities []struct {
		Generation struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"generation,omitempty"`
		Abilities []struct {
			IsHidden bool `json:"is_hidden,omitempty"`
			Slot     int  `json:"slot,omitempty"`
			Ability  any  `json:"ability,omitempty"`
		} `json:"abilities,omitempty"`
	} `json:"past_abilities,omitempty"`
	Forms []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"forms,omitempty"`
	GameIndices []struct {
		GameIndex int `json:"game_index,omitempty"`
		Version   struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"version,omitempty"`
	} `json:"game_indices,omitempty"`
	HeldItems              []any  `json:"held_items,omitempty"`
	LocationAreaEncounters string `json:"location_area_encounters,omitempty"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"move,omitempty"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at,omitempty"`
			VersionGroup   struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version_group,omitempty"`
			MoveLearnMethod struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"move_learn_method,omitempty"`
			Order any `json:"order,omitempty"`
		} `json:"version_group_details,omitempty"`
	} `json:"moves,omitempty"`
	Species struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"species,omitempty"`
	Sprites struct {
		Other struct {
			Home struct {
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontFemale      string `json:"front_female,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontShinyFemale string `json:"front_shiny_female,omitempty"`
			} `json:"home,omitempty"`
			Showdown struct {
				BackShiny        string `json:"back_shiny,omitempty"`
				BackFemale       string `json:"back_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				BackDefault      string `json:"back_default,omitempty"`
				FrontFemale      string `json:"front_female,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				BackShinyFemale  string `json:"back_shiny_female,omitempty"`
				FrontShinyFemale string `json:"front_shiny_female,omitempty"`
			} `json:"showdown,omitempty"`
			DreamWorld struct {
				FrontFemale  any    `json:"front_female,omitempty"`
				FrontDefault string `json:"front_default,omitempty"`
			} `json:"dream_world,omitempty"`
			OfficialArtwork struct {
				Versions struct {
					GenerationI struct {
						RedAndBlue struct {
							FrontDefault string `json:"front_default,omitempty"`
						} `json:"red-and-blue,omitempty"`
						RedAndGreen struct {
							FrontDefault string `json:"front_default,omitempty"`
						} `json:"red-and-green,omitempty"`
					} `json:"generation-i,omitempty"`
					GenerationIi struct {
						GoldAndSilver struct {
							FrontDefault any `json:"front_default,omitempty"`
						} `json:"gold-and-silver,omitempty"`
					} `json:"generation-ii,omitempty"`
				} `json:"versions,omitempty"`
				FrontShiny   string `json:"front_shiny,omitempty"`
				FrontDefault string `json:"front_default,omitempty"`
			} `json:"official-artwork,omitempty"`
		} `json:"other,omitempty"`
		Versions struct {
			GenerationI struct {
				Yellow struct {
					BackGbc              string `json:"back_gbc,omitempty"`
					BackGray             string `json:"back_gray,omitempty"`
					FrontGbc             string `json:"front_gbc,omitempty"`
					FrontGray            string `json:"front_gray,omitempty"`
					BackDefault          string `json:"back_default,omitempty"`
					FrontDefault         string `json:"front_default,omitempty"`
					BackTransparent      string `json:"back_transparent,omitempty"`
					FrontTransparent     string `json:"front_transparent,omitempty"`
					BackTransparentGray  string `json:"back_transparent_gray,omitempty"`
					FrontTransparentGray string `json:"front_transparent_gray,omitempty"`
				} `json:"yellow,omitempty"`
				RedBlue struct {
					BackGray             string `json:"back_gray,omitempty"`
					FrontGray            string `json:"front_gray,omitempty"`
					BackDefault          string `json:"back_default,omitempty"`
					FrontDefault         string `json:"front_default,omitempty"`
					BackTransparent      string `json:"back_transparent,omitempty"`
					FrontTransparent     string `json:"front_transparent,omitempty"`
					BackTransparentGray  string `json:"back_transparent_gray,omitempty"`
					FrontTransparentGray string `json:"front_transparent_gray,omitempty"`
				} `json:"red-blue,omitempty"`
				RedGreenJapan struct {
					BackGray     string `json:"back_gray,omitempty"`
					FrontGray    string `json:"front_gray,omitempty"`
					BackDefault  string `json:"back_default,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"red-green-japan,omitempty"`
			} `json:"generation-i,omitempty"`
			GenerationV struct {
				Icons struct {
					Animated struct {
						FrontDefault string `json:"front_default,omitempty"`
					} `json:"animated,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"icons,omitempty"`
				BlackWhite struct {
					Animated struct {
						BackShiny        string `json:"back_shiny,omitempty"`
						BackFemale       string `json:"back_female,omitempty"`
						FrontShiny       string `json:"front_shiny,omitempty"`
						BackDefault      string `json:"back_default,omitempty"`
						FrontFemale      string `json:"front_female,omitempty"`
						FrontDefault     string `json:"front_default,omitempty"`
						BackShinyFemale  string `json:"back_shiny_female,omitempty"`
						FrontShinyFemale string `json:"front_shiny_female,omitempty"`
					} `json:"animated,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"black-white,omitempty"`
			} `json:"generation-v,omitempty"`
			GenerationIi struct {
				Gold struct {
					BackShiny             string `json:"back_shiny,omitempty"`
					FrontShiny            string `json:"front_shiny,omitempty"`
					BackDefault           string `json:"back_default,omitempty"`
					FrontDefault          string `json:"front_default,omitempty"`
					BackTransparent       string `json:"back_transparent,omitempty"`
					FrontTransparent      string `json:"front_transparent,omitempty"`
					BackShinyTransparent  string `json:"back_shiny_transparent,omitempty"`
					FrontShinyTransparent string `json:"front_shiny_transparent,omitempty"`
				} `json:"gold,omitempty"`
				Silver struct {
					BackShiny             string `json:"back_shiny,omitempty"`
					FrontShiny            string `json:"front_shiny,omitempty"`
					BackDefault           string `json:"back_default,omitempty"`
					FrontDefault          string `json:"front_default,omitempty"`
					BackTransparent       string `json:"back_transparent,omitempty"`
					FrontTransparent      string `json:"front_transparent,omitempty"`
					BackShinyTransparent  string `json:"back_shiny_transparent,omitempty"`
					FrontShinyTransparent string `json:"front_shiny_transparent,omitempty"`
				} `json:"silver,omitempty"`
				Crystal struct {
					Animated struct {
						FrontShiny   string `json:"front_shiny,omitempty"`
						FrontDefault string `json:"front_default,omitempty"`
					} `json:"animated,omitempty"`
					BackShiny             string `json:"back_shiny,omitempty"`
					FrontShiny            string `json:"front_shiny,omitempty"`
					BackDefault           string `json:"back_default,omitempty"`
					FrontDefault          string `json:"front_default,omitempty"`
					BackTransparent       string `json:"back_transparent,omitempty"`
					FrontTransparent      string `json:"front_transparent,omitempty"`
					BackShinyTransparent  string `json:"back_shiny_transparent,omitempty"`
					FrontShinyTransparent string `json:"front_shiny_transparent,omitempty"`
				} `json:"crystal,omitempty"`
			} `json:"generation-ii,omitempty"`
			GenerationIv struct {
				Icons struct {
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"icons,omitempty"`
				Platinum struct {
					Animated struct {
						FrontShiny       string `json:"front_shiny,omitempty"`
						FrontFemale      string `json:"front_female,omitempty"`
						FrontDefault     string `json:"front_default,omitempty"`
						FrontShinyFemale string `json:"front_shiny_female,omitempty"`
					} `json:"animated,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"platinum,omitempty"`
				DiamondPearl struct {
					Animated struct {
						FrontShiny       string `json:"front_shiny,omitempty"`
						FrontFemale      string `json:"front_female,omitempty"`
						FrontDefault     string `json:"front_default,omitempty"`
						FrontShinyFemale string `json:"front_shiny_female,omitempty"`
					} `json:"animated,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"diamond-pearl,omitempty"`
				HeartgoldSoulsilver struct {
					Animated struct {
						FrontShiny       string `json:"front_shiny,omitempty"`
						FrontFemale      string `json:"front_female,omitempty"`
						FrontDefault     string `json:"front_default,omitempty"`
						FrontShinyFemale string `json:"front_shiny_female,omitempty"`
					} `json:"animated,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"heartgold-soulsilver,omitempty"`
			} `json:"generation-iv,omitempty"`
			GenerationIx struct {
				Champions struct {
					FrontShiny   string `json:"front_shiny,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"champions,omitempty"`
				ScarletViolet struct {
					FrontFemale  any    `json:"front_female,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"scarlet-violet,omitempty"`
			} `json:"generation-ix,omitempty"`
			GenerationVi struct {
				XY struct {
					BackShiny        string `json:"back_shiny,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"x-y,omitempty"`
				Icons struct {
					FrontFemale  any    `json:"front_female,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"icons,omitempty"`
				OmegarubyAlphasapphire struct {
					BackShiny        string `json:"back_shiny,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"omegaruby-alphasapphire,omitempty"`
			} `json:"generation-vi,omitempty"`
			GenerationIii struct {
				Icons struct {
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"icons,omitempty"`
				Emerald struct {
					Animated struct {
						BackShiny    string `json:"back_shiny,omitempty"`
						FrontShiny   string `json:"front_shiny,omitempty"`
						BackDefault  string `json:"back_default,omitempty"`
						FrontDefault string `json:"front_default,omitempty"`
					} `json:"animated,omitempty"`
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
					BackDefault  string `json:"back_default,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"emerald,omitempty"`
				RubySapphire struct {
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
					BackDefault  string `json:"back_default,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"ruby-sapphire,omitempty"`
				FireredLeafgreen struct {
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
					BackDefault  string `json:"back_default,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"firered-leafgreen,omitempty"`
			} `json:"generation-iii,omitempty"`
			GenerationVii struct {
				Icons struct {
					FrontFemale  any    `json:"front_female,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"icons,omitempty"`
				UltraSunUltraMoon struct {
					BackShiny        string `json:"back_shiny,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"ultra-sun-ultra-moon,omitempty"`
				LetsGoPikachuLetsGoEevee struct {
					Icons struct {
						FrontDefault string `json:"front_default,omitempty"`
					} `json:"icons,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"lets-go-pikachu-lets-go-eevee,omitempty"`
			} `json:"generation-vii,omitempty"`
			GenerationViii struct {
				Icons struct {
					FrontFemale  any    `json:"front_female,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"icons,omitempty"`
				BrilliantDiamondShiningPearl struct {
					FrontDefault string `json:"front_default,omitempty"`
				} `json:"brilliant-diamond-shining-pearl,omitempty"`
			} `json:"generation-viii,omitempty"`
		} `json:"versions,omitempty"`
		BackShiny        string `json:"back_shiny,omitempty"`
		BackFemale       string `json:"back_female,omitempty"`
		FrontShiny       string `json:"front_shiny,omitempty"`
		BackDefault      string `json:"back_default,omitempty"`
		FrontFemale      string `json:"front_female,omitempty"`
		FrontDefault     string `json:"front_default,omitempty"`
		BackShinyFemale  string `json:"back_shiny_female,omitempty"`
		FrontShinyFemale string `json:"front_shiny_female,omitempty"`
	} `json:"sprites,omitempty"`
	Cries struct {
		Latest string `json:"latest,omitempty"`
		Legacy string `json:"legacy,omitempty"`
	} `json:"cries,omitempty"`
	Stats []struct {
		BaseStat int `json:"base_stat,omitempty"`
		Effort   int `json:"effort,omitempty"`
		Stat     struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"stat,omitempty"`
	} `json:"stats,omitempty"`
	PastStats []struct {
		Generation struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"generation,omitempty"`
		Stats []struct {
			BaseStat int `json:"base_stat,omitempty"`
			Effort   int `json:"effort,omitempty"`
			Stat     struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"stat,omitempty"`
		} `json:"stats,omitempty"`
	} `json:"past_stats,omitempty"`
	Types []struct {
		Slot int `json:"slot,omitempty"`
		Type struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"type,omitempty"`
	} `json:"types,omitempty"`
	PastTypes []any `json:"past_types,omitempty"`
}


func GetPokemon(url string, c *pokecache.Cache) (Pokemon, error) {
	_, exists := c.Cache[url]
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		newVal, err := io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.Cache[url] = pokecache.CacheEntry {
			CreatedAt: time.Now(),
			Val: newVal,
		}
	}
	body := c.Cache[url]

	res_decoded, err := decodeJSONResponsePokemon(body.Val)
	if err != nil {
		return Pokemon{}, err
	}

	return res_decoded, nil
}

func decodeJSONResponsePokemon(jsondata []byte) (Pokemon, error) {
	var pokemon Pokemon
	if err := json.Unmarshal(jsondata, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}