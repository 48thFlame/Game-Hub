package games

func cPutPieceOnBoard(board CBoard, col int, piece CPlr) CBoard {
	for i := CRowsNum - 1; i >= 0; i-- { // going from 6-0 to start at bottom
		if board[i][col] == CNone {
			board[i][col] = piece
			break
		}
	}

	return board
}

func cGetGameState(board CBoard) CStateOfGame {
	plr1Won := cDidPlrConnect4(board, CPlr1Max)
	if plr1Won {
		return CStatePlr1Won
	}

	plr2Won := cDidPlrConnect4(board, CPlr2Min)
	if plr2Won {
		return CStatePlr2Won
	}

	draw := len(cGetAvailableMoves(board)) == 0
	if draw {
		return CStateDraw
	}

	return CStatePlaying
}

func cDidPlrConnect4(board CBoard, fPiece CPlr) (won bool) {
	// row checks
	for _, row := range board {
		won = cPlrHasNConnected(4, row[:], fPiece)
		if won {
			return won
		}
	}

	// col checks
	for i := 0; i < CColsNum; i++ {
		s := []CPlr{}

		for _, row := range board {
			s = append(s, row[i])
		}

		won = cPlrHasNConnected(4, s, fPiece)
		if won {
			return won
		}
	}

	// diagonal checks
	for i := 0; i < CRowsNum; i++ {
		for j := 0; j < CColsNum; j++ {

			s := []CPlr{}
			for k := 0; k < 4; k++ {
				if i+k < CRowsNum && j+k < CColsNum {
					s = append(s, board[i+k][j+k])
				}
			}
			won = cPlrHasNConnected(4, s, fPiece)
			if won {
				return won
			}

			s = []CPlr{}
			for k := 0; k < 4; k++ {
				if i+k < CRowsNum && j-k >= 0 {
					s = append(s, board[i+k][j-k])
				}
			}
			won = cPlrHasNConnected(4, s, fPiece)
			if won {
				return won
			}
		}
	}

	return false
}

func cGetOtherPlayer(plr CPlr) CPlr {
	switch plr {
	case CPlr1Max:
		return CPlr2Min
	case CPlr2Min:
		return CPlr1Max
	}

	return CNone
}

func cGetAvailableMoves(board CBoard) []int {
	a := []int{}

	for i := 0; i < CColsNum; i++ {
		if cMovePossible(board, i) {
			a = append(a, i)
		}
	}

	return a
}

func cMovePossible(board CBoard, col int) (possible bool) {
	for i := 0; i < CRowsNum; i++ { // start at top and f sees an empty spot there is at least one possible room
		if board[i][col] == CNone {
			return true
		}
	}

	return false
}

func cPlrHasNConnected(n int, s []CPlr, lookingFor CPlr) bool {
	if len(s) < n {
		return false
	}

	counter := 0
	for _, spot := range s {
		if spot == lookingFor {
			counter++
			if counter >= n {
				return true
			}
		} else {
			counter = 0
		}
	}

	return false
}
