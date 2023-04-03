package main

import (
	"math/rand"
	"time"

	shell "github.com/48thFlame/Command-Shell"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	mm := shell.NewCommand("mm", 0, mastermindCommandHandler)
	c := shell.NewCommand("c", 0, connect4CommandHandler)

	s, err := shell.NewShell(mm, c)
	if err != nil {
		panic(err)
	}

	s.Run()
}
