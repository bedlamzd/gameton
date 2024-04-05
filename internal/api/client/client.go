package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/bedlamzd/gameton/internal/api/model"
)

const (
	AuthToken = "660d33aaee2c7660d33aaee2cc"
	Host      = "https://datsedenspace.datsteam.dev"
)

var client = &http.Client{}

func addAuth(req *http.Request) {
	req.Header.Set("X-Auth-Token", AuthToken)
}

func GetUniverse() model.UniverseResponse {
	req, _ := http.NewRequest("GET", Host+"/player/universe", nil)
	addAuth(req)

	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result model.UniverseResponse
	json.Unmarshal(body, &result)

	return result
}

func Travel(path *[]string) (model.TravelResponse, error) {
	postBody, err := json.Marshal(map[string][]string{"planets": *path})
	if err != nil {
		return model.TravelResponse{}, err
	}
	req, err := http.NewRequest("POST", Host+"/player/travel", bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return model.TravelResponse{}, err
	}
	addAuth(req)

	res, err := client.Do(req)
	if err != nil {
		return model.TravelResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return model.TravelResponse{}, err
	}

	var result model.TravelResponse
	json.Unmarshal(body, &result)
	return result, nil
}

func Collect(garbage *model.GarbageMap) (model.CollectResponse, error) {
	postBody, err := json.Marshal(map[string]model.GarbageMap{"garbage": *garbage})
	if err != nil {
		return model.CollectResponse{}, err
	}
	req, err := http.NewRequest("POST", Host+"/player/collect", bytes.NewBuffer(postBody))
	if err != nil {
		return model.CollectResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	addAuth(req)

	res, err := client.Do(req)
	if err != nil {
		return model.CollectResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return model.CollectResponse{}, err
	}

	var r map[string]any
	json.Unmarshal(body, &r)

	var result model.CollectResponse
	json.Unmarshal(body, &result)
	return result, nil
}

func GetRounds() model.RoundsResponse {
	req, _ := http.NewRequest("GET", Host+"/player/rounds", nil)
	addAuth(req)

	res, _ := client.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result model.RoundsResponse
	json.Unmarshal(body, &result)
	return result
}

func ResetRound() model.ResetRoundResponse {
	req, _ := http.NewRequest("GET", Host+"/player/reset", nil)
	addAuth(req)

	res, _ := client.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result model.ResetRoundResponse
	json.Unmarshal(body, &result)
	return result
}
