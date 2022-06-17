package main

func main() {
	mastermind := newCommand(newAliases("mastermind", "mm", "m"), mastermindCommand)

	t := newTerminal([]*command{mastermind})
	t.run()
}