package resolvers

import (
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/primarchhorus/go-graphql-bungie/requests"
)

func base_resolver(p graphql.ResolveParams, url string, method string, ret interface{}) (interface{}, error) {

	request, err := requests.HttpRequest(url, method)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var dat map[string]interface{}
	json.Unmarshal([]byte(request), &dat)

	jsonStr, err := json.Marshal(dat["Response"])

	if err != nil {
		log.Println(err)
		return nil, err
	}

	unmarshalled_results := ret
	if err := json.Unmarshal(jsonStr, &unmarshalled_results); err != nil {
		log.Println(err)
		return nil, err
	}

	return unmarshalled_results, nil
}
