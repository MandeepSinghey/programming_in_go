package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text := " "
	if len(os.Args) > 1 {

		text = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello World", text)
}
