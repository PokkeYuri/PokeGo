package main

import (
	"PokeGo/internal/pokeapi"
	"PokeGo/internal/pokecache"
	"PokeGo/internal/pokedex"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var cache pokecache.Cache

var pokedexe pokedex.Pokedex

var cliParameters string

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Parameteters struct {
	nextMapUri, prevMapUri, locationAreaUri, pokemonUri string
}

var uriParameters = Parameteters{
	nextMapUri:      "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	prevMapUri:      "",
	locationAreaUri: "https://pokeapi.co/api/v2/location-area/",
	pokemonUri:      "https://pokeapi.co/api/v2/pokemon/",
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays name of next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays name of previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_area_name>",
			description: "List all Pokemons in named location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Catch Pokemon and adds to Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspects a Pokemon of a given name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all Pokemons who where catched",
			callback:    commandPokedex,
		},
	}
}

func commandMap() error {
	areas := pokeapi.GetLocationAreaList(uriParameters.nextMapUri, &cache)
	for i := 0; i < len(areas.Results); i++ {
		fmt.Println(areas.Results[i].Name)
	}
	uriParameters.nextMapUri = areas.Next
	if areas.Previous != nil {
		uriParameters.prevMapUri = *areas.Previous
	}

	return nil
}

func commandMapb() error {
	if uriParameters.prevMapUri == "" {
		fmt.Println("No previous Location found!")
		return nil
	}
	areas := pokeapi.GetLocationAreaList(uriParameters.prevMapUri, &cache)
	for i := 0; i < len(areas.Results); i++ {
		fmt.Println(areas.Results[i].Name)
	}
	uriParameters.nextMapUri = areas.Next
	if areas.Previous != nil {
		uriParameters.prevMapUri = *areas.Previous
	}
	return nil
}

func commandHelp() error {
	fmt.Println("\nWelcome to PokeGo!")

	for _, element := range getCommands() {
		fmt.Println(element.name + ": " + element.description)
	}
	fmt.Println("")
	return nil
}

func commandExplore() error {
	uri := uriParameters.locationAreaUri + cliParameters + "/"
	area, err := pokeapi.GetLocationArea(uri, &cache)
	if err != nil {
		fmt.Println("Location Area not found!")
		return nil
	}
	for i := 0; len(area.PokemonEncounters) > i; i++ {
		fmt.Println("- " + area.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}

func commandCatch() error {
	uri := uriParameters.pokemonUri + cliParameters + "/"

	_, alreadyOwned := pokedexe.Get(cliParameters)
	if alreadyOwned {
		fmt.Println("Pokemon already cought!")
		return nil
	}

	pokemon, err := pokeapi.GetPokemon(uri, &cache)
	if err != nil {
		fmt.Println("Pokemon not found!")
		return nil
	}
	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")

	//TODO Implement better algorythm with Pokemon.BaseExperience?
	x := rand.Intn(50)
	if x > 20 {
		pokedexe.Add(pokemon)
		fmt.Println(pokemon.Name + " was caught!")
	} else {
		fmt.Println(pokemon.Name + " escaped!")
	}

	return nil
}

func commandInspect() error {
	pokemon, ok := pokedexe.Get(cliParameters)
	if !ok {
		fmt.Println("you have not caught that pokemon!")
		return nil
	} else {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Println("Height: " + strconv.Itoa(pokemon.Height))
		fmt.Println("Weight: " + strconv.Itoa(pokemon.Height))
		fmt.Println("Stats:")
		for i := 0; i < len(pokemon.Stats); i++ {
			fmt.Println("  -" + pokemon.Stats[i].Stat.Name + ": " + strconv.Itoa(pokemon.Stats[i].BaseStat))
		}
		fmt.Println("Types:")
		for i := 0; i < len(pokemon.Types); i++ {
			fmt.Println("  -" + pokemon.Types[i].Type.Name)
		}
	}
	return nil
}

func commandPokedex() error {
	pokemonMap := pokedexe.GetAll()
	fmt.Println("Your Pokedex:")
	for k, _ := range pokemonMap {
		fmt.Println("  -" + k)
	}
	return nil
}

func commandExit() error {
	os.Exit(1)
	return nil
}

func main() {
	cache = pokecache.NewCache(time.Minute)
	pokedexe = pokedex.NewPokedex()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("PokeGo > ")
		scanner.Scan()
		inputText := scanner.Text()
		textArr := strings.Split(inputText, " ")

		var text string
		if len(textArr) > 1 {
			text = textArr[0]
			cliParameters = textArr[1]
		} else {
			text = inputText
		}

		if len(text) != 0 {
			if _, ok := getCommands()[text]; ok {
				getCommands()[text].callback()
			} else {
				fmt.Println("Command '" + text + "' not known. Please try again!")
			}
		}
	}
}
