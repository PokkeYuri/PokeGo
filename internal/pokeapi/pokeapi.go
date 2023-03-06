package pokeapi

import (
	"PokeGo/internal/pokecache"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type LocationAreaList struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type PokeAnimal struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IsMainSeries bool   `json:"is_main_series"`
	Generation   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	EffectEntries []struct {
		Effect      string `json:"effect"`
		ShortEffect string `json:"short_effect"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`
	EffectChanges []struct {
		VersionGroup struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
		EffectEntries []struct {
			Effect   string `json:"effect"`
			Language struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
		} `json:"effect_entries"`
	} `json:"effect_changes"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		VersionGroup struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	Pokemon []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Pokemon  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon"`
}

func GetLocationAreaList(uri string, c *pokecache.Cache) LocationAreaList {
	var data []byte
	value, ok := c.Get(uri)
	if ok {
		data = value
	} else {
		data = HttpGet(uri)
		c.Add(uri, data)
	}
	locationAreasList := LocationAreaList{}
	err := json.Unmarshal(data, &locationAreasList)

	if err != nil {
		log.Fatal(err)
	}
	return locationAreasList
}

func GetLocationArea(uri string, c *pokecache.Cache) (LocationArea, error) {
	var data []byte
	value, ok := c.Get(uri)
	if ok {
		data = value
	} else {
		data = HttpGet(uri)
		c.Add(uri, data)
	}

	locationArea := LocationArea{}
	err := json.Unmarshal(data, &locationArea)

	if err != nil {
		//log.Fatal(err)
	}
	return locationArea, err
}

func GetPokemon(uri string, c *pokecache.Cache) (PokeAnimal, error) {
	var data []byte
	value, ok := c.Get(uri)
	if ok {
		data = value
	} else {
		data = HttpGet(uri)
		c.Add(uri, data)
	}

	pokeAnimal := PokeAnimal{}
	err := json.Unmarshal(data, &pokeAnimal)

	if err != nil {

	}
	return pokeAnimal, err
}

func HttpGet(uri string) []byte {
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		//log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		//log.Fatal(err)
	}

	return body
}
