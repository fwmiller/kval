package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

var kvaldir string

func main() {
	fmt.Println("kval (C) Frank W Miller")

	/* Get path for kval directory */
	usr, _ := user.Current()
	kvaldir = usr.HomeDir + "/.kval"
	fmt.Printf("kval directory = %v\n", kvaldir)

	/* Check if kval directory exists */
	_, err := os.Stat(kvaldir)
	if os.IsNotExist(err) {
		fmt.Printf("Create kval directory %v", kvaldir)
		err = os.Mkdir(kvaldir, 0777)
		if err != nil {
			fmt.Printf(" failed\n")
			os.Exit(0)
		}
		fmt.Printf("\n")
	}
	/* Command line client loop */
	stdin := bufio.NewReader(os.Stdin)
	var s1 string
	for {
		fmt.Printf("kval> ")
		s1, _ = stdin.ReadString('\n')
		s2 := strings.Trim(s1, "\n")
		s3 := strings.TrimSpace(s2)
		s4 := strings.SplitAfterN(s3, " ", 2)
		if len(s4[0]) > 0 {
			switch strings.TrimSpace(s4[0]) {
			case "quit":
				os.Exit(0)
			case "create", "c":
				if len(s4) > 1 {
					CliCreate(s4[1])
				} else {
					fmt.Println("Missing argument")
				}

			case "remove", "r":
				if len(s4) > 1 {
					CliRemove(s4[1])
				} else {
					fmt.Println("Missing argument")
				}

			case "set", "s":
				if len(s4) > 1 {
					CliSet(s4[1])
				} else {
					fmt.Println("Missing argument")
				}

			case "get", "g":
				if len(s4) > 1 {
					CliGet(s4[1])
				} else {
					fmt.Println("Missing argument")
				}

			case "del", "d":
				if len(s4) > 1 {
					CliDel(s4[1])
				} else {
					fmt.Println("Missing argument")
				}

			case "help", "h":
				CliHelp()

			default:
				CliHelp()
			}
		}
	}
}
