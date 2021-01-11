package jobs

import (
	"nightbot-go/pkg/server/auth"
	"time"
)

func RefreshAuthHourly() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for {
			select {
			case <- ticker.C: {
				_ = auth.RefreshAccessToken()
			}
			}
		}
	}()
}
