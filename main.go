package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fwmiller/kval/internal/kval"
)

var (
	currdb string
	db     kval.DB
)

func main() {
	fmt.Println("kval (C) Frank W Miller")

	var err error
	db, err = kval.New()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	/* Command line client loop */
	stdin := bufio.NewReader(os.Stdin)
	var s1 string
	for {
		if currdb != "" {
			fmt.Printf("%s", currdb)
		}
		fmt.Printf("> ")

		s1, _ = stdin.ReadString('\n')
		s2 := strings.Trim(s1, "\n")
		s3 := strings.TrimSpace(s2)
		s4 := strings.SplitAfterN(s3, " ", 2)
		if len(s4) == 0 {
			fmt.Println("Missing argument")
			continue
		}

		switch strings.TrimSpace(s4[0]) {
		case "quit":
			os.Exit(0)
		case "select":
			getArg(s4, CliDb)
		case "create", "c":
			getArg(s4, CliCreate)
		case "remove", "r":
			getArg(s4, CliRemove)
		case "keys", "k":
			CliKeys()
		case "set", "s":
			getArg(s4, CliSet)
		case "get", "g":
			getArg(s4, CliGet)
		case "del", "d":
			getArg(s4, CliDel)
		case "help", "h":
			CliHelp()
		default:
			CliHelp()
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
