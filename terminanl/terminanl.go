package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func newAliases(a ...string) []string {
	return a
}

func newCommand(aliases []string, handler func([]string)) *command {
	return &command{
		aliases: aliases,
		handler: handler,
	}
}

type command struct {
	aliases []string
	handler func([]string)
}

func newTerminal(cmds []*command) *terminal {
	t := &terminal{
		cmdMap: make(map[string]*command),
	}

	cmds = append(
		cmds,
		newCommand(
			newAliases("e", "exit"), func(args []string) {
				os.Exit(0)
			},
		),
	)

	for _, c := range cmds {
		for _, name := range c.aliases {
			t.cmdMap[name] = c
		}
	}

	return t
}

type terminal struct {
	cmdMap map[string]*command
}

func (t *terminal) handleInput(input []string) {
	if len(input[0]) == 0 {
		return
	}

	if cmd, ok := t.cmdMap[input[0]]; ok {
		cmd.handler(input[1:])
	} else {
		fmt.Printf("Command '%v' not found\n", input[0])
	}
}

func (t *terminal) run() {
	fmt.Println("Boost shell, type 'exit' to exit")
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "exit" {
			break
		}

		t.handleInput(strings.Split(line, " "))
	}
}

func main() {
	myCommand := newCommand(
		newAliases("c", "cmd"),
		func(args []string) {
			fmt.Println("Test command called with args:", args)
		},
	)

	t := newTerminal([]*command{myCommand})
	t.run()
}
