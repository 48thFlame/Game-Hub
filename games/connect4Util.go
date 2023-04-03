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

// remove top piece from given col
func cRemovePieceFromBoard(board CBoard, col int) CBoard {
	col-- // turn from 1-7 to index 0-6

	for i := 0; i < CRowsNum; i++ {
		if board[i][col] != CNone {
			board[i][col] = CNone
			break
		}
	}

	return board
}
