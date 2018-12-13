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
	var rawcmdline string
	for {
		fmt.Printf("kval> ")
		rawcmdline, _ = stdin.ReadString('\n')
		cmdline := strings.TrimSuffix(rawcmdline, "\n")
		if len(cmdline) > 0 {
			fmt.Println(cmdline)
		}
	}
}
