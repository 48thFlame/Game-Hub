package commands

import (
	"fmt"
	"os"

	"github.com/avitar64/Boost-bot/discord"
	"github.com/avitar64/Boost-bot/discord/data"
	"github.com/avitar64/Boost-bot/games"
	dg "github.com/bwmarrin/discordgo"
)

// var game games.MastermindGame
// var loadedGame bool

func Mastermind(s *dg.Session, i *dg.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	switch options[0].Name {
	case "rules":
		// responding in python ;)
	case "main":
		var err error
		game := &games.MastermindGame{}
		id := discord.GetInteractionUser(i.Interaction).ID

		hasGame := data.DataExists(data.GetMastermindFileName(id))
		if hasGame {
			err = data.LoadData(data.GetMastermindFileName(id), game)
			if err != nil {
				discord.Error(fmt.Errorf("error loading mastermind game: %v", err))
			}
		} else {
			game = games.NewMastermindGame()
			err = data.SaveData(data.GetMastermindFileName(id), game)
			if err != nil {
				discord.Error(fmt.Errorf("error saving mastermind game: %v", err))
			}
		}

		embed := discord.NewEmbed().
			SetupEmbed().
			SetDescription(game.String()).
			SetAuthor("attachment://mastermind.png", "Mastermind", "").MessageEmbed

		boostR, err := os.Open("./discord/assets/boost.png")
		if err != nil {
			discord.Error(fmt.Errorf("error opening boost.png: %v", err))
		}

		mastermindR, err := os.Open("./discord/assets/mastermind.png")
		if err != nil {
			discord.Error(fmt.Errorf("error opening mastermind.png: %v", err))
		}

		err = discord.InteractionRespond(
			s,
			i.Interaction,
			discord.InstaMessage,
			&dg.InteractionResponseData{
				Embeds: []*dg.MessageEmbed{embed},
				Files:  []*dg.File{{Name: "boost.png", Reader: boostR}, {Name: "mastermind.png", Reader: mastermindR}},
			},
		)
		if err != nil {
			discord.Error(fmt.Errorf("error responding to mastermind command interaction: %v", err))
		}
	}
}
