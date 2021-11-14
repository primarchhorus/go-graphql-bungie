package main

import (
	"github.com/tkanos/gonfig"
	"os"
)

func read_config() (*Configuration, error) {
	configuration := Configuration{}
	err := gonfig.GetConf("bungieAPI_config.json", &configuration)
	if err != nil {
        return nil, err
    }
	return &configuration, nil
}

func getApiKey() (string, error) {
	xApiKey := os.Getenv("X_API_KEY") 
	return xApiKey, nil
}