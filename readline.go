package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text := ""
	for text != "exit" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("dql> ")
		text, _ = reader.ReadString('\n')
		fmt.Println(text)
	}
}
