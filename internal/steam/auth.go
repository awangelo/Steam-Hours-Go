package steam

import (
	"log"

	"github.com/Philipp15b/go-steam/v3"
)

func Login(user, pass, authCode string) *steam.Client {
	client := steam.NewClient()
	detais := steam.LogOnDetails{
		Username:      user,
		Password:      pass,
		TwoFactorCode: authCode,
	}
	client.Auth.LogOn(&detais)

	for event := range client.Events() {
		switch e := event.(type) {
		case *steam.LoggedOnEvent:
			log.Println("Logged in!")
			return client
		case *steam.LogOnFailedEvent:
			log.Fatalf("Failed to log in: %v", e.Result)
		}
	}
	return nil
}
