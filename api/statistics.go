package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/48thFlame/Game-Hub/games"
)

const masterDBFolderName = "masterGames"

type SManagerType struct {
	statsMasterId int
}

var statsManager = initStatsManager()

func initStatsManager() *SManagerType {
	s := &SManagerType{}
	s.statsMasterId = initializeMasterId()
	return s
}

func initializeMasterId() int {
	files, err := os.ReadDir(fmt.Sprintf("./db/%v", masterDBFolderName))
	if err != nil {
		log.Fatalf("Could not read %v directory!! %v\n", masterDBFolderName, err)
	}

	var largest int

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".json") {
			continue
		}

		parts := strings.Split(f.Name(), ".")
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("While in stats init, %v does not start with a number\n", f)
			continue
		}

		if num > largest {
			largest = num
		}
	}
	return largest + 1
}

func getDBMasterIdName() string {
	idStat := statsManager.statsMasterId
	id := strconv.Itoa(idStat)
	statsManager.statsMasterId++
	return id
}

func saveMastermindGame(game *games.MastermindGame) {
	// should only save if game finished
	over := game.Won || len(game.Guesses) >= games.MasterGameLen

	if over {
		err := saveJsonTo(fmt.Sprintf("%v/%v.json", masterDBFolderName, getDBMasterIdName()), game)
		if err != nil {
			log.Println("!! ERROR:", err)
		}
	}
}

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
