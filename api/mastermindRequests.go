package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/48thFlame/Game-Hub/games"
)

func handleMastermindRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Add("Content-Type", "application/json")

	if r.Method == "GET" {
		handleMasterGET(w, r)
	} else if r.Method == "POST" {
		handleMasterPOST(w, r)
	}
}

func handleMasterGET(w http.ResponseWriter, r *http.Request) {
	game := games.NewMastermindGame()

	jsonData, err := marshalMastermindGHame(game)
	if err != nil {
		log.Println("!! Error marshalling master in GET: ", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, jsonData)
}

func handleMasterPOST(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var postStruct = &MastermindGuessPOSTRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(postStruct)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	postStruct.Game.Guess(postStruct.Guess)

	jsonData, err := marshalMastermindGHame(postStruct.Game)
	if err != nil {
		log.Println("!! Error marshalling master in POST: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	saveMastermindGame(postStruct.Game)

	fmt.Fprint(w, jsonData)
}

type MastermindGuessPOSTRequest struct {
	Game  *games.MastermindGame `json:"game"`
	Guess [4]games.MasterColor  `json:"guess"`
}

func marshalMastermindGHame(game *games.MastermindGame) (string, error) {
	b, err := json.Marshal(game)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
