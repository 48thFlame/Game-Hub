package games

import (
	"fmt"
)

const inf int = 99999720

const (
	CWonS              = inf // won/lost
	CCenterS           = 243 // col 4
	CMidS              = 59  // col 3/5
	CWinnableConnect3S = 480
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
	possibleMoves := cGetAvailableMoves(game.Board)
	minEval := inf
	var bestMove int

	for moveI, moveCol := range possibleMoves {
		posEval := minMax(cPutPieceOnBoard(game.Board, moveCol, CPlr2Min), 7, -inf, inf, true)
		totalPosAnalyzed += timeCalledEval
		timeCalledEval = 0

		if posEval < minEval {
			minEval = posEval
			bestMove = moveI
		}
	}

	fmt.Println("Total pos:", totalPosAnalyzed)
	totalPosAnalyzed = 0

	return possibleMoves[bestMove]
}

var timeCalledEval, totalPosAnalyzed int

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

	possibleMoves := cGetAvailableMoves(board)
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
	timeCalledEval++
	score += getGameOverScore(board)
	score += getLocationScore(board)
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
