package games

import (
	"errors"
	"strings"
)

const (
	connectRowsNum = 6
	connectColsNum = 7
)

type connectSpot int

const (
	ConnectEmpty connectSpot = iota
	Connect1                 // being one players pieces
	Connect2                 // being the other players pieces
)

func (c connectSpot) String() string {
	switch c {
	case ConnectEmpty:
		return "🔳"
	case Connect1:
		return "🔵"
	case Connect2:
		return "🔴"
	default:
		return ""
	}
}

func NewConnect4Game() *Connect4Game {
	return &Connect4Game{}
}

type Connect4Game [connectRowsNum][connectColsNum]connectSpot

func (c *Connect4Game) String() string {
	var sb strings.Builder

	for i := 0; i < connectRowsNum; i++ {
		for j := 0; j < connectColsNum; j++ {
			sb.WriteString(c[i][j].String())
		}

		if i < connectRowsNum-1 { // if its the last row shouldn't add newline char
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func has4connected(s []connectSpot, lookingFor connectSpot) bool {
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

func (c *Connect4Game) won(lookingFor connectSpot) (won bool) {
	// vertical checks
	for _, row := range c {
		won = has4connected(row[:], lookingFor)
		if won {
			return
		}
	}

	// horizontal checks
	s := []connectSpot{}
	for i := 0; i < connectColsNum; i++ {
		for _, row := range c {
			s = append(s, row[i])
		}
		won = has4connected(s, lookingFor)
		if won {
			return
		}
	}

	// diagonal checks
	for i := 0; i < connectRowsNum; i++ {
		for j := 0; j < connectColsNum; j++ {

			s = []connectSpot{}
			for k := 0; k < 4; k++ {
				if i+k < connectRowsNum && j+k < connectColsNum {
					s = append(s, c[i+k][j+k])
				}
			}
			won = has4connected(s, lookingFor)
			if won {
				return
			}

			s = []connectSpot{}
			for k := 0; k < 4; k++ {
				if i+k < connectRowsNum && j-k >= 0 {
					s = append(s, c[i+k][j-k])
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

// places a pieces for the given plrNum 1|2 at the col returns whether plr who placed won(the bool), and whether was successful (whether the col not filled)(the error)
// must be given a valid col num, between 1-7
func (c *Connect4Game) Place(plrNum, col int) (won bool, placed error) {
	var toPlace connectSpot
	col--

	switch plrNum {
	case 1:
		toPlace = Connect1
	case 2:
		toPlace = Connect2
	default:
		return false, errors.New("invalid plrNum")
	}

	placed = errors.New("column full")         // if placed doesn't become nil then the column is full
	for i := connectRowsNum - 1; i >= 0; i-- { // going from 6-0 to start at bottom
		if c[i][col] == ConnectEmpty {
			c[i][col] = toPlace
			placed = nil
			break
		}
	}

	won = c.won(toPlace)

	return
}
