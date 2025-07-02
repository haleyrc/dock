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
	case "list":
		if err := c.List(ctx); err != nil {
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
	fmt.Fprintln(os.Stderr, "  clean    Kills and removes all containers and images")
	fmt.Fprintln(os.Stderr, "  cleanall Same as clean, but also prunes the build cache")
	fmt.Fprintln(os.Stderr, "  list     Lists all containers and images")
	fmt.Fprintln(os.Stderr)
}
