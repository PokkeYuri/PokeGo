# PokeGo: A CLI Pokemon Adventure

## Overview
Dive into the world of Pokemon with PokeGo, a fun and simple CLI game written in Go! Explore different locations, catch a variety of Pokemon, and build your own Pokedex, all from your terminal.

The primary purpose of this project is for my own learning and development. It serves as an exercise to explore and practice programming in Go and uses the [PokéAPI](https://pokeapi.co/) for fetching data about Pokémon.

## Features
- **Explore Various Locations**: Travel through numerous areas in the Pokemon world to discover and catch Pokemon.
- **Catch and Release Pokemons**: Use your skills to catch Pokemon and add them to your Pokedex.
- **Build Your Pokedex**: Keep track of all the Pokemon you've caught and learn more about each one.

## Installation

To get started with PokeGo, clone this repository and build the project:
1. `git clone https://github.com/PokkeYuri/PokeGo.git`
2. `cd PokeGo`
3. `go build`

### Prerequisites
- Go 1.20 or later

## Usage
To start your adventure, run the following command: 
- LINUX: `./PokeGo`
- WINDOWS: `./PokeGo.exe`

## Commands
- `help`: Displays a help message with available commands.
- `exit`: Exit the Pokedex and the game.
- `map`: Shows the next 20 location areas to explore.
- `mapb`: Displays the previous 20 location areas.
- `explore <location_area_name>`: List all Pokemons in the specified location area.
- `catch <pokemon-name>`: Attempt to catch a named Pokemon.
- `inspect <pokemon-name>`: Inspect a caught Pokemon's details.
- `pokedex`: View all caught Pokemons in your Pokedex.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Disclaimer
"Pokémon" is a trademark of Nintendo, Creatures Inc, and GAME FREAK inc. This project is a fan-made game and is not affiliated, associated, authorized, endorsed by, or in any way officially connected with the Pokémon Company, Nintendo, Creatures, or Game Freak, or any of their subsidiaries or affiliates. The official Pokémon website can be found at [https://www.pokemon.com](https://www.pokemon.com). The names "Nintendo", "Pokémon", and related names, marks, emblems, and images are registered trademarks of their respective owners.
