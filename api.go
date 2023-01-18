package main

import (
	"log"

	"github.com/48thFlame/Game-Hub/api"
)

func main() {
	err := api.RunApi(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
