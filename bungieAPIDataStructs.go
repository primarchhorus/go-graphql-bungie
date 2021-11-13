package main

import "time"

type Applications struct {
	applications []Application
}
	
type Application struct {
	ApplicationID  int       `json:"applicationId"`
	Name           string    `json:"name"`
	RedirectURL    string    `json:"redirectUrl"`
	Link           string    `json:"link"`
	Scope          string    `json:"scope"`
	Origin         string    `json:"origin"`
	Status         int       `json:"status"`
	CreationDate   time.Time `json:"creationDate"`
	StatusChanged  time.Time `json:"statusChanged"`
	FirstPublished time.Time `json:"firstPublished"`
	Team           []struct {
		Role           int `json:"role"`
		APIEulaVersion int `json:"apiEulaVersion"`
		User           struct {
			SupplementalDisplayName string `json:"supplementalDisplayName"`
			IconPath                string `json:"iconPath"`
			CrossSaveOverride       int    `json:"crossSaveOverride"`
			IsPublic                bool   `json:"isPublic"`
			MembershipType          int    `json:"membershipType"`
			MembershipID            string `json:"membershipId"`
			DisplayName             string `json:"displayName"`
		} `json:"user"`
	} `json:"team"`
}

type UserNameSearchResults struct {
	SearchResults []struct {
		BungieGlobalDisplayName     string `json:"bungieGlobalDisplayName"`
		BungieGlobalDisplayNameCode int    `json:"bungieGlobalDisplayNameCode"`
		BungieNetMembershipID       string `json:"bungieNetMembershipId"`
		DestinyMemberships          []struct {
			IconPath                    string `json:"iconPath"`
			CrossSaveOverride           int    `json:"crossSaveOverride"`
			ApplicableMembershipTypes   []int  `json:"applicableMembershipTypes"`
			IsPublic                    bool   `json:"isPublic"`
			MembershipType              int    `json:"membershipType"`
			MembershipID                string `json:"membershipId"`
			DisplayName                 string `json:"displayName"`
			BungieGlobalDisplayName     string `json:"bungieGlobalDisplayName"`
			BungieGlobalDisplayNameCode int    `json:"bungieGlobalDisplayNameCode"`
		} `json:"destinyMemberships"`
	} `json:"searchResults"`
}