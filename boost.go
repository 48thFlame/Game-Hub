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

	// cmd := exec.Command(pyInterpreter, pyCommandsFile)
	// cmd.Stdout = os.Stdout

	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("Error running python script:\n%v\n", err)
	// }
	// bot, err := discord.NewBot("TOKEN.txt", "755001834418208840")
	// if err != nil {
	// 	log.Fatalf("Error creating Discord session:\n%v\n", err)
	// }

	// bot.S.AddHandler(func(s *dg.Session, r *dg.Ready) {
	// 	log.Printf("%v is now online!\n", bot.S.State.User)
	// })

	// // cmd, err := bot.Session.ApplicationCommandCreate(bot.Session.State.User.ID, "755001834418208840", &dg.ApplicationCommand{Name: "ping", Description: "ping"})
	// // if err != nil {
	// // 	log.Fatalf("Error creating slash-command:\n%v\n", err)
	// // }

	// err = bot.S.Open()
	// if err != nil {
	// 	log.Fatalf("Error opening Discord session:\n%v\n", err)
	// }

	// bot.AddCommand(&dg.ApplicationCommand{Name: "ping", Description: "ping"}, func(s *dg.Session, i *dg.InteractionCreate) {
	// 	s.InteractionRespond(i.Interaction, &dg.InteractionResponse{Type: dg.InteractionResponseChannelMessageWithSource, Data: &dg.InteractionResponseData{Content: "pong"}})
	// })
	// bot.RegisterCommands()

	// err = bot.S.UpdateListeningStatus("any given feedback!")
	// if err != nil {
	// 	log.Fatalf("Error setting listening status:\n%v\n", err)
	// }

	// defer bot.S.Close()

	// // Wait for a quit signal to quit
	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, os.Interrupt)
	// log.Println("Press Ctrl+C to exit")
	// <-stop

	// log.Println("Gracefully  shutting down...")
	// bot.RemoveCommands()
	// // defered thing will happen
}

// func pingCommand(bot *dg.Session, m *dg.MessageCreate) {
// 	if m.Content == "+ping" {
// 		bot.ChannelMessageSend(m.ChannelID, "Pong!")
// 		bot.ChannelMessageSendEmbed(m.ChannelID, &dg.MessageEmbed{Title: "Pong!", Fields: []*dg.MessageEmbedField{{Name: "Ping", Value: "Pong!"}}})
// 	}
// }
