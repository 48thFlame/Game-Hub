package main

import (
	"fmt"
	"os"

	"github.com/48thFlame/Game-hub/boost/discord/data"
)

func main() {
	id := os.Args[1]
	user, err := data.LoadUser(id)
	if err != nil {
		panic(fmt.Errorf("error loading user to ban him from the feedback command: %v", err))
	}

	user.Feedback = true

	err = data.SaveData(data.GetUserFileName(id), user)
	if err != nil {
		panic(fmt.Errorf("error saving user after banning him from using the feedback command: %v", err))
	}
	fmt.Printf("User %v banned from using the feedback command\n", id)
}
