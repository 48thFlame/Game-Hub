package data

import (
	"fmt"
)

func GetMastermindFileName(userID string) string {
	return fmt.Sprintf("live-games/master-%v.json", userID)
}
