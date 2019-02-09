package cli

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fwmiller/kval/internal/kval"
)

var dbNameCheck = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`).MatchString

type Client struct {
	db     kval.DB
	Currdb string
}

func (c *Client) Loop() {
	/* Command line client loop */
	stdin := bufio.NewReader(os.Stdin)
	var s1, s2 string
	s3 := ""
	multiline := false
	for {
		if c.Currdb != "" {
			fmt.Printf("%s", c.Currdb)
		}
		if multiline {
			fmt.Printf("\\ ")
		} else {
			fmt.Printf("> ")
		}

		s1, _ = stdin.ReadString('\n')
		s2 = strings.TrimSpace(strings.Trim(s1, "\n"))
		if len(s2) == 0 {
			continue
		}
		if strings.HasSuffix(s2, "\\") {
			multiline = true
			s3 += strings.TrimRight(s2, "\\")
			s3 += "\n"
			continue
		}
		multiline = false
		s3 += s2

		s4 := strings.SplitAfterN(s3, " ", 2)
		if len(s4) == 0 {
			fmt.Println("Missing argument")
			continue
		}
		s3 = ""

		switch strings.TrimSpace(s4[0]) {
		case "quit", "q":
			os.Exit(0)
		case "select":
			getArg(s4, c.Select)
		case "create", "c":
			getArg(s4, c.Create)
		case "remove", "r":
			getArg(s4, c.Remove)
		case "keys", "k":
			c.Keys()
		case "set", "s":
			getArg(s4, c.Set)
		case "get", "g":
			getArg(s4, c.Get)
		case "del", "d":
			getArg(s4, c.Del)
		case "list", "l":
			c.List()
		case "time", "t":
			c.Time()
		case "help", "h":
			c.Help()
		case "exists", "e":
			getArg(s4, c.Exists)
		default:
			c.Help()
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
