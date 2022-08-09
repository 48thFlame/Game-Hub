package main

func main() {
	mastermind := newCommand(newAliases("mastermind", "mm", "m"), mastermindCommand)
	connect4 := newCommand(newAliases("connect4", "c4", "4"), connect4Command)

	t := newTerminal([]*command{mastermind, connect4})
	t.run()
}