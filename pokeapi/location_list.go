package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (ResLocations, error) {

	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl	
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResLocations{}, err
	}

	resLocations := ResLocations{}
	err = json.Unmarshal(data, &resLocations)
	if err != nil {
		return ResLocations{}, err
	}

	return resLocations, nil

}