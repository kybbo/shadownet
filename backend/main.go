package main

import (
	"fmt"
	"os"

	"shadownet/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: shadownet <command>")
		return
	}

	switch os.Args[1] {
	case "init":
		cmd.RunInit()
	default:
		fmt.Println("Unkown command:", os.Args[1])
	}
}
