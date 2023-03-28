package games

import (
	"math/rand"
	"reflect"
	"sort"
)

const MasterGameLen = 7

type MasterColor int

const (
	Blank MasterColor = iota
	Red
	Orange
	Yellow
	Green
	Blue
	Purple
)

var masterColors = []MasterColor{Red, Orange, Yellow, Green, Blue, Purple}

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
		return Blank
	}
}

type MasterResult int

const (
	Empty MasterResult = iota
	White
	Black
)

var perfectGuessResult = []MasterResult{Black, Black, Black, Black}

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
		Answer:  getNewRandomColorSet(),
		Guesses: [][4]MasterColor{},
		Results: [][]MasterResult{},
	}
}

type MastermindGame struct {
	Won     bool             `json:"won"`
	Answer  [4]MasterColor   `json:"answer"`
	Guesses [][4]MasterColor `json:"guesses"`
	Results [][]MasterResult `json:"results"`
}

// guess on the game, returns whether user won or not
func (m *MastermindGame) Guess(guess [4]MasterColor) bool {
	results := getGuessResult(guess, m.Answer)

	m.Guesses = append(m.Guesses, guess)
	m.Results = append(m.Results, results)

	won := reflect.DeepEqual(results, perfectGuessResult)
	m.Won = won

	return won
}

// fills results for all already guessed (use when taking game out of json for example)
func (m *MastermindGame) FillResults() {
	for _, guess := range m.Guesses {
		m.Results = append(m.Results, getGuessResult(guess, m.Answer))
	}
}

func getGuessResult(guess, answer [4]MasterColor) []MasterResult {
	results := make([]MasterResult, 0)

	for cI, c := range answer {
		if guess[cI] == c {
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

	return results
}

func contains(s []MasterColor, i MasterColor) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}
