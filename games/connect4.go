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

func (game *Connect4Game) place(col int) {
	col-- // turn from 1-7 to index 0-6

	for i := CRowsNum - 1; i >= 0; i-- { // going from 6-0 to start at bottom
		if game.Board[i][col] == CNone {
			game.Board[i][col] = game.PlrTurn
			break
		}
	}
}

func (game *Connect4Game) Turn(col int) (good bool) {
	good = cMovePossible(game.Board, col)

	if good {
		game.place(col)
		game.setGameState()
		game.nextPlayer()
	}

	return good
}

// changed player to next turn
func (game *Connect4Game) nextPlayer() {
	switch game.PlrTurn {
	case CPlr1Max:
		game.PlrTurn = CPlr2Min
	case CPlr2Min:
		game.PlrTurn = CPlr1Max
	}
}

func (game *Connect4Game) setGameState() {
	plr1Won := cPlrHas4Connected(game.Board, CPlr1Max)
	if plr1Won {
		game.GameState = CStatePlr1Won
		return
	}

	plr2Won := cPlrHas4Connected(game.Board, CPlr2Min)
	if plr2Won {
		game.GameState = CStatePlr2Won
		return
	}

	draw := len(cGetAvailableMoves(game.Board)) == 0
	if draw {
		game.GameState = CStateDraw
		return
	}

	game.GameState = CStatePlaying
}
