package pokedex

import (
	"PokeGo/internal/pokeapi"
	"sync"
)

type Pokedex struct {
	mu       *sync.Mutex
	Counters map[string]pokeapi.Pokemon
}

func (pd Pokedex) Add(p pokeapi.Pokemon) {
	pd.mu.Lock()
	defer pd.mu.Unlock()
	pd.Counters[p.Name] = p
}

func NewPokedex() Pokedex {
	return Pokedex{
		mu:       &sync.Mutex{},
		Counters: make(map[string]pokeapi.Pokemon),
	}
}
