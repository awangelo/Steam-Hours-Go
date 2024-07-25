package main

import (
	"flag"
	"log"
	"os"

	"github.com/awangelo/Steam-Hours-Go/internal/steam"
	"github.com/awangelo/Steam-Hours-Go/pkg/logger"
)

func main() {
	var (
		user = flag.String("username", "", "Steam username")
		pass = flag.String("password", "", "Steam password")
	)

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Println("Usage: steam-go <app-id>")
		os.Exit(1)
	}
	appId := args[0]

	logger.Init()

	client := steam.Login(*user, *pass)
	steam.StartFarm(client, appId)
	// Chamar a func em auth (pass, user e talvez auth code) que vai
	// retornar um *Client que sera usado em StartHours(mudar depois?).
	// por enquanto, as logs continuaram assim no stdout

	// Adcionar logs do go-steam

	select {} // Manter
}
