package discord

import (
	"log"

	dg "github.com/bwmarrin/discordgo"
)

const (
	InstaMessage    = dg.InteractionResponseChannelMessageWithSource
	WillEditMessage = dg.InteractionResponseDeferredChannelMessageWithSource
)

func Error(err error) {
	log.Printf("!! Error: %v\n", err)
}

func InteractionRespond(s *dg.Session, i *dg.Interaction, t dg.InteractionResponseType, r *dg.InteractionResponseData) error {
	return s.InteractionRespond(
		i,
		&dg.InteractionResponse{
			Type: t,
			Data: r,
		},
	)
}
