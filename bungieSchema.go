package main

import (
	"github.com/graphql-go/graphql"
)

//
// App Endpoint Schema
//
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
				Type: graphql.DateTime,
			},
			"Team": &graphql.Field{
				Type: graphql.NewList(Team),
			},
		},
	},
)

var Team = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Team",
		Fields: graphql.Fields{
			"Role": &graphql.Field{
				Type: graphql.Int,
			},
			"APIEulaVersion": &graphql.Field{
				Type: graphql.Int,
			},
			"User": &graphql.Field{
				Type: BungieUser,
			},
		},
	},
)

var BungieUser = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"SupplementalDisplayName": &graphql.Field{
				Type: graphql.Int,
			},
			"IconPath": &graphql.Field{
				Type: graphql.Int,
			},
			"CrossSaveOverride": &graphql.Field{
				Type: graphql.Int,
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipID": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var AppList = graphql.NewList(AppType)

//
// User Endpoint Schema
//
var Memberships = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Memberships",
		Fields: graphql.Fields{
			"IconPath": &graphql.Field{
				Type: graphql.String,
			},
			"CrossSaveOverride": &graphql.Field{
				Type: graphql.Int,
			},
			"ApplicableMembershipTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipID": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var CharacterSearch = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CharacterSearch",
		Fields: graphql.Fields{
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
			"BungieNetMembershipID": &graphql.Field{
				Type: graphql.String,
			},
			"DestinyMemberships": &graphql.Field{
				Type: graphql.NewList(Memberships),
			},
		},
	},
)

var ChracterSearchResults = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ChracterSearchResults",
		Fields: graphql.Fields{
			"SearchResults": &graphql.Field{
				Type: graphql.NewList(CharacterSearch),
			},
		},
	},
)
