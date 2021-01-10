package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"nightbot-go/pkg/constants"
	"nightbot-go/pkg/server/auth"
	"nightbot-go/pkg/store"
	"nightbot-go/pkg/types"
	"nightbot-go/pkg/util"
)

func ChannelHandler(w http.ResponseWriter, r *http.Request) {
	err := auth.RefreshAccessToken()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	accessToken, _ := store.DataStore.Get(store.AccessToken)
	resp, err := util.AttemptToMakeAuthorizedAPICall("GET", constants.GetChannelURL, accessToken.(string))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	if resp.StatusCode >= 400 {
		refreshErr := auth.RefreshAccessToken()
		if refreshErr != nil {
			http.Error(w, refreshErr.Error(), 400)
			return
		}

		resp, err = util.AttemptToMakeAuthorizedAPICall("GET", constants.GetChannelURL, accessToken.(string))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var channelResponse types.ChannelResponse
	marshalErr := json.Unmarshal(body, &channelResponse)
	if marshalErr != nil {
		log.Println(marshalErr)
		http.Error(w, marshalErr.Error(), 500)
		return
	}

	log.Println(channelResponse)
	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(&channelResponse)
	if encodeErr != nil {
		log.Println(encodeErr)
		http.Error(w, encodeErr.Error(), 500)
		return
	}
}
