package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ClientRequest(pageUrl *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, nil
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationArea{}, nil
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationArea{}, nil
	}

	loc := LocationArea{}
	err = json.Unmarshal(body, &loc)
	if err != nil {
		return LocationArea{}, nil
	}

	return loc, nil
}
