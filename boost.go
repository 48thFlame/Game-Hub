package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/avitar64/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

const (
	pyInterpreter  = "python3.10"
	pyCommandsFile = "./discord/commands.py"
)

func main() {
	var err error
	var bot *discord.Bot

	bot, err = discord.NewBot("./discord/TOKEN.txt", pyInterpreter, pyCommandsFile)
	if err != nil {
		log.Fatalf("Error creating bot: %v\n", err)
	}

	bot.AddCommandHandler(
		"ping",
		func(s *dg.Session, i *dg.InteractionCreate) {
			s.InteractionRespond(
				i.Interaction,
				&dg.InteractionResponse{
					Type: dg.InteractionResponseChannelMessageWithSource,
					Data: &dg.InteractionResponseData{
						Content: "Pong!",
					},
				},
			)
		},
	)

	err = bot.S.Open()
	if err != nil {
		log.Fatalf("Error opening bot session: %v\n", err)
	}
	defer bot.S.Close()

	err = bot.S.UpdateListeningStatus("any given feedback!")
	if err != nil {
		log.Fatalf("Error setting listening status: %v\n", err)
	}

	// Wait for a quit signal to quit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down...")
}
