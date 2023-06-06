package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/48thFlame/Game-Hub/games"
)

const (
	masterDBFolderName   = "masterGames"
	connect4DBFolderName = "connect4Games"
)

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

func saveConnect4Game(game *games.Connect4Game) {
	// should only save if game finished
	over := game.GameState != games.CStatePlaying

	if over {
		err := saveJsonTo(fmt.Sprintf("%v/%v.json", connect4DBFolderName, getDBConnect4IdName()), game)
		if err != nil {
			log.Println("!! Error saving connect4 game:", err)
		}
	}
}
