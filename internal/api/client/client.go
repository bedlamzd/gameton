package client

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bedlamzd/gameton/internal/api/model"
)

const (
	AuthToken = "660d33aaee2c7660d33aaee2cc"
	Host      = "https://datsedenspace.datsteam.dev"
)

var client = http.Client{}

func GetUniverse() model.UniverseResponse {
	req, _ := http.NewRequest("GET", Host+"/player/universe", nil)
	req.Header.Add("X-Auth-Token", AuthToken)

	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result model.UniverseResponse
	json.Unmarshal(body, &result)

	return result
}
