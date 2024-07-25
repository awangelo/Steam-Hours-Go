package main

import (
	"flag"
	"log"
	"os"

	"github.com/awangelo/Steam-Hours-Go/pkg/logger"
)

func main() {
	var (
		user = flag.String("username", "", "Steam username")
		pass = flag.String("password", "", "Steam password")
	)

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 || len(args) > 1 {
		log.Println("Usage: steam-go <app-id>")
		os.Exit(1)
	}
	appId := args[0]

	logger.Init()
}
