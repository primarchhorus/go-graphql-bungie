package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

var config = read_config()

func getFirstPartyAppList() (interface{}, error) {
	firstparty, err := httpRequest(config.API.App.GetBungieApplications.Endpoint,
		config.API.App.GetBungieApplications.Method)
	if err != nil {
		log.Println(err)
	}
	var apps []interface{}
	var dat map[string]interface{}
	json.Unmarshal([]byte(firstparty), &dat)

	jsonStr, err := json.Marshal(dat["Response"])
	if err != nil {
		log.Println(err)
	}

	if err := json.Unmarshal(jsonStr, &apps); err != nil {
		log.Println(err)
	}

	var s []interface{}
	for _, app := range apps {
		mappedApp := app.(map[string]interface{})

		jsonString, _ := json.Marshal(mappedApp)
		appid := Application{}
		json.Unmarshal(jsonString, &appid)
		s = append(s, appid)
	}
	return s, nil
}

func getFirstPartyAppById(p graphql.ResolveParams) (interface{}, error) {
	firstparty, err := httpRequest("/App/FirstParty/", "GET")
	if err != nil {
		log.Println(err)
	}
	var apps []interface{}
	var dat map[string]interface{}
	json.Unmarshal([]byte(firstparty), &dat)

	jsonStr, err := json.Marshal(dat["Response"])
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(jsonStr, &apps); err != nil {
		fmt.Println(err)
	}

	id, ok := p.Args["applicationId"].(float64)
	if ok {
		for _, app := range apps {
			mappedApp := app.(map[string]interface{})
			if mappedApp["applicationId"] == id {
				jsonString, _ := json.Marshal(mappedApp)
				appid := Application{}
				json.Unmarshal(jsonString, &appid)
				return appid, nil
			}

		}
	}
	return nil, nil
}

func searchUserName(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.User.SearchByGlobalNamePrefix.Endpoint, p.Args["username"].(string), "/0")
	// url := "/User/Search/Prefix/" + p.Args["username"].(string) + "/0"
	fmt.Println(url)
	firstparty, err := httpRequest(url, config.API.User.SearchByGlobalNamePrefix.Method)
	if err != nil {
		log.Println(err)
	}

	var dat map[string]interface{}
	json.Unmarshal([]byte(firstparty), &dat)

	jsonStr, err := json.Marshal(dat["Response"])

	if err != nil {
		log.Println(err)
	}
	// var search map[string]interface{}
	searchResults := UserNameSearchResults{}
	if err := json.Unmarshal(jsonStr, &searchResults); err != nil {
		log.Println(err)
	}

	return searchResults, nil
}
