package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRequest(url string) (*http.Response, error) {
	respon, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if respon.StatusCode != http.StatusOK {
		respon.Body.Close()
		return nil, fmt.Errorf("can't get %s, %s", url, respon.Status)
	}

	return respon, nil
}

func JsonDecoding(responce *http.Response) (*Issue, error) {
	var result Issue
	if err := json.NewDecoder(responce.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func StartGetRequestWithJsonResponce(url string) (*Issue, error) {
	responce, err := GetRequest(url)
	if err != nil {
		return nil, err
	}

	json, jErr := JsonDecoding(responce)
	if err != nil {
		return nil, jErr
	}
	responce.Body.Close()
	return json, nil
}
