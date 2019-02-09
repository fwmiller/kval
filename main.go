package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/fwmiller/kval/internal/cli"
	"github.com/fwmiller/kval/internal/kval"
)

func main() {
	fmt.Println("kval (C) Frank W Miller")

	go startTcpServer()

	db, err := kval.New()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	cli := cli.New(db)

	/* Command line client loop */
	stdin := bufio.NewReader(os.Stdin)
	var s1, s2 string
	s3 := ""
	multiline := false
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
		case "list", "l":
			cli.List()
		case "time", "t":
			cli.Time()
		case "help", "h":
			cli.Help()
		case "exists", "e":
			getArg(s4, cli.Exists)
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

const (
	ConnHost = "localhost"
	ConnPort = "6380"
	ConnType = "tcp"
)

func startTcpServer() {
	l, err := net.Listen(ConnType, ConnHost + ":" + ConnPort)
	if err != nil {
		fmt.Println("Listen failed (", err.Error, ")")
		return
	}
	defer l.Close()

	fmt.Println("Listening on port " + ConnPort)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Accept failed (", err.Error, ")")
			continue
		}
		go handleTcpConnection(c)
	}
}

func handleTcpConnection(c net.Conn) {
	buf := make([]byte, 1024)
	_, err := c.Read(buf)
	if (err != nil) {
		fmt.Println("Read failed (", err.Error, ")")
	} else {
		c.Write(buf)
	}
	c.Close()
}
