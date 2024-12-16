package main

import(
	"time"
	"github.com/mbrunoon/pokedex/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeApiClient: client,
	}

	startRepl(cfg)
}