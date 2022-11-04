package commands

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/48thFlame/Game-hub/boost/discord"
	"github.com/48thFlame/Game-hub/boost/discord/data"
	dg "github.com/bwmarrin/discordgo"
)

func Ping(s *dg.Session, i *dg.InteractionCreate) {
	err := discord.InteractionRespond(s, i.Interaction, discord.InstaMessage, &dg.InteractionResponseData{Content: fmt.Sprintf("Pong! %v", s.HeartbeatLatency())})

	if err != nil {
		discord.Error(fmt.Errorf("error responding to ping command interaction: %v", err), s, i.Interaction)
		return
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
		discord.Error(fmt.Errorf("error opening boost.png: %v", err), s, i.Interaction)
		return
	}

	err = discord.InteractionRespond(s, i.Interaction, dg.InteractionResponseChannelMessageWithSource, &dg.InteractionResponseData{
		Embeds: []*dg.MessageEmbed{embed},
		Files:  []*dg.File{{Name: "boost.png", Reader: r}},
	})

	if err != nil {
		discord.Error(fmt.Errorf("error responding to ball8 command interaction: %v", err), s, i.Interaction)
		return
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
		discord.Error(fmt.Errorf("error opening dice.png: %v", err), s, i.Interaction)
		return
	}

	boostR, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err), s, i.Interaction)
		return
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
		discord.Error(fmt.Errorf("error responding to dice command interaction: %v", err), s, i.Interaction)
		return
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
		discord.Error(fmt.Errorf("error opening coinflip.png: %v", err), s, i.Interaction)
		return
	}

	boostR, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err), s, i.Interaction)
		return
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
		discord.Error(fmt.Errorf("error responding to coinflip command interaction: %v", err), s, i.Interaction)
		return
	}
}

const (
	botInviteLink      = "https://bit.ly/3aykHRP"
	boostDiscordServer = "https://discord.gg/ZR2EspdHFQ"
)

func Info(s *dg.Session, i *dg.InteractionCreate) {
	embed := discord.NewEmbed().
		SetupEmbed().
		SetAuthor("attachment://info.png", "Info", "").
		AddField("Bot invite link:", botInviteLink, false).
		AddField("Boost Discord server:", boostDiscordServer, false).
		AddField("Icons by:", "https://icons8.com", false).MessageEmbed

	infoR, err := os.Open("./discord/assets/info.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening info.png: %v", err), s, i.Interaction)
		return
	}

	boostR, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err), s, i.Interaction)
		return
	}

	err = discord.InteractionRespond(
		s,
		i.Interaction,
		dg.InteractionResponseChannelMessageWithSource,
		&dg.InteractionResponseData{
			Embeds: []*dg.MessageEmbed{embed},
			Files:  []*dg.File{{Name: "info.png", Reader: infoR}, {Name: "boost.png", Reader: boostR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to info command interaction: %v", err), s, i.Interaction)
		return
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
		AddField(fmt.Sprintf("%v's poll:", discord.GetInteractionUser(i.Interaction)), poll, false).MessageEmbed

	pollR, err := os.Open("./discord/assets/poll.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening poll.png: %v", err), s, i.Interaction)
		return
	}

	boostR, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err), s, i.Interaction)
		return
	}

	err = discord.InteractionRespond(s, i.Interaction, discord.DefferSendMessage, nil)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to poll command interaction with deffer thing: %v", err), s, i.Interaction)
		return
	}

	msg, err := s.ChannelMessageSendComplex(
		channel.ID,
		&dg.MessageSend{
			Embed: embed,
			Files: []*dg.File{{Name: "poll.png", Reader: pollR}, {Name: "boost.png", Reader: boostR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error sending poll: %v", err), s, i.Interaction)
		return
	}

	s.MessageReactionAdd(msg.ChannelID, msg.ID, pollReactionPositive)
	s.MessageReactionAdd(msg.ChannelID, msg.ID, pollReactionNegative)

	content := "Done!"
	err = discord.InteractionEdit(s, i.Interaction, &dg.WebhookEdit{Content: &content})
	if err != nil {
		discord.Error(fmt.Errorf("error editing poll interaction: %v", err), s, i.Interaction)
		return
	}
}

const feedbackChannelId = "845017352264351764"

func Feedback(s *dg.Session, i *dg.InteractionCreate) {
	var err error
	options := i.ApplicationCommandData().Options
	interactionUser := discord.GetInteractionUser(i.Interaction)

	userData, err := data.LoadUser(interactionUser.ID)
	if err != nil {
		discord.Error(fmt.Errorf("error loading user data: %v", err), s, i.Interaction)
		return
	}

	if userData.Feedback {
		err = discord.InteractionRespond(
			s,
			i.Interaction,
			discord.InstaMessage,
			&dg.InteractionResponseData{Content: "You are banned from using the feedback command!"},
		)
		if err != nil {
			discord.Error(fmt.Errorf("error responding to feedback command interaction: %v", err), s, i.Interaction)
			return
		}

		return
	}

	err = discord.InteractionRespond(
		s,
		i.Interaction,
		discord.DefferSendMessage,
		nil,
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to feedback command interaction: %v", err), s, i.Interaction)
		return
	}

	embed := discord.NewEmbed().
		SetAuthor("", "Feedback", "").
		SetTitle(fmt.Sprintf("%v, name at time: %v", interactionUser.ID, interactionUser.String())).
		SetDescription(options[0].StringValue()).MessageEmbed

	_, err = s.ChannelMessageSendEmbed(
		feedbackChannelId,
		embed,
	)
	if err != nil {
		discord.Error(fmt.Errorf("error sending feedback: %v", err), s, i.Interaction)
		return
	}

	content := "Thank you for your feedback! 😎\n\u200b\n||**WARNING:** DO NOT SPAM! spamming will lead to a ban from using the feedback command!||"
	err = discord.InteractionEdit(
		s,
		i.Interaction,
		&dg.WebhookEdit{
			Content: &content,
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to feedback command interaction with deffer thing: %v", err), s, i.Interaction)
		return
	}
}

func Statistics(s *dg.Session, i *dg.InteractionCreate) {
	id := discord.GetInteractionUser(i.Interaction).ID
	userData, err := data.LoadUser(id)
	if err != nil {
		discord.Error(fmt.Errorf("error loading user data: %v", err), s, i.Interaction)
		return
	}

	wins, losses, rounds := userData.Stats.Mastermind.Wins, userData.Stats.Mastermind.Losses, userData.Stats.Mastermind.Rounds
	totalGames := wins + losses

	mastermindStatsStr := ""
	mastermindStatsStr += fmt.Sprintf("> Wins: %v\n", wins)
	mastermindStatsStr += fmt.Sprintf("> Losses: %v\n", losses)
	mastermindStatsStr += fmt.Sprintf("> Total games: %v\n", totalGames)
	mastermindStatsStr += fmt.Sprintf("> Average rounds to victory: %v\n", float64(rounds)/float64(wins))

	embed := discord.NewEmbed().
		SetupEmbed().
		SetAuthor("attachment://stats.png", "Statistics", "").
		AddField("Mastermind:", mastermindStatsStr, false).MessageEmbed

	statsR, err := os.Open("./discord/assets/stats.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening stats.png: %v", err), s, i.Interaction)
		return
	}

	boostR, err := os.Open("./discord/assets/boost.png")
	if err != nil {
		discord.Error(fmt.Errorf("error opening boost.png: %v", err), s, i.Interaction)
	}

	err = discord.InteractionRespond(
		s,
		i.Interaction,
		discord.InstaMessage,
		&dg.InteractionResponseData{
			Embeds: []*dg.MessageEmbed{embed},
			Files:  []*dg.File{{Name: "stats.png", Reader: statsR}, {Name: "boost.png", Reader: boostR}},
		},
	)
	if err != nil {
		discord.Error(fmt.Errorf("error responding to statistics command interaction: %v", err), s, i.Interaction)
	}
}
