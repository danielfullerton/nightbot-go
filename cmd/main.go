package main

import (
	_ "github.com/joho/godotenv/autoload"
	"nightbot-go/pkg/server"
	"nightbot-go/pkg/server/auth"
	"nightbot-go/pkg/util"
)

func main() {
	util.KillIfEnvironmentVariablesNotSet()
	auth.StoreEnvTokensIfPresent()
	server.Start()
}
