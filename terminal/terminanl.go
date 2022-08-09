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

func newCommand(aliases []string, handler commandHandler) *command {
	return &command{
		aliases: aliases,
		handler: handler,
	}
}

type commandHandler func(*terminal, []string)
type command struct {
	aliases []string
	handler commandHandler
}

func newTerminal(cmds []*command) *terminal {
	t := &terminal{
		cmds:       make([]*command, 0),
		handlerMap: make(map[string]commandHandler),
		data:       make(map[string]any),
	}

	cmds = append(
		cmds,
		newCommand(
			newAliases("exit", "e"),
			func(_ *terminal, args []string) {
				os.Exit(0)
			},
		),
		newCommand(
			newAliases("help", "h"),
			func(t *terminal, args []string) {
				for _, c := range t.cmds {
					fmt.Println(c.aliases[0])
				}
			},
		),
	)

	t.cmds = cmds

	for _, c := range cmds {
		for _, name := range c.aliases {
			t.handlerMap[name] = c.handler
		}
	}

	return t
}

type terminal struct {
	cmds       []*command
	handlerMap map[string]commandHandler
	data       map[string]any
}

func (t *terminal) handleInput(input []string) {
	if len(input[0]) == 0 {
		return
	}

	if cmd, ok := t.handlerMap[input[0]]; ok {
		cmd(t, input[1:])
	} else {
		fmt.Printf("Command '%v' not found\n", input[0])
	}
}

func (t *terminal) run() {
	fmt.Println("Boost shell, type 'help' to see all commands")
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

func (t *terminal) Error(msg string) {
	fmt.Println("!!Error:", msg)
}
