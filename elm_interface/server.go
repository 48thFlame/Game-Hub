package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/avitar64/Boost-bot/games"
)

func newMasterJson() (string, error) {
	masterGame := games.NewMastermindGame()

	b, err := json.Marshal(masterGame)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func handleMastermindGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// fmt.Println(r.Body)
	// q := r.URL.Query()
	jsonData, err := newMasterJson()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("!! ERROR: %v\n", err)
	}
	fmt.Fprint(w, jsonData)
}

func main() {
	http.HandleFunc("/mastermind", handleMastermindGET)
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Println("Running...")

	http.ListenAndServe(":8080", nil)
}
