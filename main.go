package main

import (
	"fmt"
	"os"

	"github.com/sojebsikder/go-mirror/cmd"
)

var version = "0.0.1"

func showUsage() {
	fmt.Println("Usage:")
	fmt.Println("  mirror start")
	fmt.Println()
	fmt.Println("  mirror help")
	fmt.Println("  mirror version")
	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	cmdName := os.Args[1]

	switch cmdName {
	case "start":
		cmd.Mirror()
	case "help":
		showUsage()
	case "version":
		fmt.Println("mirror version " + version)
	default:
		fmt.Println("Unknown command:", cmdName)
		fmt.Println("Use 'mirror help' to see available commands.")
		os.Exit(1)
	}
}
