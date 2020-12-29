package auth

import (
	"github.com/patrickmn/go-cache"
	"nightbot-go/pkg/store"
	"os"
)

func StoreEnvTokensIfPresent() {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientID != "" && clientSecret != "" {
		store.DataStore.Set(store.ClientID, clientID, cache.NoExpiration)
		store.DataStore.Set(store.ClientSecret, clientSecret, cache.NoExpiration)
	}
}
