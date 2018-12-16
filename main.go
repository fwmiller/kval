package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"regexp"
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
		s2 := strings.TrimSpace(s1)

		collapsespace := regexp.MustCompile(`\s+`)
		s3 := collapsespace.ReplaceAllString(s2, " ")

		if len(s3) > 0 {
			s4 := strings.Split(s3, " ")
			switch s4[0] {
			case "quit":
				os.Exit(0)
			case "create", "c":
				Create(s4)
			case "remove", "r":
				Remove(s4)
			case "set", "s":
				Set(s4)
			case "get", "g":
				Get(s4)
			case "del", "d":
				Del(s4)
			}
		}
	}
}
