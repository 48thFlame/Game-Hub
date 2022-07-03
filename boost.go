package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/avitar64/Boost-bot/discord"
	"github.com/avitar64/Boost-bot/discord/commands"
)

const (
	pyFilePath                = "./discord/boost.py"
	pyInterpreterNameFilePath = "./pyInterpreterName.txt"
)

var pyInterpreter string

func main() {
	rand.Seed(time.Now().UnixNano())
	var err error

	b, err := ioutil.ReadFile(pyInterpreterNameFilePath)
	if err != nil {
		log.Fatalf("error opening pyInterpreterName.txt: %v", err)
	}
	pyInterpreter = string(b)

	var bot *discord.Bot

	bot, err = discord.NewBot("./discord/TOKEN.txt", pyInterpreter, pyFilePath)
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
