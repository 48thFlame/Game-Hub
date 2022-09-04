package discord

import (
	"log"
	"os"
	"os/exec"

	dg "github.com/bwmarrin/discordgo"
)

type SlashCommandHandlerType func(s *dg.Session, i *dg.InteractionCreate)

func NewBot(tokenFilePath, pyInterpreter, pyCommandsFile string) (*Bot, error) {
	bot := &Bot{}

	token, err := os.ReadFile(tokenFilePath)
	if err != nil {
		return nil, err
	}

	s, err := dg.New("Bot " + string(token))
	if err != nil {
		return nil, err
	}

	cmdsHandlers := make(map[string]SlashCommandHandlerType)

	s.AddHandler(func(s *dg.Session, i *dg.InteractionCreate) {
		if i.Type == dg.InteractionApplicationCommand {
			if cmdHandler, exists := bot.cmdsHandlers[i.ApplicationCommandData().Name]; exists {
				cmdHandler(s, i)
			} else {
				log.Printf("Command '%v' was executed by user, bot no handler was defined.\n", i.ApplicationCommandData().Name)
			}
		}
	})

	s.AddHandler(func(s *dg.Session, r *dg.Ready) {
		log.Printf("%v is now online!\n", bot.S.State.User)
	})

	bot.S = s
	bot.cmdsHandlers = cmdsHandlers
	bot.pyInterpreter = pyInterpreter
	bot.pyCommandsFile = pyCommandsFile

	go func() {
		err := bot.runPyScript()
		if err != nil {
			log.Fatalf("Error while running python script: %v\n", err)
		}
	}()

	return bot, nil
}

type Bot struct {
	S                             *dg.Session
	cmdsHandlers                  map[string]SlashCommandHandlerType
	pyInterpreter, pyCommandsFile string
}

func (b *Bot) runPyScript() (err error) {
	cmd := exec.Command(b.pyInterpreter, b.pyCommandsFile)
	cmd.Stdout = log.Default().Writer()

	err = cmd.Run()

	return
}

func (b *Bot) AddCommandHandler(name string, handler SlashCommandHandlerType) {
	b.cmdsHandlers[name] = handler
}
