package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ClientRequestExplore(areaName string) (AreaInfo, error) {
	url := baseURL + "/location-area/" + areaName

	//check cache
	val, ok := c.cache.Get(url)
	if ok == true {
		areaInfo := AreaInfo{}
		err := json.Unmarshal(val, &areaInfo)
		if err != nil {
			return AreaInfo{}, err
		}
		return areaInfo, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaInfo{}, nil
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return AreaInfo{}, nil
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return AreaInfo{}, nil
	}

	c.cache.Add(url, body)

	areaInfo := AreaInfo{}
	err = json.Unmarshal(body, &areaInfo)
	if err != nil {
		return AreaInfo{}, nil
	}

	return areaInfo, nil

}
