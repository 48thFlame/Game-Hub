package main

import (
	"fmt"
	"strconv"

	"github.com/avitar64/Boost-bot/games"
)

const (
	connect4CommandDataGameName    = "connect4"
	connect4CommandDataPlrTurnName = "plrTurn"
)

func connect4Command(t *terminal, args []string) {
	if _, ok := t.data[connect4CommandDataGameName]; !ok {
		fmt.Println("Creating new connect4 game...")
		game := games.NewConnect4Game()
		t.data[connect4CommandDataGameName] = game
		t.data[connect4CommandDataPlrTurnName] = 1
	}

	game := t.data[connect4CommandDataGameName].(*games.Connect4Game)
	argsLen := len(args)
	var err error
	var going, won bool
	var plrGoingNum, colNum int

	switch argsLen {
	case 1:
		going = true
		plrGoingNum = t.data[connect4CommandDataPlrTurnName].(int)

		// should get col num
		colNum, err = strconv.Atoi(args[0])
		if err != nil || colNum < 1 || colNum > 7 {
			t.Error(fmt.Sprintf("\"%v\" isn't a valid column number", args[0]))
			return
		}
	case 0:
	default:
		t.Error("You need to provide 1 or 0 args not more.")
		return
	}

	if going {
		won, err = game.Place(plrGoingNum, colNum)
		if err != nil {
			t.Error(err.Error())
			return
		}

		// should switch to next plr's turn
		var toSetPlrTurn int
		if plrGoingNum == 1 {
			toSetPlrTurn = 2
		} else {
			toSetPlrTurn = 1
		}
		t.data[connect4CommandDataPlrTurnName] = toSetPlrTurn
	}

	fmt.Println(game)
	if won {
		fmt.Printf("Player %v, won!\n", plrGoingNum)
		delete(t.data, connect4CommandDataGameName)
		delete(t.data, connect4CommandDataPlrTurnName)
	}
}

// func connect4Command(t *terminal, args []string) {
// 	var won bool
// 	var success error
// 	la := len(args)

// 	if _, ok := t.data[connect4CommandDataGameName]; !ok {
// 		fmt.Println("Creating new connect4 game...")
// 		t.data[connect4CommandDataGameName] = games.NewConnect4Game() // create the game object
// 		t.data[connect4CommandDataPlrTurnName] = 1
// 	}

// 	game := t.data[connect4CommandDataGameName].(*games.Connect4Game)
// 	plrNum := t.data[connect4CommandDataPlrTurnName].(int)

// 	fmt.Println(plrNum)

// 	going := la > 0
// 	if going {
// 		column, err := strconv.Atoi(args[0])
// 		if err != nil || column < 1 || column > 7 {
// 			t.Error(fmt.Sprintf("\"%v\" isn't a valid column number", args[0]))
// 			return
// 		}

// 		won, success = game.Place(plrNum, column)
// 		if success != nil {
// 			t.Error(success.Error())
// 			return
// 		}
// 		t.data[connect4CommandDataPlrTurnName] = ((plrNum + 1) % 2) // cycles between 1 and 2
// 		fmt.Println(((plrNum + 1) % 2))
// 	}

// 	fmt.Println(game)
// 	if won {
// 		fmt.Printf("Player %v won!\n", plrNum)
// 		delete(t.data, connect4CommandDataGameName)
// 		delete(t.data, connect4CommandDataPlrTurnName)
// 	}
// }
