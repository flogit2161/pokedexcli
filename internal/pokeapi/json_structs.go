package pokeapi

// LocationArea JSON Struct for locations
type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

// AreaInfo JSON Struct for Pokemon list
type AreaInfo struct {
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}
