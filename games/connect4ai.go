package games

import "sort"

func FlameAiGetMove(game Connect4Game) (col int) {
	var depth int

	if game.TurnNum < 3 {
		depth = 7
	} else if game.TurnNum < 15 {
		depth = 8
	} else if game.TurnNum < 22 {
		depth = 9
	} else if game.TurnNum < 26 {
		depth = 11
	} else {
		depth = game.TurnNum * game.TurnNum
	}

	initOrdered := cGetOrderedAvailableMoves(game.Board)
	quickSearchResults := cMoveRatingsToMoves(cSearch(game, 4, initOrdered))
	mRatings := cSearch(game, depth, quickSearchResults)

	bestMove := mRatings[0]

	return bestMove.move
}

func cSearch(game Connect4Game, depth int, moves []int) (s []cMoveRating) {
	isMaximizingPlayer := game.PlrTurn == CPlr1Max
	moveRatingCh := make(chan cMoveRating)

	for _, moveCol := range moves {
		go func(m int) {
			posEval := cMinMax(
				cPutPieceOnBoard(game.Board, m, game.PlrTurn),
				depth,
				-cHighNumber,
				cHighNumber,
				// opposite of max because is placing now the other one
				!isMaximizingPlayer)

			moveRatingCh <- cMoveRating{move: m, eval: posEval}
		}(moveCol)
	}

	numOfPossible := len(moves)
	for i := 0; i < numOfPossible; i++ { // for each goroutine started, get its result
		t := <-moveRatingCh

		s = append(s, t)
	}

	if isMaximizingPlayer {
		cSortMovesHighLow(s)
	} else {
		cSortMovesLowHigh(s)
	}

	return s
}

func cMinMax(board CBoard, depth int, alpha int, beta int, maximizingPlayer bool) int {
	if gs := cGetGameState(board); gs != CStatePlaying {
		switch gs {
		case CStateDraw:
			return 0
		case CStatePlr1Won:
			return cHighNumber + depth
		case CStatePlr2Won:
			return -cHighNumber - depth
		}
	} else if depth == 0 {
		return cStaticEval(&board)
	}

	possibleMoves := cGetOrderedAvailableMoves(board)
	if maximizingPlayer {
		maxEval := -cHighNumber

		for _, move := range possibleMoves {
			newEval := cMinMax(cPutPieceOnBoard(board, move, CPlr1Max), depth-1, alpha, beta, false)
			maxEval = cMax(maxEval, newEval)

			alpha = cMax(alpha, newEval)
			if beta <= alpha {
				break
			}
		}

		return maxEval
	} else {
		minEval := cHighNumber

		for _, move := range possibleMoves {
			newEval := cMinMax(cPutPieceOnBoard(board, move, CPlr2Min), depth-1, alpha, beta, true)
			minEval = cMin(minEval, newEval)

			beta = cMin(beta, newEval)
			if beta <= alpha {
				break
			}
		}

		return minEval
	}
}

const cHighNumber int = 100_000_000_000

const (
	cWonS              = cHighNumber // won/lost
	cCenterS           = 381         // col 4
	cMidS              = 59          // col 3/5
	cWinnableConnect3S = 387
	cWinnableConnect2S = 152
)

const ( // index num starts at 0
	cCenterCol = 3
	cMidCol1   = 2
	cMidCol2   = 4
)

type cMoveRating struct {
	move int
	eval int
}

func cMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func cMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func cInSlice(s []int, a int) bool {
	for _, obj := range s {
		if obj == a {
			return true
		}
	}

	return false
}

func cWinnableN(s []CPlr, piece CPlr, n int) bool {
	counter := 0

	for _, spot := range s {
		switch spot {
		case piece:
			counter++
			continue
		case CNone:
			continue
		default: // its opponent
			return false
		}
	}

	return counter == n // only equal not bigger so connect 3 don't count also as connect 2
}

func cGetCombinations(board CBoard) [][]CPlr {
	/*
		[[2 0 0 0 2 2 2]
			[1 1 1 0 1 0 1]
			[1 2 2 0 0 0 2]
			[2 1 1 1 1 1 1]
			[1 0 0 2 2 2 2]
			[0 0 1 1 1 1 2]]
	*/
	var combinations [][]CPlr

	// row combinations
	for _, row := range board {
		for colI := 0; colI < 4; colI++ {
			combination := row[colI : colI+4]
			combinations = append(combinations, combination)
		}
	}

	// col combination no because they can just be blocked
	// // col combinations
	// for col := 0; col < CColsNum; col++ {
	// 	for rowI := 0; rowI < 3; rowI++ {
	// 		combination := []CPlr{board[rowI][col], board[rowI+1][col], board[rowI+2][col], board[rowI+3][col]}
	// 		combinations = append(combinations, combination)
	// 	}
	// }

	for rowI := 0; rowI < 3; rowI++ {
		for colI := 0; colI < 4; colI++ {
			combination := []CPlr{board[rowI][colI], board[rowI+1][colI+1], board[rowI+2][colI+2], board[rowI+3][colI+3]}
			combinations = append(combinations, combination)
		}
	}

	// iterate over every diagonal (starting from top-right corner)
	for rowI := 0; rowI < 3; rowI++ {
		for j := 3; j < 7; j++ {
			combination := []CPlr{board[rowI][j], board[rowI+1][j-1], board[rowI+2][j-2], board[rowI+3][j-3]}
			combinations = append(combinations, combination)
		}
	}

	return combinations
}

func cGetOrderedAvailableMoves(board CBoard) []int {
	s := cGetAvailableMoves(board)
	prob := []int{3, 4, 2, 5, 1, 6, 0}

	c := make([]int, 0)

	for _, good := range prob {
		if cInSlice(s, good) {
			c = append(c, good)
		}
	}

	return c
}

func cSortMovesHighLow(s []cMoveRating) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].eval > s[j].eval
	})
}

func cSortMovesLowHigh(s []cMoveRating) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].eval < s[j].eval
	})
}

func cMoveRatingsToMoves(s []cMoveRating) []int {
	r := make([]int, 0)
	for _, m := range s {
		r = append(r, m.move)
	}

	return r
}

func cStaticEval(board *CBoard) (score int) {
	score += cGetConnectionScore(board)
	score += cGetLocationScore(board)

	return score
}

func cGetLocationScore(board *CBoard) (score int) {
	centerPlr1, centerPlr2 := cGetCenters(board)
	midPl1, midPlr2 := cGetMiddles(board)

	score += (centerPlr1 - centerPlr2) * cCenterS
	score += (midPl1 - midPlr2) * cMidS
	return score
}

func cGetCenters(board *CBoard) (centerPlr1, centerPlr2 int) {
	for _, row := range board {
		spot := row[cCenterCol]
		if spot == CPlr1Max {
			centerPlr1++
		} else if spot == CPlr2Min {
			centerPlr2++
		}
	}

	return centerPlr1, centerPlr2
}

func cGetMiddles(board *CBoard) (midPlr1, midPlr2 int) {
	for _, row := range board {
		spot1 := row[cMidCol1]
		spot2 := row[cMidCol2]

		if spot1 == CPlr1Max || spot2 == CPlr1Max {
			midPlr1++
		} else if spot1 == CPlr2Min || spot2 == CPlr2Min {
			midPlr2++
		}
	}

	return midPlr1, midPlr2
}

func cGetConnectionScore(board *CBoard) (score int) {
	b := *board
	possibleMoves := cGetAvailableMoves(b)
	for _, move := range possibleMoves {
		b = cPutPieceOnBoard(b, move, CPlr(-1))
	}
	s := cGetCombinations(b)

	number3Plr1, number2Plr1 := cGetConnectionsForPlr(s, CPlr1Max)
	number3Plr2, number2Plr2 := cGetConnectionsForPlr(s, CPlr2Min)

	score += (number3Plr1 - number3Plr2) * cWinnableConnect3S
	score += (number2Plr1 - number2Plr2) * cWinnableConnect2S

	return score
}

func cGetConnectionsForPlr(s [][]CPlr, p CPlr) (number3s, number2s int) {
	var winnable2, winnable3 bool

	for _, combination := range s {
		winnable3 = cWinnableN(combination, p, 3)
		if winnable3 {
			number3s++
			continue
		}

		winnable2 = cWinnableN(combination, p, 2)
		if winnable2 {
			number2s++
		}
	}

	return number3s, number2s
}
