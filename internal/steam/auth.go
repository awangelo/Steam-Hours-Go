package steam

import (
	"log"

	"github.com/Philipp15b/go-steam/v3"
)

func Login(user, pass string) *steam.Client {
	client := steam.NewClient()
	detais := steam.LogOnDetails{Username: user, Password: pass}
	client.Auth.LogOn(&detais)

	go clientListen(client)

	for event := range client.Events() {
		switch e := event.(type) {
		case *steam.LoggedOnEvent:
			log.Println("Logged in!")
			return client
		case *steam.LogOnFailedEvent:
			log.Fatalf("Failed to log in: %v", e.Result)
		default:
			log.Printf("Event: %v", e)
		}
	}

	return nil
}

// clientListen logs all events from the new client.
func clientListen(client *steam.Client) {
	for e := range client.Events() {
		log.Printf("Event: %v", e)
	}
}
