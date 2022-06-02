package commands

import (
	"github.com/avitar64/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

type exportedCommands map[string]discord.SlashCommandHandlerType

func emptySlashCommandHandler(*dg.Session, *dg.InteractionCreate) {}

func ExportCommands() exportedCommands {
	ec := make(exportedCommands)

	ec["ping"] = Ping
	ec["8ball"] = Ball8
	ec["dice"] = Dice
	ec["coinflip"] = Coinflip
	ec["poll"] = Poll
	ec["info"] = Info

	ec["calculator"] = emptySlashCommandHandler

	return ec
}
