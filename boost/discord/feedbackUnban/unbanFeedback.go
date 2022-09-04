package main

import (
	"fmt"
	"os"

	"github.com/avitar64/Boost-bot/boost/discord/data"
)

func main() {
	id := os.Args[1]
	user, err := data.LoadUser(id)
	if err != nil {
		panic(fmt.Errorf("error loading user to unban him from the feedback command: %v", err))
	}

	user.Feedback = false

	err = data.SaveData(data.GetUserFileName(id), user)
	if err != nil {
		panic(fmt.Errorf("error saving user after unbanning him from using the feedback command: %v", err))
	}

	fmt.Printf("User %v unbanned from using the feedback command\n", id)
}
