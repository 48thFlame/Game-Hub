package commands

import (
	"fmt"
	"os"

	"github.com/48thFlame/Game-hub/boost/discord"
	"github.com/48thFlame/Game-hub/boost/discord/data"
	"github.com/48thFlame/Game-hub/games"
	dg "github.com/bwmarrin/discordgo"
)

func getMasterColorFromMastermindCommandOptionsStringValue(c string) games.MasterColor {
	switch c {
	case "🟥 - red":
		return games.Red
	case "🟧 - orange":
		return games.Orange
	case "🟨 - yellow":
		return games.Yellow
	case "🟩 - green":
		return games.Green
	case "🟦 - blue":
		return games.Blue
	case "🟪 - purple":
		return games.Purple
	default:
		return games.Empty
	}
}

func Mastermind(s *dg.Session, i *dg.InteractionCreate) {
	subCommands := i.ApplicationCommandData().Options
	cmdName := subCommands[0].Name

	if cmdName == "rules" {
		return // responding in python ;)
	}

	// init vars
	var err error
	var needsToSave bool
	var won, lost bool
	game := &games.MastermindGame{}
	id := discord.GetInteractionUser(i.Interaction).ID
	hasGame := data.DataExists(data.GetMastermindFileName(id))
	interaction := i.Interaction

	if cmdName == "new-game" {
		if hasGame {
			err = data.DeleteData(data.GetMastermindFileName(id))
			if err != nil {
				discord.Error(fmt.Errorf("error deleting mastermind game: %v", err), s, interaction)
			}
			user, err := data.LoadUser(id)
			if err != nil {
				discord.Error(fmt.Errorf("error loading user: %v", err), s, interaction)
				return
			}
			user.Stats.Mastermind.Losses++
			err = data.SaveData(data.GetUserFileName(id), user)
			if err != nil {
				discord.Error(fmt.Errorf("error saving user: %v", err), s, interaction)
			}
		}
		discord.InteractionRespond(
			s,
			i.Interaction,
			discord.InstaMessage,
			&dg.InteractionResponseData{
				Content: "New mastermind game successfully started!",
			},
		)
		return
	}

	// load game data or create new game
	if hasGame {
		err = data.LoadData(data.GetMastermindFileName(id), game)
		if err != nil {
			discord.Error(fmt.Errorf("error loading mastermind game: %v", err), s, interaction)
		}
	} else {
		game = games.NewMastermindGame()
		needsToSave = true
	}
	game.FillResults()

	// if guessed should guess
	if cmdName == "guess" {
		options := subCommands[0].Options
		guess := [4]games.MasterColor{}

		for cI, c := range options {
			cStr := c.StringValue()
			guess[cI] = getMasterColorFromMastermindCommandOptionsStringValue(cStr)
		}
		won = game.Guess(guess)
		lost = len(game.Guesses) == games.MasterGameLen && !won
		needsToSave = !won && !lost // only if didn't win or lose should save otherwise will show the user now that they won no keeping data necessary
	}

	// only if updated game data should write to disk
	if needsToSave {
		err = data.SaveData(data.GetMastermindFileName(id), game)
		if err != nil {
			discord.Error(fmt.Errorf("error saving mastermind game: %v", err), s, interaction)
		}
	}

	embed := discord.NewEmbed().
		SetupEmbed().
		SetAuthor("attachment://mastermind.png", "Mastermind", "").
		SetDescription(game.String())
	if won {
		embed.SetTitle("Congratulations! 🥳 You won!")
	} else if lost {
		embed.SetTitle(fmt.Sprintf("You lost! 😢\nThe answer was: %v", game.GetAnswerString(" ")))
	}

	// should delete data if won or lost
	if won || lost {
		// should delete the mastermind game
		data.DeleteData(data.GetMastermindFileName(id))

		// should update the user stats
		user, err := data.LoadUser(id)
		if err != nil {
			discord.Error(fmt.Errorf("error loading user: %v", err), s, interaction)
		}
		if won {
			user.Stats.Mastermind.Wins++
			user.Stats.Mastermind.Rounds += len(game.Guesses)
		} else {
			user.Stats.Mastermind.Losses++
		}

		// should save the user
		err = data.SaveData(data.GetUserFileName(id), user)
		if err != nil {
			discord.Error(fmt.Errorf("error saving user: %v", err), s, interaction)
		}
	}

	boostR, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err), s ,interaction)
	}

	mastermindR, err := os.Open("./discord/assets/mastermind.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening mastermind.png: %v", err), s, interaction)
	}

	err = discord.InteractionRespond(
		s,
		i.Interaction,
		discord.InstaMessage,
		&dg.InteractionResponseData{
			Embeds: []*dg.MessageEmbed{embed.MessageEmbed},
			Files:  []*dg.File{{Name: "boost.png", Reader: boostR}, {Name: "mastermind.png", Reader: mastermindR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to mastermind command interaction: %v", err), s ,interaction)
	}
}
