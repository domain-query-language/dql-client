package main

import (
	"fmt"
	"github.com/pkg/term"
)

func main() {

	char := ""
	arrow := ""
	newline := false
	for char != "q" {
		char, arrow, newline, _ = getChar()
		if arrow == "up" || arrow == "down" {
			fmt.Printf("\rPressed %s/        ", arrow)
		} else if newline {
			fmt.Println("")
		} else {
			fmt.Print(char)
		}
	}
	fmt.Println("Bye")
}

func getChar() (char string, keyCode string, newline bool, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)
	char = ""
	keyCode = ""
	newline = false

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
		if bytes[0] == '\n' {
			newline = true
		} else {
			char = string(bytes[0])
		}
	} else {
		newline = true
	}
	t.Restore()
	t.Close()
	return
}
