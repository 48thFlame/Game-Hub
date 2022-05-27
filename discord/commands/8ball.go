package commands

import (
	"math/rand"
	"os"

	"github.com/avitar64/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

var ball8Answers = [...]string{
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes, definitely",
	"You may rely on it",
	"As I see it, yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

func Ball8(s *dg.Session, i *dg.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	question := options[0].StringValue()
	answer := ball8Answers[rand.Intn(len(ball8Answers))]

	embed := discord.NewEmbed().
		SetupEmbed().
		SetAuthor("", "🎱 Magic 8-ball", "").
		AddField("Question", question, false).
		AddField("Answer", answer, false).MessageEmbed

	r, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(err)
	}

	err = discord.InteractionRespond(s, i.Interaction, dg.InteractionResponseChannelMessageWithSource, &dg.InteractionResponseData{
		Embeds: []*dg.MessageEmbed{embed},
		Files:  []*dg.File{{Name: "boost.png", Reader: r}},
	})

	if err != nil {
		discord.Error(err)
	}
}
