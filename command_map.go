package main

import(
	"fmt"
	"errors"
)

func commandMapf(cfg *config) error {

	resLocations, err := cfg.pokeApiClient.ListLocations(cfg.nextLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resLocations.Next
	cfg.prevLocationsURL = resLocations.Previous

	for _, loc := range resLocations.Results {
		fmt.Println(loc.Name)
	}

	return nil

}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL != nil {
		return errors.New("this is the first page")
	}

	resLocations, err := cfg.pokeApiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resLocations.Next
	cfg.prevLocationsURL = resLocations.Previous

	for _, loc := range resLocations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}