package games

import (
	"errors"
)

const (
	CRowsNum = 6
	CColsNum = 7
)

type CSpot int

const (
	CNone CSpot = iota
	CPlr1       // being player 1
	CPlr2       // being player 2
)

type CBoard = [CRowsNum][CColsNum]CSpot

func NewConnect4Game() *Connect4Game {
	return &Connect4Game{
		Board:   CBoard{},
		PlrTurn: CPlr1,
		PlrWon:  CNone,
	}
}

type Connect4Game struct {
	Board   CBoard `json:"board"`   // an 2d array representing the board
	PlrTurn CSpot  `json:"plrTurn"` // players turn 1/2
	PlrWon  CSpot  `json:"plrWon"`  // player who won 1/2 or none - 0
}

func has4connected(s []CSpot, lookingFor CSpot) bool {
	if len(s) < 4 {
		return false
	}

	counter := 0
	for _, spot := range s {
		if spot == lookingFor {
			counter++
			if counter >= 4 {
				return true
			}
		} else {
			counter = 0
		}
	}

	return false
}

func (c *Connect4Game) won() (won bool) {
	lookingFor := c.PlrTurn
	// vertical checks
	for _, row := range c.Board {
		won = has4connected(row[:], lookingFor)
		if won {
			return
		}
	}

	// horizontal checks
	s := []CSpot{}
	for i := 0; i < CColsNum; i++ {
		for _, row := range c.Board {
			s = append(s, row[i])
		}
		won = has4connected(s, lookingFor)
		if won {
			return
		}
	}

	// diagonal checks
	for i := 0; i < CRowsNum; i++ {
		for j := 0; j < CColsNum; j++ {

			s = []CSpot{}
			for k := 0; k < 4; k++ {
				if i+k < CRowsNum && j+k < CColsNum {
					s = append(s, c.Board[i+k][j+k])
				}
			}
			won = has4connected(s, lookingFor)
			if won {
				return
			}

			s = []CSpot{}
			for k := 0; k < 4; k++ {
				if i+k < CRowsNum && j-k >= 0 {
					s = append(s, c.Board[i+k][j-k])
				}
			}
			won = has4connected(s, lookingFor)
			if won {
				return
			}
		}
	}

	return
}

func (c *Connect4Game) nextPlayer() {
	switch c.PlrTurn {
	case CPlr1:
		c.PlrTurn = CPlr2
	case CPlr2:
		c.PlrTurn = CPlr1
	}
}

// places a pieces for the current turned plr at the col returns whether plr who placed won(the bool), and whether was successful (whether the col not filled)(the error)
// must be given a valid col num, between 1-7
func (c *Connect4Game) Place(col int) (won bool, placed error) {
	col-- // turn from 1-7 to index 0-6

	// if placed doesn't become nil then the column is full
	placed = errors.New("column full")

	for i := CRowsNum - 1; i >= 0; i-- { // going from 6-0 to start at bottom
		if c.Board[i][col] == CNone {
			c.Board[i][col] = c.PlrTurn
			placed = nil
			break
		}
	}

	won = c.won()

	if !won {
		c.nextPlayer()
	}

	return
}
