package games

import (
	"math/rand"
)

func Connect4GetAiMove(game Connect4Game) (col int) {
	a := cGetAvailableMoves(game.Board)
	col = a[rand.Intn(len(a))]

	return col
}
