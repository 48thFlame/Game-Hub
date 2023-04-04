package main

import (
	"fmt"
	"strconv"
	"strings"

	shell "github.com/48thFlame/Command-Shell"

	"github.com/48thFlame/Game-Hub/games"
)

const connect4CommandDataGameName = "connect4"

func connect4CommandHandler(i *shell.CommandInput) error {
	placing := false

	argsNum := len(i.Args)
	if argsNum > 0 {
		if argsNum == 1 {
			placing = true
		} else {
			return fmt.Errorf(
				"invalid guess arguments number expected exactly 1 found %v, Note: if you don't want to place pass 0 arguments",
				argsNum,
			)
		}
	}

	if _, ok := i.Cmd.Data[connect4CommandDataGameName]; !ok {
		fmt.Fprintln(i.Stdout, "Creating new connect4 game...")
		i.Cmd.Data[connect4CommandDataGameName] = games.NewConnect4Game()
	}
	game := i.Cmd.Data[connect4CommandDataGameName].(*games.Connect4Game)

	if placing {
		col, err := strconv.Atoi(i.Args[0])
		if err != nil {
			return fmt.Errorf("argument is not a colmun number, convertetd with error: %v", err)
		}

		good := game.Turn(col - 1)
		if !good {
			return fmt.Errorf("colmun %v is full", col)
		}

		if game.GameState == games.CStateDraw {
			fmt.Fprintf(i.Stdout, "%v", connect4GameToString(game))
			fmt.Fprintln(i.Stdout, "The game ended in a draw...")
			delete(i.Cmd.Data, connect4CommandDataGameName)

		} else if game.GameState == games.CStatePlr1Won { // ! ! ! change to be dynamic player can go second
			fmt.Fprintf(i.Stdout, "%v", connect4GameToString(game))
			fmt.Fprintln(i.Stdout, "Congratulations!🥳 You won!")
			delete(i.Cmd.Data, connect4CommandDataGameName)

		} else {
			aiCol := games.Connect4GetAiMove(*game)
			_ = game.Turn(aiCol)
			fmt.Fprintf(i.Stdout, "Ai went at col: %v\n%v", aiCol+1, connect4GameToString(game))

			if game.GameState == games.CStatePlr2Won {
				fmt.Fprintln(i.Stdout, "The ai won, you lost :(")
				delete(i.Cmd.Data, connect4CommandDataGameName)
			} else if game.GameState == games.CStateDraw {
				fmt.Fprintln(i.Stdout, "The game ended in a draw...")
				delete(i.Cmd.Data, connect4CommandDataGameName)
			}
		}
	}

	return nil
}

func connect4GameToString(game *games.Connect4Game) string {
	var sb strings.Builder

	for _, row := range game.Board {
		for _, spot := range row {
			sb.WriteString(cSpotToString(spot))
		}
		sb.WriteRune('\n')
	}

	sb.WriteString("|1|2|3|4|5|6|7|")

	return sb.String()
}

func cSpotToString(c games.CPlr) string {
	switch c {
	case games.CNone:
		return "🔳"
	case games.CPlr1Max:
		return "🔵"
	case games.CPlr2Min:
		return "🔴"
	default:
		return ""
	}
}
