package store

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var DataStore = cache.New(5*time.Minute, 10*time.Minute)

const ClientID = "ClientID"
const ClientSecret = "ClientSecret"

const AuthCode = "AuthCode"
const AccessToken = "AccessToken"
const RefreshToken = "RefreshToken"
