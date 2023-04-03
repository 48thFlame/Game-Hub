package main

import (
	"fmt"
	"strings"

	shell "github.com/48thFlame/Command-Shell"

	"github.com/48thFlame/Game-Hub/games"
)

const mastermindCommandDataName = "mastermind"

func mastermindCommandHandler(i *shell.CommandInput) error {
	guessed := false

	argsNum := len(i.Args)
	if argsNum > 0 {
		if argsNum == 4 {
			guessed = true
		} else {
			return fmt.Errorf(
				"invalid guess arguments number expected exactly 4 found %v, Note: if you don't want to guess pass 0 arguments",
				argsNum,
			)
		}
	}

	if _, ok := i.Cmd.Data[mastermindCommandDataName]; !ok {
		fmt.Fprintln(i.Stdout, "Creating new mastermind game...")
		i.Cmd.Data[mastermindCommandDataName] = games.NewMastermindGame()
	}
	game := i.Cmd.Data[mastermindCommandDataName].(*games.MastermindGame)

	var won bool
	if guessed {
		guess := [4]games.MasterColor{}
		for j := 0; j < 4; j++ {
			guess[j] = games.ConvertLetterToColor(i.Args[j])
		}
		won = game.Guess(guess)
	}

	fmt.Fprintln(i.Stdout, masterGameToString(game))

	if won {
		fmt.Fprintln(i.Stdout, "Congratulations!🥳 You won!")
		delete(i.Cmd.Data, mastermindCommandDataName)
	}

	return nil
}

func masterGameToString(m *games.MastermindGame) (str string) {
	masterSecretEmoji := "❓"
	masterBoardSeparator := " -- "

	str += "Answer:" + "\n"
	str += strings.Repeat(masterSecretEmoji+" ", 4) +
		masterBoardSeparator +
		strings.Repeat(masterResultToString(games.Black)+" ", 4) +
		"\n"

	var guess [4]games.MasterColor
	var result []games.MasterResult

	for i := 0; i < games.MasterGameLen; i++ {
		if len(m.Guesses) > i { //if guessed up until now, use the guess, otherwise use a blank/empty guess
			guess = m.Guesses[i]
			result = m.Results[i]
		} else {
			guess = [4]games.MasterColor{games.Blank, games.Blank, games.Blank, games.Blank}
			result = []games.MasterResult{games.Empty, games.Empty, games.Empty, games.Empty}
		}

		str += "Round " + fmt.Sprint(i+1, ":")
		str += "\n"

		for _, color := range guess {
			str += masterColorToString(color)
			str += " "
		}
		str += masterBoardSeparator
		for _, result := range result {
			str += masterResultToString(result)
			str += " "
		}
		if i != games.MasterGameLen-1 { // if its not the last round then should add new line char
			str += "\n"
		}
	}

	return
}

func masterResultToString(mr games.MasterResult) string {
	switch mr {
	case games.Empty:
		return "🔳"
	case games.White:
		return "❎"
	case games.Black:
		return "✅"
	default:
		return ""
	}
}

func masterColorToString(c games.MasterColor) string {
	switch c {
	case games.Blank:
		return "🔳"
	case games.Red:
		return "🟥"
	case games.Orange:
		return "🟧"
	case games.Yellow:
		return "🟨"
	case games.Green:
		return "🟩"
	case games.Blue:
		return "🟦"
	case games.Purple:
		return "🟪"
	default:
		return ""
	}
}
