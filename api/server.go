package api

import (
	"log"
	"net/http"
)

func handleConnect4Requests(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Connect4"))
}

func RunApi(port string) error {
	http.HandleFunc("/mastermind", handleMastermindRequests)

	http.HandleFunc("/connect4", handleConnect4Requests)

	log.Printf("Running API at localhost:%v\n", port)

	return http.ListenAndServe(port, nil)
}
