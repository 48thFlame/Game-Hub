package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/avitar64/Boost-bot/discord"
	"github.com/avitar64/Boost-bot/discord/commands"
)

const (
	pyInterpreter  = "python3.10"
	pyCommandsFile = "./discord/boost.py"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var err error
	var bot *discord.Bot

	bot, err = discord.NewBot("./discord/TOKEN.txt", pyInterpreter, pyCommandsFile)
	if err != nil {
		log.Fatalf("Error creating bot: %v\n", err)
	}

	commands := commands.ExportCommands()
	for name, handler := range commands {
		bot.AddCommandHandler(name, handler)
	}

	err = bot.S.Open()
	if err != nil {
		log.Fatalf("Error opening bot session: %v\n", err)
	}
	defer bot.S.Close()

	// Wait for a quit signal to quit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down...")
}
