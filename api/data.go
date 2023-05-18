package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/48thFlame/Game-Hub/games"
)

const masterDBFolderName = "masterGames"

// saveToJson saves the given json to path in the db/ folder
func saveJsonTo(path string, data interface{}) error {
	f, err := os.Create(fmt.Sprintf("./db/%v", path))
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	err = enc.Encode(data)
	if err != nil {
		return err
	}

	return nil

}

func saveMastermindGame(game *games.MastermindGame) {
	// should only save if game finished
	over := game.Won || len(game.Guesses) >= games.MasterGameLen

	if over {
		err := saveJsonTo(fmt.Sprintf("%v/%v.json", masterDBFolderName, getDBMasterIdName()), game)
		if err != nil {
			log.Println("!! Error saving master game:", err)
		}
	}
}
