package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"nightbot-go/pkg/constants"
	"nightbot-go/pkg/util"
	"os"
)

func ManualAuthHandler(w http.ResponseWriter, r *http.Request) {
	clientID := os.Getenv("CLIENT_ID")
	nightBotURL := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&state=1",
		constants.AuthCodeURL,
		clientID,
		util.GetRedirectURI(),
		url.QueryEscape(constants.Scopes),
	)
	http.Redirect(w, r, nightBotURL, 302)
}
