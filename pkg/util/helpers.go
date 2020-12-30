package util

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

var client = &http.Client{}
var baseURL = os.Getenv("BASE_URL")
var port = os.Getenv("PORT")

func GetRedirectURI() string {
	return fmt.Sprintf("%s:%s/api/config/token", baseURL, port)
}

func GetCloseURI() string {
	return fmt.Sprintf("%s:%s/close", baseURL, port)
}

func AttemptToMakeAuthorizedAPICall(method, apiUrl, accessToken string) (*http.Response, error) {
	reqUrl, err := url.Parse(apiUrl)
	if err != nil {
		return &http.Response{}, err
	}

	request := &http.Request{
		Method: method,
		URL: reqUrl,
		Header: map[string][]string {
			"Content-Type": {"application/json; charset=UTF-8"},
			"Authorization": {fmt.Sprintf("Bearer %s", accessToken)},
		},
	}

	return client.Do(request)
}

func KillIfEnvironmentVariablesNotSet() {
	if
		os.Getenv("BASE_URL") == "" ||
		os.Getenv("PORT") == "" {
		log.Fatal("Please provide at least the following environment variables: BASE_URL, PORT")
	}
}
