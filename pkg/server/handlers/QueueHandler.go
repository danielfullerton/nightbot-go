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

func QueueHandler(w http.ResponseWriter, r *http.Request) {
	err := auth.RefreshAccessToken()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	accessToken, _ := store.DataStore.Get(store.AccessToken)

	resp, err := util.AttemptToMakeAuthorizedAPICall("GET", constants.GetQueueURL, accessToken.(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if resp.StatusCode >= 400 {
		refreshErr := auth.RefreshAccessToken()
		if refreshErr != nil {
			http.Error(w, refreshErr.Error(), 400)
			return
		}

		resp, err = util.AttemptToMakeAuthorizedAPICall("GET", constants.GetQueueURL, accessToken.(string))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var queueResponse types.QueueResponse
	marshalErr := json.Unmarshal(body, &queueResponse)
	if marshalErr != nil {
		http.Error(w, marshalErr.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(&queueResponse)
	if encodeErr != nil {
		http.Error(w, encodeErr.Error(), 500)
		return
	}
}
