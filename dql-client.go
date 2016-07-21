package main

import (
	"flag"
	"fmt"
	"github.com/carmark/pseudo-terminal-go/terminal"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var serverUrl string

func main() {
	term, err := terminal.NewWithStdInOut()
	if err != nil {
		panic(err)
	}

	flag.StringVar(&serverUrl, "h", "localhost", "specify a server to connect to.  defaults to 'localhost'.")
	flag.Parse()

	defer term.ReleaseFromStdInOut() // defer this
	fmt.Println("\nWelcome to DQL Client V0.1.0 (Alpha)\n\nTo quit, type 'exit'\n")
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
	for _, command := range commands {
		response, err := sendCommand(command)
		if err {
			fmt.Print("Error: ")
		}
		fmt.Println(response + "\n")
	}
}

func sendCommand(command string) (string, bool) {

	resp, _ := http.PostForm(
		serverUrl+"/api/command",
		url.Values{"statement": {command}})
	defer resp.Body.Close()

	message, _ := ioutil.ReadAll(resp.Body)
	err := resp.StatusCode != 200
	return string(message), err
}
