package main

import (
	"io/ioutil"
	"net/http"
 )

func httpRequest(endpoint string,  method string) (string, error) {
	conf, err := read_config()
	if err != nil {
		processError(err)
	}
	client := http.Client{}
	url := conf.BasePath + endpoint
	req , err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}

	key, err := getApiKey()
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", key)	
		
	res , err := client.Do(req)
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