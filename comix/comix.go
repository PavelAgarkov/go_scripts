package comix

import (
	git "../github"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Comix struct {
	Month string `json:"month"`
	Num   int    `json:"num"`
}

const ROOT_URL = "https://xkcd.com"

func GetAllComix() {

	i := 1
	comix := make(map[string]interface{})
	list := make(map[string]map[string]interface{})
	for {
		url := []string{ROOT_URL, strconv.Itoa(i), "info.0.json"}
		adress := strings.Join(url, "/")
		responce, error := git.GetRequest(adress)

		if responce.StatusCode != http.StatusOK || error != nil || i == 100 {
			break
		}
		body, readErr := ioutil.ReadAll(responce.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		jsonErr := json.Unmarshal(body, &comix)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		list[strconv.Itoa(i)] = comix
		i++
	}

}
