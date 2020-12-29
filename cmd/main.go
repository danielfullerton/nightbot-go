package main

import (
	"nightbot-go/pkg/server"
	"nightbot-go/pkg/server/auth"
)

func main() {
	auth.StoreEnvTokensIfPresent()
	server.Start()
}
