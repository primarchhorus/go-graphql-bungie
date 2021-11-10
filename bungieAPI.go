package main

import (
	"io/ioutil"
	"net/http"
	// "bytes"
	// "net/url"
	"fmt"
	// "time"
	"encoding/json"
	"log"
	"os"
	"github.com/tkanos/gonfig"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	// "github.com/graphql-go/handler"
	// "reflect"
	// "github.com/mitchellh/mapstructure"
 )

type Configuration struct {
	XAPIKey string `json:"x-api-key"`
	BasePath string `json:"base-path"`
}

func read_config() (*Configuration, error) {
	configuration := Configuration{}
	err := gonfig.GetConf("bungieAPI_config.json", &configuration)
	if err != nil {
        return nil, err
    }
	return &configuration, nil
}

func http_request(endpoint string,  method string, conf *Configuration) (string, error) {
	client := http.Client{}
	url := conf.BasePath + endpoint
	req , err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", conf.XAPIKey)	
		
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

func GetDestinyManifest(conf *Configuration) (*ManifestStruct, error) {
	res, err := http_request("/Destiny2/Manifest/", "GET", conf)
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
// var Team = graphql.NewObject( 
// 	graphql.ObjectConfig{
// 	}
// )

var AppType = graphql.NewObject( 
	graphql.ObjectConfig{
		Name: "Application",
		Fields: graphql.Fields{
			"ApplicationID": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"RedirectURL": &graphql.Field{
				Type: graphql.String,
			},
			"Link": &graphql.Field{
				Type: graphql.String,
			},
			"Scope": &graphql.Field{
				Type: graphql.String,
			},
			"Origin": &graphql.Field{
				Type: graphql.String,
			},
			"Status": &graphql.Field{
				Type: graphql.Int,
			},
			"CreationDate": &graphql.Field{
				Type: graphql.DateTime,
			},
			"StatusChanged": &graphql.Field{
				Type: graphql.Int,
			},
			"FirstPublished": &graphql.Field{
				Type: graphql.Int,
			},
		},
    },
)

func graphqlquery(w http.ResponseWriter, req *http.Request){
	
	
	
}

func main() {

	conf, err := read_config()
	if err != nil {
		processError(err)
	}

	// res, err := GetDestinyManifest(conf)
	// if err != nil {
	// 	processError(err)
	// }
	// log.Printf(res.Response.Version)

	fields := graphql.Fields{
		"applications": &graphql.Field{
			Type: AppType,
			Args: graphql.FieldConfigArgument{
                "applicationId": &graphql.ArgumentConfig{
                    Type: graphql.Float,
                },
            },
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				firstparty, err := http_request("/App/FirstParty/", "GET", conf)
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
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	// query := `
	// 	{
	// 		applications(applicationId:29068) {
	// 			CreationDate
	// 		}
	// 	}
	// `

	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
	})
	http.Handle("/graphql", handler)
	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))
	// params := graphql.Params{Schema: schema, RequestString: query}
	// r := graphql.Do(params)
	// if len(r.Errors) > 0 {
	// 	log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	// }
	// log.Printf("%s, ", r)

}