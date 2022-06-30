package commands

import (
	"github.com/avitar64/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

type exportedCommands map[string]discord.SlashCommandHandlerType

func emptySlashCommandHandler(*dg.Session, *dg.InteractionCreate) {}

func ExportCommands() exportedCommands {
	ec := make(exportedCommands)

	// non games
	ec["ping"] = Ping
	ec["8ball"] = Ball8
	ec["dice"] = Dice
	ec["coinflip"] = Coinflip
	ec["poll"] = Poll
	ec["info"] = Info
	ec["feedback"] = Feedback

	// games
	ec["mastermind"] = Mastermind

	// python-executed commands
	ec["calculator"] = emptySlashCommandHandler
	ec["help"] = emptySlashCommandHandler

	return ec
}
