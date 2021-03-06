package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"net/http"
	"nightbot-go/pkg/constants"
	"nightbot-go/pkg/store"
	"nightbot-go/pkg/types"
	"nightbot-go/pkg/util"
)

func StoreCredentialsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("skipCredentials") != "1" {
		var req types.Credentials
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "failed due to marshaling error: " + err.Error(), 500)
			return
		}

		store.DataStore.Set(store.ClientID, req.ClientID, cache.NoExpiration)
		store.DataStore.Set(store.ClientSecret, req.ClientSecret, cache.NoExpiration)
	}

	clientId, _ := store.DataStore.Get(store.ClientID)

	nightBotURL := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&state=1",
		constants.AuthCodeURL,
		clientId.(string),
		util.GetRedirectURI(),
		constants.Scopes,
	)

	response := types.RedirectURLResponse{
		RedirectURL: nightBotURL,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, "failed due to marshaling error: " + err.Error(), 500)
		return
	}
}
