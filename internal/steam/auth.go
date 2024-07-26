package steam

import (
	"log"

	"github.com/Philipp15b/go-steam/v3"
	"github.com/Philipp15b/go-steam/v3/protocol/steamlang"
)

func Login(user, pass, authCode string) *steam.Client {
	detais := steam.LogOnDetails{
		Username:      user,
		Password:      pass,
		TwoFactorCode: authCode,
	}
	client := steam.NewClient()
	//client.Auth.LogOn(&detais)

	client.Connect()
	for event := range client.Events() {
		switch e := event.(type) {
		case *steam.ConnectedEvent:
			log.Println("ConnectedEvent triggered")
			client.Auth.LogOn(&detais)
			client.Social.SetPersonaState(steamlang.EPersonaState_Online)
			return client
		case *steam.LoggedOnEvent:
			log.Println("LoggedOnEvent triggered")
			client.Social.SetPersonaState(steamlang.EPersonaState_Online)
		case steam.FatalErrorEvent:
			log.Print(e)
		default:
			log.Print(e)
		}
	}
	return nil
}
