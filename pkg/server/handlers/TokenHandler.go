package handlers

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"nightbot-go/pkg/constants"
	"nightbot-go/pkg/store"
	"nightbot-go/pkg/types"
	"nightbot-go/pkg/util"
)

var client = &http.Client{}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")

	clientId, _ := store.DataStore.Get(store.ClientID)
	clientSecret, _ := store.DataStore.Get(store.ClientSecret)

	values := url.Values{}
	values.Add("client_id", clientId.(string))
	values.Add("client_secret", clientSecret.(string))
	values.Add("code", authCode)
	values.Add("grant_type", "authorization_code")
	values.Add("redirect_uri", util.GetRedirectURI())

	response, err := client.PostForm(constants.AccessTokenURL, values)
	if err != nil {
		http.Error(w, "failed to get access token: " + err.Error(), 500)
		return
	}

	dbytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "failed to parse access token: " + err.Error(), 500)
		return
	}

	var accessTokenResponse types.AccessTokenResponse
	err = json.Unmarshal(dbytes, &accessTokenResponse)
	if err != nil {
		http.Error(w, "failed to parse access token: " + err.Error(), 500)
		return
	}

	log.Println(accessTokenResponse.AccessToken)
	log.Println(accessTokenResponse.RefreshToken)

	store.DataStore.Set(store.AccessToken, accessTokenResponse.AccessToken, cache.NoExpiration)
	store.DataStore.Set(store.RefreshToken, accessTokenResponse.RefreshToken, cache.NoExpiration)

	http.Redirect(w, r, util.GetCloseURI(), 302)
}
