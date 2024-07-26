package steam

import (
	"log"
	"strconv"
	"time"

	"github.com/Philipp15b/go-steam/v3"
	"github.com/Philipp15b/go-steam/v3/protocol/steamlang"
)

// func (a *Auth) handleLogOnResponse(packet *protocol.Packet) {

// ...} else if result == steamlang.EResult_Fail || result == steamlang.EResult_ServiceUnavailable || result == steamlang.EResult_TryAnotherCM {
// 	// some error on Steam's side, we'll get an EOF later

// Broken method prob. ↑

// Connect failed: dial tcp 146.66.155.8:27017: connect: connection timed out
// example bot ↑
func StartFarm(client *steam.Client, appId string) {
	log.Printf("Starting farming for AppID: %s", appId)

	client.Social.SetPersonaState(steamlang.EPersonaState_Online)

	game, err := strconv.Atoi(appId)
	if err != nil {
		log.Fatalf("Failed to convert AppID to int: %s", err)
	}

	client.GC.SetGamesPlayed(uint64(game))

	for {
		for event := range client.Events() {
			switch e := event.(type) {
			case error:
				log.Printf("Error received: %v", e)

				log.Println("Connection closed unexpectedly. Reconnecting...")
				time.Sleep(5 * time.Second)            // Wait before retrying
				client.Connect()                       // Reconnect the client
				client.GC.SetGamesPlayed(uint64(game)) // Resume playing the game

			default:
				log.Println(e)
			}
		}
	}
}

// ver https://pkg.go.dev/github.com/Philipp15b/go-steam/v3@v3.0.0/protocol/gamecoordinator#GCMsg

// type Client struct {
// 	Auth          *Auth
// 	Social        *Social
// 	Web           *Web
// 	Notifications *Notifications
// 	Trading       *Trading
// 	GC            *GameCoordinator

// 	ConnectionTimeout time.Duration
// 	// contains filtered or unexported fields
// }
