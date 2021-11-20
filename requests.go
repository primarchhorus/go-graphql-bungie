package main

import (
	"io/ioutil"
	"net/http"
)

var (
	conf   = read_config()
	client = http.Client{}
	key    = getApiKey()
)

func httpRequest(endpoint string, method string) (string, error) {

	url := conf.BasePath + endpoint
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", key)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body1, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	sb1 := string(body1)

	return sb1, nil
}
