package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/48thFlame/Game-Hub/games"
)

type Connect4PostRequest struct {
	Game  *games.Connect4Game   `json:"game"`
	Level games.Connect4AiLevel `json:"level"`
}

func handleConnect4Requests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Add("Content-Type", "application/json")

	if r.Method == "GET" {
		handleConnect4GET(w, r)
	} else if r.Method == "POST" {
		handleConnect4POST(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleConnect4GET(w http.ResponseWriter, r *http.Request) {
	game := games.NewConnect4Game()

	jsonData, err := marshalConnect4Game(game)
	if err != nil {
		log.Println("!! Error marshalling connect4 in GET: ", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, jsonData)
}

func handleConnect4POST(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var postData = &Connect4PostRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(postData)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	gameOver := postData.Game.GameState != games.CStatePlaying

	if !gameOver {
		// only if playing should make a move
		aiCol := games.FlameAiGetMove(*postData.Game, postData.Level)

		postData.Game.Turn(aiCol)
	}

	sendBackJson, err := marshalConnect4Game(postData.Game)
	if err != nil {
		log.Println("!! Error marshalling connect4 in POST: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Fprint(w, sendBackJson)

	saveConnect4Game(postData.Game)

}

func marshalConnect4Game(game *games.Connect4Game) (string, error) {
	b, err := json.Marshal(game)

	if err != nil {
		return "", err
	}
	return string(b), nil
}
