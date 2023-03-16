package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/OrbKonSec/mailintel/cmd/utils"
)

func Hunter(email string) {
	defer utils.ProgressBar.Add(10)
	key := utils.GetAPIKey("Hunter")
	if key == "" {
		return
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", utils.HunterAPI+strings.Split(email, "@")[1]+"&api_key="+key, nil)
	req.Header.Set("User-Agent", "mosint")
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &utils.Hunter_result)
}