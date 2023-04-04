package games

const inf int = 99999720

const (
	CWonS              = inf // won/lost
	CCenterS           = 381 // col 4
	CMidS              = 59  // col 3/5
	CWinnableConnect3S = 387
	CWinnableConnect2S = 152
)

const ( // index num starts at 0
	cCenterCol = 3
	cMidCol1   = 2
	cMidCol2   = 4
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Connect4GetAiMove(game Connect4Game) (col int) {
	possibleMoves := getOrderedMoves(game.Board)

	var depth int

	if game.TurnNum < 5 {
		depth = 6
	} else if game.TurnNum < 14 {
		depth = 7
	} else if game.TurnNum < 19 {
		depth = 9
	} else if game.TurnNum < 23 {
		depth = 11
	} else {
		depth = game.TurnNum / 2
	}

	isMaximizingPlayer := game.PlrTurn == CPlr1Max
	evalCh := make(chan [2]int)

	for _, moveCol := range possibleMoves {
		go func(m int) {
			// opposite of max because is placing now the other one
			posEval := minMax(cPutPieceOnBoard(game.Board, m, game.PlrTurn), depth, -inf, inf, !isMaximizingPlayer)
			evalCh <- [2]int{m, posEval}
		}(moveCol)
	}

	var extremeEval int
	if isMaximizingPlayer {
		extremeEval = -inf
	} else {
		extremeEval = inf
	}
	var bestMove int

	numOfPossible := len(possibleMoves)
	for i := 0; i < numOfPossible; i++ { // for each goroutine started, get its result
		t := <-evalCh

		posEval := t[1]

		if isMaximizingPlayer {
			if posEval > extremeEval {
				extremeEval = posEval
				bestMove = t[0]
			}
		} else {
			if posEval < extremeEval {
				extremeEval = posEval
				bestMove = t[0]
			}
		}
	}

	return bestMove
}

func inSlice(s []int, a int) bool {
	for _, obj := range s {
		if obj == a {
			return true
		}
	}

	return false
}

func getOrderedMoves(board CBoard) []int {
	s := cGetAvailableMoves(board)
	prob := []int{3, 4, 2, 5, 1, 6, 0}

	c := make([]int, 0)

	for _, good := range prob {
		if inSlice(s, good) {
			c = append(c, good)
		}
	}

	return c
}

func minMax(board CBoard, depth int, alpha int, beta int, maximizingPlayer bool) int {
	if depth == 0 {
		return staticEval(&board)
	} else if gs := cGetGameState(board); gs != CStatePlaying {
		switch gs {
		case CStateDraw:
			return 0
		case CStatePlr1Won:
			return (inf * depth) // make it smaller the deeper in this happens so comp will rather stop early threats
		case CStatePlr2Won:
			return (-inf * depth)
		}
	}

	possibleMoves := getOrderedMoves(board)
	if maximizingPlayer {
		maxEval := -inf

		for _, move := range possibleMoves {
			newEval := minMax(cPutPieceOnBoard(board, move, CPlr1Max), depth-1, alpha, beta, false)
			maxEval = max(maxEval, newEval)

			alpha = max(alpha, newEval)
			if beta <= alpha {
				break
			}
		}

		return maxEval
	} else {
		minEval := inf

		for _, move := range possibleMoves {
			newEval := minMax(cPutPieceOnBoard(board, move, CPlr2Min), depth-1, alpha, beta, true)
			minEval = min(minEval, newEval)

			beta = min(beta, newEval)
			if beta <= alpha {
				break
			}
		}

		return minEval
	}
}

func staticEval(board *CBoard) (score int) {
	score += getGameOverScore(board)
	score += getLocationScore(board)
	score += getConnectionScore(board)

	return score
}

func getGameOverScore(board *CBoard) (score int) {
	gs := cGetGameState(*board)

	switch gs {
	case CStatePlr1Won:
		return inf
	case CStatePlr2Won:
		return -inf
	}

	return 0
}

func getLocationScore(board *CBoard) (score int) {
	score += getPlr1CenterScore(board)
	score += getPlr2CenterScore(board)
	score += getPlr1MiddleScore(board)
	score += getPlr2MiddleScore(board)
	return
}

func getPlr1CenterScore(board *CBoard) int {
	centers := 0
	for _, row := range board {
		if row[cCenterCol] == CPlr1Max {
			centers++
		}
	}

	return centers * CCenterS
}

func getPlr2CenterScore(board *CBoard) int {
	centers := 0
	for _, row := range board {
		if row[cCenterCol] == CPlr2Min {
			centers--
		}
	}

	return centers * CCenterS
}

func getPlr1MiddleScore(board *CBoard) int {
	middles := 0
	for _, row := range board {
		if row[cMidCol1] == CPlr1Max {
			middles++
		}
		if row[cMidCol2] == CPlr1Max {
			middles++
		}
	}

	return middles * CMidS
}

func getPlr2MiddleScore(board *CBoard) int {
	middles := 0
	for _, row := range board {
		if row[cMidCol1] == CPlr2Min {
			middles--
		}
		if row[cMidCol2] == CPlr2Min {
			middles--
		}
	}

	return middles * CMidS
}

func getConnectionScore(board *CBoard) (score int) {
	s := getCombinations(board)
	score += getConnectionScoreForPlr1(s)
	score += getConnectionScoreForPlr2(s)

	return score
}

func getConnectionScoreForPlr1(s [][]CPlr) int {
	score := 0
	var winnable2, winnable3 bool

	for _, combination := range s {
		winnable3 = winnableN(combination, CPlr1Max, 3)
		if winnable3 {
			score += CWinnableConnect3S
			continue
		}

		winnable2 = winnableN(combination, CPlr1Max, 2)
		if winnable2 {
			score += CWinnableConnect2S
			continue
		}
	}

	return score
}

func getConnectionScoreForPlr2(s [][]CPlr) int {
	score := 0
	var winnable2, winnable3 bool

	for _, combination := range s {
		winnable3 = winnableN(combination, CPlr2Min, 3)
		if winnable3 {
			score -= CWinnableConnect3S
			continue
		}

		winnable2 = winnableN(combination, CPlr2Min, 2)
		if winnable2 {
			score -= CWinnableConnect2S
			continue
		}
	}

	return score
}

func getCombinations(board *CBoard) [][]CPlr {
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

	// col combinations
	for col := 0; col < CColsNum; col++ {
		for rowI := 0; rowI < 3; rowI++ {
			combination := []CPlr{board[rowI][col], board[rowI+1][col], board[rowI+2][col], board[rowI+3][col]}
			combinations = append(combinations, combination)
		}
	}

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

func winnableN(s []CPlr, piece CPlr, n int) bool {
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
