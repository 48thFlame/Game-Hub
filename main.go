package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token, err := os.ReadFile("TOKEN.txt")
	if err != nil {
		log.Fatalf("Could not read TOKEN.txt:\n%v\n", err)
	}

	discord, err := discordgo.New("Bot " + string(token))
	if err != nil {
		log.Fatalf("Error creating Discord session:\n%v\n", err)
	}

	err = discord.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session:\n%v\n", err)
	}
	log.Printf("%v is now online!\n", discord.State.User)

	
	defer discord.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Shutting down...")
}
