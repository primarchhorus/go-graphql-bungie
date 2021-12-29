package main

import (
	"github.com/graphql-go/graphql"
)

var (
	fields = graphql.Fields{
		"applications": &graphql.Field{
			Type: AppType,
			Args: graphql.FieldConfigArgument{
				"applicationId": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return appById(p)
			},
		},
		"applicationslist": &graphql.Field{
			Type: AppList,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return appList()
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
				return searchUser(p)
			},
		},
	}
)
