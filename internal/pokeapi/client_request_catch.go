package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ClientRequestCatch(pokemonName string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemonName

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, nil
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return PokemonInfo{}, nil
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return PokemonInfo{}, nil
	}

	pokemonInfo := PokemonInfo{}
	err = json.Unmarshal(body, &pokemonInfo)
	if err != nil {
		return PokemonInfo{}, nil
	}

	return pokemonInfo, nil
}
