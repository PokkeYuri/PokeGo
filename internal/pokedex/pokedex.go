package pokedex

import (
	"PokeGo/internal/pokeapi"
	"sync"
)

type Pokedex struct {
	mu       *sync.Mutex
	counters map[string]pokeapi.Pokemon
}

func (pd Pokedex) Add(p pokeapi.Pokemon) {
	pd.mu.Lock()
	pd.counters[p.Name] = p
	pd.mu.Unlock()
}

func (pd Pokedex) Get(name string) (pokeapi.Pokemon, bool) {
	pd.mu.Lock()
	poke, ok := pd.counters[name]
	pd.mu.Unlock()
	return poke, ok
}

func (pd Pokedex) GetAll() map[string]pokeapi.Pokemon {
	pd.mu.Lock()
	defer pd.mu.Unlock()
	return pd.counters
}
func NewPokedex() Pokedex {
	return Pokedex{
		mu:       &sync.Mutex{},
		counters: make(map[string]pokeapi.Pokemon),
	}
}
