package util

import (
	"fmt"
	"os"
)

func GetRedirectURI() string {
	return fmt.Sprintf("http://localhost:%s/token", os.Getenv("PORT"))
}

func GetCloseURI() string {
	return fmt.Sprintf("http://localhost:%s/close", os.Getenv("PORT"))
}
