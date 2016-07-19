package main

import (
	"fmt"
	"github.com/carmark/pseudo-terminal-go/terminal"
	"io"
	"strings"
)

func main() {
	term, err := terminal.NewWithStdInOut()
	if err != nil {
		panic(err)
	}
	defer term.ReleaseFromStdInOut() // defer this
	fmt.Println("Welcome to DQL")
	term.SetPrompt("dql> ")

	line := ""
	var fullCommands []string
	partialCommand := ""

	for {
		line, err := term.ReadLine()
		if err == io.EOF {
			term.Write([]byte(line))
			fmt.Println()
			return
		}
		if strings.TrimSpace(line) == "exit" {
			fmt.Println("bye")
			return
		}

		if strings.TrimSpace(line) != "" {
			if trimmedStringEndsWith(line, ";") {
				term.SetPrompt("dql> ")
			} else {
				term.SetPrompt("   > ")
			}

			fullCommands, partialCommand = extractCommands(line, partialCommand)

			runCommands(fullCommands)
		}
	}
	term.Write([]byte(line))
}

func trimmedStringEndsWith(text string, end string) bool {
	text = strings.TrimSpace(text)
	return text[len(text)-1:] == end
}

func extractCommands(line string, partial string) ([]string, string) {
	fullLine := partial + line
	commands := strings.Split(fullLine, ";")

	partialCommand := strings.TrimSpace(commands[len(commands)-1])
	fullCommands := commands[:len(commands)-1]

	return fullCommands, partialCommand
}

func runCommands(commands []string) {
	if len(commands) == 0 {
		return
	}
	fmt.Print(len(commands))
	fmt.Println(" commands run\n")
}
