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
	fmt.Println("kval (C) Frank W Miller")

	db, err := kval.New()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	cli := cli.New(db)

	/* Command line client loop */
	stdin := bufio.NewReader(os.Stdin)
	var s1 string
	for {
		if cli.Currdb != "" {
			fmt.Printf("%s", cli.Currdb)
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
		case "quit", "q":
			os.Exit(0)
		case "select":
			getArg(s4, cli.Select)
		case "create", "c":
			getArg(s4, cli.Create)
		case "remove", "r":
			getArg(s4, cli.Remove)
		case "keys", "k":
			cli.Keys()
		case "set", "s":
			getArg(s4, cli.Set)
		case "get", "g":
			getArg(s4, cli.Get)
		case "del", "d":
			getArg(s4, cli.Del)
		case "help", "h":
			cli.Help()
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
