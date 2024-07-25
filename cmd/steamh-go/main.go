package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/awangelo/Steam-Hours-Go/internal/steam"
	"github.com/awangelo/Steam-Hours-Go/pkg/logger"
)

func main() {
	var (
		user string
		pass string
	)

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Println("Usage: steam-go <app-id>")
		os.Exit(1)
	}
	appId := args[0]

	fmt.Print("Steam username: ")
	fmt.Scanln(&user)
	if user == "" {
		log.Println("Username cannot be empty")
		os.Exit(1)
	}

	fmt.Print("Steam password: ")
	fmt.Scanln(&pass)
	if pass == "" {
		log.Println("Password cannot be empty")
		os.Exit(1)
	}

	clearTerm()
	logger.Init()

	client := steam.Login(user, pass)
	steam.StartFarm(client, appId)
	// Chamar a func em auth (pass, user e talvez auth code) que vai
	// retornar um *Client que sera usado em StartHours(mudar depois?).
	// por enquanto, as logs continuaram assim no stdout
}

func clearTerm() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
