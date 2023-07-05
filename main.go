package main

import (
	"context"
	"fmt"
	"os"

	"github.com/haleyrc/dock/internal/dock"
)

func main() {
	ctx := context.Background()

	cmd := "help"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	c, err := dock.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	switch cmd {
	case "clean":
		if err := c.Clean(ctx); err != nil {
			panic(err)
		}
	case "cleanall":
		if err := c.CleanAll(ctx); err != nil {
			panic(err)
		}
	case "help":
		printUsage()
		os.Exit(2)
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: dock COMMAND")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "COMMANDS:")
	fmt.Fprintln(os.Stderr, "\tclean\t\tKills and removes all containers and images")
	fmt.Fprintln(os.Stderr, "\tcleanall\tSame as clean, but also prunes the build cache")
	fmt.Fprintln(os.Stderr)
}
