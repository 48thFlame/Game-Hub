package commands

import (
	"github.com/avitar64/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

func Ping(s *dg.Session, i *dg.InteractionCreate) {
	err := discord.InteractionRespond(s, i.Interaction, discord.InstaMessage, &dg.InteractionResponseData{Content: "Pong!"})

	if err != nil {
		discord.Error(err)
	}
}
