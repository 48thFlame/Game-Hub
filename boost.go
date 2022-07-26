package main

import (
	"fmt"
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
	err := run()
	if err != nil {
		log.Fatalf("Error running bot, with error:\n%v\n", err)
	}
}

func run() (err error) {
	rand.Seed(time.Now().UnixNano())
	log.Default().SetOutput(os.Stdout)

	b, err := ioutil.ReadFile(pyInterpreterNameFilePath)
	if err != nil {
		return fmt.Errorf("error opening pyInterpreterName.txt: %v", err)
	}
	pyInterpreter = string(b)

	var bot *discord.Bot

	bot, err = discord.NewBot("./discord/TOKEN.txt", pyInterpreter, pyFilePath)
	if err != nil {
		return fmt.Errorf("error creating bot: %v", err)
	}

	commands := commands.ExportCommands()
	for name, handler := range commands {
		bot.AddCommandHandler(name, handler)
	}

	err = bot.S.Open()
	if err != nil {
		return fmt.Errorf("error opening bot session: %v", err)
	}
	defer bot.S.Close()

	// Wait for a quit signal to quit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down...")

	return nil
}
