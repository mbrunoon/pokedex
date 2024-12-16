package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/mbrunoon/pokedex/pokeapi"
)

type config struct {
	pokeApiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		cmd, ok := getCommands()[words[0]]
		if !ok {
			fmt.Println("Unknow command")
			continue
		}

		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

		continue
	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Pokedex Help",
			callback:    commandHelp,
		},
		"map": {
				name: "map",
				description: "Get locations list",
				callback: commandMapf,
		},
		"mapb": {
				name: "mapb",
				description: "Previous Locations list",
				callback: commandMapb,
		},
	}
}