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
	s.connect4IdGameNum = initializeConnect4Id()

	return s
}

type dbManagerType struct {
	masterIdGameNum   int
	connect4IdGameNum int
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
			log.Printf("While in stats init for master, %v does not start with a number\n", f.Name())
			continue
		}

		if num > largest {
			largest = num
		}
	}
	return largest + 1
}

func initializeConnect4Id() int {
	files, err := os.ReadDir(fmt.Sprintf("./db/%v", connect4DBFolderName))
	if err != nil {
		log.Fatalf("Could not read %v directory!! %v\n", connect4DBFolderName, err)
	}

	var largest int

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".json") {
			continue
		}

		parts := strings.Split(f.Name(), ".")
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("While in stats init for connect4, %v does not start with a number\n", f.Name())
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

func getDBConnect4IdName() string {
	connect4GameId := dbManager.connect4IdGameNum
	id := strconv.Itoa(connect4GameId)

	dbManager.connect4IdGameNum++ // increment to next id
	return id
}
