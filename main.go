package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func newBot(tokenFilePath string) (*discordgo.Session, error) {
	token, err := os.ReadFile(tokenFilePath)
	if err != nil {
		return nil, err
	}

	bot, err := discordgo.New("Bot " + string(token))
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func main() {
	bot, err := newBot("TOKEN.txt")
	if err != nil {
		log.Fatalf("Error creating Discord session:\n%v\n", err)
	}

	err = bot.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session:\n%v\n", err)
	}
	log.Printf("%v is now online!\n", bot.State.User)

	bot.AddHandler(pingCommand)
	err = bot.UpdateListeningStatus("To any given feedback!")
	if err != nil {
		log.Fatalf("Error setting listening status:\n%v\n", err)
	}

	// Wait for a quit signal to quit
	defer bot.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("💤💤💤Gracefully  shutting down...")
}

func pingCommand(bot *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "+ping" {
		bot.ChannelMessageSend(m.ChannelID, "Pong!")
		bot.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{Title: "Pong!", Fields: []*discordgo.MessageEmbedField{{Name: "Ping", Value: "Pong!"}}})
	}
}
