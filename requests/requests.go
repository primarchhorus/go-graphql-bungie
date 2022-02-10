package requests

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/primarchhorus/go-graphql-bungie/configuration"
)

var (
	client   = http.Client{}
	key      = getApiKey()
	clientId = os.Getenv("CLIENT_ID")
	// clientSecret = os.Getenv("CLIENT_SECRET")
)

func getApiKey() string {
	xApiKey := os.Getenv("X_API_KEY")
	if len(xApiKey) == 0 {
		err := errors.New("X_API_KEY environment variable empty")
		log.Fatalln(err)
	}
	return xApiKey
}

type error interface {
	Error() string
}

func HttpRequest(endpoint string, method string) (string, error) {

	url := configuration.BasePath + endpoint
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", key)
	req.Header.Set("Authorization", "Bearer "+configuration.ReadTempTokens().AccessToken)

	res, err := client.Do(req)
	if res.StatusCode != 200 {
		log.Println(res.Status)
		return "", errors.New(res.Status)
	}
	if err != nil {
		log.Println(err)
		return "", err
	}
	body1, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	sb1 := string(body1)

	return sb1, nil
}

func oauth_step_one() {
	urls := "https://www.bungie.net/en/oauth/authorize"
	param := url.Values{}
	param.Add("client_id", clientId)
	param.Add("response_type", "code")
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {

	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", key)

	res, err := client.Do(req)
	if err != nil {

	}
	body1, err := ioutil.ReadAll(res.Body)
	if err != nil {

	}
	sb1 := string(body1)
	fmt.Println(sb1)
}
