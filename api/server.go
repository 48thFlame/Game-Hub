package api

import (
	"log"
	"net/http"
)

func RunApi(port string) error {
	http.HandleFunc("/mastermind", handleMastermindRequests)

	log.Printf("Running API on port \"%v\"\n", port)

	return http.ListenAndServe(port, nil)
}
