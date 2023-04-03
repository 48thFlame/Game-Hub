package games

func cGetAvailableMoves(board CBoard) []int {
	a := []int{}

	for i := 1; i < CColsNum+1; i++ {
		if cMovePossible(board, i) {
			a = append(a, i)
		}
	}

	return a
}

func cMovePossible(board CBoard, col int) (possible bool) {
	col-- // turn from 1-7 to index 0-6

	for i := CRowsNum - 1; i >= 0; i-- { // going from 6-0 to start at bottom
		if board[i][col] == CNone {
			possible = true
			break
		}
	}

	return possible
}

func cPlrHas4Connected(board CBoard, fPiece CPlr) (won bool) {
	// row checks
	for _, row := range board {
		won = has4connected(row[:], fPiece)
		if won {
			return won
		}
	}

	// col checks
	s := []CPlr{}
	for i := 0; i < CColsNum; i++ {

		for _, row := range board {
			s = append(s, row[i])
		}

		won = has4connected(s, fPiece)
		if won {
			return won
		}
	}

	// diagonal checks
	for i := 0; i < CRowsNum; i++ {
		for j := 0; j < CColsNum; j++ {

			s = []CPlr{}
			for k := 0; k < 4; k++ {
				if i+k < CRowsNum && j+k < CColsNum {
					s = append(s, board[i+k][j+k])
				}
			}
			won = has4connected(s, fPiece)
			if won {
				return won
			}

			s = []CPlr{}
			for k := 0; k < 4; k++ {
				if i+k < CRowsNum && j-k >= 0 {
					s = append(s, board[i+k][j-k])
				}
			}
			won = has4connected(s, fPiece)
			if won {
				return won
			}
		}
	}

	return false
}

func has4connected(s []CPlr, lookingFor CPlr) bool {
	n := 4
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
