package data

import (
	"fmt"

	"github.com/avitar64/Boost-bot/games"
)

type Mastermind = games.MastermindGame

func GetMastermindFileName(userID string) string {
	return fmt.Sprintf("live-games/master-%v.json", userID)
}
