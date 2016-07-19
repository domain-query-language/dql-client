package main

import (
	"fmt"
	"github.com/pkg/term"
	str "strings"
)

func main() {
	fmt.Print("\nWelcome to the DQL Client.\n\nCopyright 2016\n\n")
	readLine := true

	char := ";"
	arrow := ""
	for readLine {
		if isFullCommand(char) {
			fmt.Print("dql> ")
		} else {
			fmt.Print("  -> ")
		}
		char, arrow, _ = getChar()

		if arrow != "" {
			fmt.Print(arrow)
		}

		if char != "" {
			//fmt.Print(char)
		}

		if str.Contains(char, "exit") {
			readLine = false
			fmt.Print("Goodbye\n\n")
		}
	}
}

func isFullCommand(text string) bool {
	return text[len(text)-1:] == ";"
}

func getChar() (char string, keyCode string, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)
	char = ""
	keyCode = ""

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {

		if bytes[2] == 65 {
			// Up
			keyCode = "up"
		} else if bytes[2] == 66 {
			// Down
			keyCode = "down"
		} else if bytes[2] == 67 {
			// Right
			keyCode = "right"
		} else if bytes[2] == 68 {
			// Left
			keyCode = "left"
		}
	} else if numRead == 1 {
		char = string(bytes[0])
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
}
