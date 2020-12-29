package util

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

var client = &http.Client{}

func GetRedirectURI() string {
	return fmt.Sprintf("http://localhost:%s/token", os.Getenv("PORT"))
}

func GetCloseURI() string {
	return fmt.Sprintf("http://localhost:%s/close", os.Getenv("PORT"))
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
