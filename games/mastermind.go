package games

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
)

const (
	masterGameLen        = 7
	masterHighlighter    = "**"
	masterSecretEmoji    = "❓"
	masterBoardSeparator = " -- "
)

type MasterColor int

const (
	Empty MasterColor = iota
	Red
	Orange
	Yellow
	Green
	Blue
	Purple
)

var masterColors = []MasterColor{Red, Orange, Yellow, Green, Blue, Purple}

func (c MasterColor) String() string {
	switch c {
	case Empty:
		return "🔳"
		// return "empty"
	case Red:
		return "🟥"
		// return "red"
	case Orange:
		return "🟧"
		// return "orange"
	case Yellow:
		return "🟨"
		// return "yellow"
	case Green:
		return "🟩"
		// return "green"
	case Blue:
		return "🟦"
		// return "blue"
	case Purple:
		return "🟪"
		// return "purple"
	default:
		return ""
	}
}

func ConvertLetterToColor(letter string) MasterColor {
	switch letter {
	case "r":
		return Red
	case "o":
		return Orange
	case "y":
		return Yellow
	case "g":
		return Green
	case "b":
		return Blue
	case "p":
		return Purple
	default:
		return Empty
	}
}

type MasterResult int

const (
	Blank MasterResult = iota
	White
	Black
)

var perfectGuessResult = []MasterResult{Black, Black, Black, Black}

func (mr MasterResult) String() string {
	switch mr {
	case Blank:
		return "🔳"
	case White:
		return "❎"
	case Black:
		return "✅"
	default:
		return ""
	}
}

// returns a new colorSet, usually will be used to generate the answer for a mastermind game.
func getNewRandomColorSet() [4]MasterColor {
	shuffledColors := masterColors[:]
	rand.Shuffle(
		len(shuffledColors),
		func(i, j int) { shuffledColors[i], shuffledColors[j] = shuffledColors[j], shuffledColors[i] },
	)
	answer := [4]MasterColor{}
	copy(answer[:], shuffledColors)

	return answer
}

func NewMastermindGame() *MastermindGame {
	return &MastermindGame{
		// Answer:  [4]MasterColor{Red, Orange, Yellow, Green},
		Answer:  getNewRandomColorSet(),
		Guesses: [][4]MasterColor{},
		results: [][]MasterResult{},
	}
}

type MastermindGame struct {
	Answer  [4]MasterColor   `json:"answer"`
	Guesses [][4]MasterColor `json:"guesses"`
	results [][]MasterResult
}

func (m *MastermindGame) String() (str string) {
	str += masterHighlighter + "Answer:" + masterHighlighter + "\n"
	str += strings.Repeat(masterSecretEmoji + " ", 4) + masterBoardSeparator + strings.Repeat(Black.String() + " ", 4) + "\n"

	var guess [4]MasterColor
	var result []MasterResult

	for i := 0; i < masterGameLen; i++ {
		if len(m.Guesses) > i { //if guessed up until now, use the guess, otherwise use a blank/empty guess
			guess = m.Guesses[i]
			result = m.results[i]
		} else {
			guess = [4]MasterColor{Empty, Empty, Empty, Empty}
			result = []MasterResult{Blank, Blank, Blank, Blank}
		}

		str += masterHighlighter + "Round " + fmt.Sprint(i+1, ":") + masterHighlighter
		str += "\n"

		for _, color := range guess {
			str += color.String()
			str += " "
		}
		str += masterBoardSeparator
		for _, result := range result {
			str += result.String()
			str += " "
		}
		if i != masterGameLen-1 { // if its not the last round then should add new line char
			str += "\n"
		}
	}

	return
}

// guess on the game, returns whether user won or not
func (m *MastermindGame) Guess(guess [4]MasterColor) bool {
	results := make([]MasterResult, 0)

	for cI, c := range m.Answer {
		if c == guess[cI] {
			results = append(results, Black)
		} else if contains(guess[:], c) {
			results = append(results, White)
		}
	}

	sort.Slice(
		results,
		func(i, j int) bool {
			return int(results[i]) > int(results[j])
		},
	)

	m.Guesses = append(m.Guesses, guess)
	m.results = append(m.results, results)

	return reflect.DeepEqual(results, perfectGuessResult)
}

func contains(s []MasterColor, i MasterColor) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}
