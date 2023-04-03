package games

const (
	CRowsNum = 6
	CColsNum = 7
)

type CPlr int

const (
	CNone    CPlr = iota
	CPlr1Max      // being player 1
	CPlr2Min      // being player 2
)

type CStateOfGame int

const (
	CStatePlaying CStateOfGame = iota
	CStatePlr1Won
	CStatePlr2Won
	CStateDraw
)

func NewConnect4Game() *Connect4Game {
	return &Connect4Game{
		Board:     CBoard{},
		PlrTurn:   CPlr1Max,
		GameState: CStateDraw,
	}
}

type CBoard [CRowsNum][CColsNum]CPlr
type Connect4Game struct {
	Board     CBoard       `json:"board"`   // an 2d array representing the board
	PlrTurn   CPlr         `json:"plrTurn"` // players turn 1/2
	GameState CStateOfGame `json:"gameState"`
}

func (game *Connect4Game) Turn(col int) (good bool) {
	good = cMovePossible(game.Board, col)

	if good {
		game.Board = putPieceOnBoard(game.Board, col, game.PlrTurn)
		game.GameState = getGameState(game.Board)
		game.PlrTurn = getOtherPlayer(game.PlrTurn)
	}

	return good
}

func putPieceOnBoard(board CBoard, col int, piece CPlr) CBoard {
	col-- // turn from 1-7 to index 0-6

	for i := CRowsNum - 1; i >= 0; i-- { // going from 6-0 to start at bottom
		if board[i][col] == CNone {
			board[i][col] = piece
			break
		}
	}

	return board
}

func getGameState(board CBoard) CStateOfGame {
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
	s := []CPlr{}
	for i := 0; i < CColsNum; i++ {

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

			s = []CPlr{}
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

func getOtherPlayer(plr CPlr) CPlr {
	switch plr {
	case CPlr1Max:
		return CPlr2Min
	case CPlr2Min:
		return CPlr1Max
	}

	return CNone
}
