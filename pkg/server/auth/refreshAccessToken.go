package auth

import (
	"encoding/json"
	"errors"
	"github.com/patrickmn/go-cache"
	"net/http"
	"net/url"
	"nightbot-go/pkg/constants"
	"nightbot-go/pkg/store"
	"nightbot-go/pkg/types"
	"nightbot-go/pkg/util"
)

var client = &http.Client{}

func RefreshAccessToken() error {
	clientId, _ := store.DataStore.Get(store.ClientID)
	clientSecret, _ := store.DataStore.Get(store.ClientSecret)
	refreshToken, _ := store.DataStore.Get(store.RefreshToken)

	values := url.Values{}
	values.Add("client_id", clientId.(string))
	values.Add("client_secret", clientSecret.(string))
	values.Add("grant_type", "refresh_token")
	values.Add("redirect_uri", util.GetRedirectURI())
	values.Add("refresh_token", refreshToken.(string))

	resp, err := client.PostForm(constants.AccessTokenURL, values)

	if err != nil || resp.StatusCode >= 400 {
		return errors.New("failed to refresh access token, please re-authorize application")
	}

	var response types.AccessTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return errors.New("failed to decode access token: " + err.Error())
	}

	store.DataStore.Set(store.AccessToken, response.AccessToken, cache.NoExpiration)
	store.DataStore.Set(store.RefreshToken, response.RefreshToken, cache.NoExpiration)

	return nil
}
