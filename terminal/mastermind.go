package main

import (
	"fmt"

	"github.com/avitar64/Boost-bot/games"
)

const (
	mastermindCommandDataName = "mastermind"
	mastermindGuessName       = "g"
)

func mastermindCommand(t *terminal, args []string) {
	var guessed, won bool

	// should check for errors before any action is done.
	if len(args) > 0 {
		switch args[0] {
		case mastermindGuessName:
			if len(args) < 5 {
				t.Error("not enough arguments to use mastermind guess command.")
				return
			} else if len(args) > 5 {
				t.Error("too many arguments to use mastermind guess command.")
				return
			}
			guessed = true
		}
	}

	if _, ok := t.data[mastermindCommandDataName]; !ok {
		fmt.Println("Creating new mastermind game...")
		t.data[mastermindCommandDataName] = games.NewMastermindGame()
	}
	game := t.data[mastermindCommandDataName].(*games.MastermindGame)

	if guessed {
		guess := [4]games.MasterColor{}
		for i := 0; i < 4; i++ {
			guess[i] = games.ConvertLetterToColor(args[i+1])
		}
		won = game.Guess(guess)
	}
	fmt.Println(game)
	if won {
		fmt.Println("Congratulations!🥳 You won!")
		delete(t.data, mastermindCommandDataName)
	}
}
