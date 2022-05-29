package commands

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/avitar64/Boost-bot/discord"
	dg "github.com/bwmarrin/discordgo"
)

func Ping(s *dg.Session, i *dg.InteractionCreate) {
	err := discord.InteractionRespond(s, i.Interaction, discord.InstaMessage, &dg.InteractionResponseData{Content: fmt.Sprintf("Pong! %v", s.HeartbeatLatency())})

	if err != nil {
		discord.Error(fmt.Errorf("error responding to ping command interaction: %v", err))
	}
}

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
		SetAuthor("", "🎱 \u200b Magic 8-ball", "").
		AddField("Question", question, false).
		AddField("Answer", answer, false).MessageEmbed

	r, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err))
	}

	err = discord.InteractionRespond(s, i.Interaction, dg.InteractionResponseChannelMessageWithSource, &dg.InteractionResponseData{
		Embeds: []*dg.MessageEmbed{embed},
		Files:  []*dg.File{{Name: "boost.png", Reader: r}},
	})

	if err != nil {
		discord.Error(fmt.Errorf("error responding to ball8 command interaction: %v", err))
	}
}

func Dice(s *dg.Session, i *dg.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	sides := int(options[0].IntValue())
	dice := int(options[1].IntValue())

	rolledDice := make([]string, 0)

	for i := 0; i < dice; i++ {
		rolledDice = append(rolledDice, fmt.Sprint(rand.Intn(sides)+1))
	}

	embed := discord.NewEmbed().
		SetupEmbed().
		SetAuthor("attachment://dice.png", "Dice", "").
		AddField(fmt.Sprintf("You rolled %v dice, each with %v sides", dice, sides), strings.Join(rolledDice, ", "), false).MessageEmbed

	diceR, err := os.Open("./discord/assets/dice.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening dice.png: %v", err))
	}

	boostR, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err))
	}

	err = discord.InteractionRespond(
		s,
		i.Interaction,
		discord.InstaMessage,
		&dg.InteractionResponseData{
			Embeds: []*dg.MessageEmbed{embed},
			Files:  []*dg.File{{Name: "dice.png", Reader: diceR}, {Name: "boost.png", Reader: boostR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to dice command interaction: %v", err))
	}

}

var coinStates = [...]string{"heads", "tails"}

func Coinflip(s *dg.Session, i *dg.InteractionCreate) {
	res := coinStates[rand.Intn(len(coinStates))]

	embed := discord.NewEmbed().
		SetupEmbed().
		SetAuthor("attachment://coinflip.png", "Coinflip", "").
		AddField("Flipping a coin", fmt.Sprintf("It's %v!", res), false).MessageEmbed

	cfr, err := os.Open("./discord/assets/coinflip.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening coinflip.png: %v", err))
	}

	boostR, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err))
	}

	err = discord.InteractionRespond(s,
		i.Interaction,
		dg.InteractionResponseChannelMessageWithSource,
		&dg.InteractionResponseData{
			Embeds: []*dg.MessageEmbed{embed},
			Files:  []*dg.File{{Name: "coinflip.png", Reader: cfr}, {Name: "boost.png", Reader: boostR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to coinflip command interaction: %v", err))
	}
}

const (
	pollReactionPositive = "👍"
	pollReactionNegative = "👎"
)

func Poll(s *dg.Session, i *dg.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	channel := options[0].ChannelValue(nil)
	poll := options[1].StringValue()

	embed := discord.NewEmbed().
		SetupEmbed().
		SetAuthor("attachment://poll.png", "Poll", "").
		AddField(fmt.Sprintf("%v's poll:", i.Member.User), poll, false).MessageEmbed

	pollR, err := os.Open("./discord/assets/poll.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening poll.png: %v", err))
	}

	boostR, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err))
	}

	err = discord.InteractionRespond(s, i.Interaction, discord.DefferSendMessage, nil)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to poll command interaction with deffer thing: %v", err))
	}

	msg, err := s.ChannelMessageSendComplex(
		channel.ID,
		&dg.MessageSend{
			Embed: embed,
			Files: []*dg.File{{Name: "poll.png", Reader: pollR}, {Name: "boost.png", Reader: boostR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error sending poll: %v", err))
	}

	s.MessageReactionAdd(msg.ChannelID, msg.ID, pollReactionPositive)
	s.MessageReactionAdd(msg.ChannelID, msg.ID, pollReactionNegative)

	err = discord.InteractionEdit(s, i.Interaction, &dg.WebhookEdit{Content: "Done!"})
	if err != nil {
		discord.Error(fmt.Errorf("error editing poll interaction: %v", err))
	}
}
