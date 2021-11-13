package main

import (
	"github.com/tkanos/gonfig"
)

func read_config() (*Configuration, error) {
	configuration := Configuration{}
	err := gonfig.GetConf("bungieAPI_config.json", &configuration)
	if err != nil {
        return nil, err
    }
	return &configuration, nil
}