package api

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func initDbManager() *dbManagerType {
	s := &dbManagerType{}
	s.masterIdGameNum = initializeMasterId()

	return s
}

type dbManagerType struct {
	masterIdGameNum int
}

var dbManager = initDbManager()

// initializeMasterId gets the highest game number in the masterGames folder to be used in creating new games
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
			log.Printf("While in stats init, %v does not start with a number\n", f.Name())
			continue
		}

		if num > largest {
			largest = num
		}
	}
	return largest + 1
}

func getDBMasterIdName() string {
	masterGameId := dbManager.masterIdGameNum
	id := strconv.Itoa(masterGameId)

	dbManager.masterIdGameNum++ // increment to next id
	return id
}
