package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

// func GetDestinyManifest() (*ManifestStruct, error) {
// 	res, err := httpRequest("/Destiny2/Manifest/", "GET")
// 	if err != nil {
// 		return nil, err
// 	}

// 	// // read our opened json as a byte array.
// 	// byteValue, err := ioutil.ReadAll([]byte(res))
// 	// if err != nil {
// 	//     return nil, err
// 	// we initialize our auth struct
// 	var mani ManifestStruct

// 	// we unmarshal our byteArray which contains our
// 	// jsonFile's content into 'auth' which we defined above
// 	json.Unmarshal([]byte(res), &mani)
// 	return &mani, nil
// }

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

	testthing := "/test/%s/string/"
	epoint := fmt.Sprintf(testthing, "the")
	fmt.Println(epoint)
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
