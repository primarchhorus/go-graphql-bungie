package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
 )

type Configuration struct {
	XAPIKey string `json:"x-api-key"`
	BasePath string `json:"base-path"`
}

func test_oauth() (string, error) {
	
	// bearer := auth.ClientID + " " + auth.ClientToken
	client := http.Client{}
	req , err := http.NewRequest("GET", "https://www.bungie.net/Platform//Destiny2/Manifest/", nil)
	if err != nil {
		return "", err
	}

	// req.Header.Set("client_id", "37774")
	// req.Header.Set("response_type", "code")

	res , err := client.Do(req)
	if err != nil {
		return "", err
	}
	body1, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	sb1:= string(body1)
	log.Printf(sb1)
	return sb1, nil
}

func processError(err error) {
    fmt.Println(err)
    os.Exit(2)
}

func GetDestinyManifest() (*ManifestStruct, error) {
	res, err := httpRequest("/Destiny2/Manifest/", "GET")
	if err != nil {
		return nil, err
	}

    // // read our opened json as a byte array.
    // byteValue, err := ioutil.ReadAll([]byte(res))
	// if err != nil {
    //     return nil, err
    // we initialize our auth struct
    var mani ManifestStruct

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'auth' which we defined above
    json.Unmarshal([]byte(res), &mani)
	return &mani, nil
}

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

func getFirstPartyAppList() (interface{}, error) {
	firstparty, err := httpRequest("/App/FirstParty/", "GET")
	if err != nil {
		processError(err)
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
		processError(err)
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
	url := "/User/Search/Prefix/" + p.Args["username"].(string) + "/0"
	fmt.Println(url)
	firstparty, err := httpRequest(url, "GET")
	if err != nil {
		processError(err)
	}
	var dat map[string]interface{}
	json.Unmarshal([]byte(firstparty), &dat)

	jsonStr, err := json.Marshal(dat["Response"])
	fmt.Println(firstparty)
	if err != nil {
		fmt.Println(err)
	}
	// var search map[string]interface{}
	searchResults := UserNameSearchResults{}
	if err := json.Unmarshal(jsonStr, &searchResults); err != nil {
		fmt.Println(err)
	}

	return searchResults, nil
}

func main() {

	fields := graphql.Fields{
		"applications": &graphql.Field{
			Type: AppType,
			Args: graphql.FieldConfigArgument{
                "applicationId": &graphql.ArgumentConfig{
                    Type: graphql.Float,
                },
            },
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return getFirstPartyAppById(p)
			},
		},
		"applicationslist": &graphql.Field{
			Type: AppList,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return getFirstPartyAppList()
			},
		},
		"searchuser": &graphql.Field{
			Type: ChracterSearchResults,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
            },
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return searchUserName(p)
			},
		},
	}

	
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
	})
	
	http.Handle("/graphql", handler)
	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))

}