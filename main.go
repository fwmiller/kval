package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fwmiller/kval/internal/cli"
	"github.com/fwmiller/kval/internal/kval"
)

func main() {
	multiline := false

	fmt.Println("kval (C) Frank W Miller")

	db, err := kval.New()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	cli := cli.New(db)

	/* Command line client loop */
	stdin := bufio.NewReader(os.Stdin)
	var line string = ""
	var s1 string
	for {
		if cli.Currdb != "" {
			fmt.Printf("%s", cli.Currdb)
		}
		if multiline {
			fmt.Printf("\\ ")
		} else {
			fmt.Printf("> ")
		}
		s1, _ = stdin.ReadString('\n')
		s2 := strings.Trim(s1, "\n")
		s3 := strings.TrimSpace(s2)
		if len(s3) == 0 && !multiline {
			continue
		}
		// Check for multiline input
		if (strings.HasSuffix(s3, "\\")) {
			multiline = true
			line += strings.TrimRight(s3, "\\")
			line += "\n"
			continue
		}
		multiline = false
		line += s3

		tokens := strings.SplitAfterN(line, " ", 2)
		if len(tokens) == 0 {
			fmt.Println("Missing argument")
			continue
		}
		line = ""
		fmt.Println(tokens)

		switch strings.TrimSpace(tokens[0]) {
		case "quit", "q":
			os.Exit(0)
		case "select":
			getArg(tokens, cli.Select)
		case "create", "c":
			getArg(tokens, cli.Create)
		case "remove", "r":
			getArg(tokens, cli.Remove)
		case "keys", "k":
			cli.Keys()
		case "set", "s":
			getArg(tokens, cli.Set)
		case "get", "g":
			getArg(tokens, cli.Get)
		case "del", "d":
			getArg(tokens, cli.Del)
		case "list", "l":
			cli.List()
		case "time", "t":
			cli.Time()
		case "help", "h":
			cli.Help()
		case "exists", "e":
			getArg(tokens, cli.Exists)
		default:
			cli.Help()
		}
	}
}

func getArg(args []string, f func(arg string)) {
	if len(args) < 2 {
		fmt.Println("Missing argument")
	} else {
		f(args[1])
	}
}
