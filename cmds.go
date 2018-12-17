package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var dbNameCheck = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`).MatchString

func CliCreate(args string) {
	dbname := strings.TrimSpace(args)
	if !dbNameCheck(dbname) {
		fmt.Printf("Illegal characters in %v\n", dbname)
		return
	}
	dbpath := kvaldir + "/" + dbname
	fmt.Printf("Create database %v", dbpath)
	err := os.Mkdir(dbpath, 0777)
	if err != nil {
		fmt.Printf(" failed")
	}
	fmt.Printf("\n")
}

func CliRemove(args string) {
	fmt.Printf("Remove args = %v\n", args)
}

func CliSet(args string) {
	fmt.Printf("Set key-value pair args = %v\n", args)
}

func CliGet(args string) {
	fmt.Printf("Get value for key args = %v\n", args)
}

func CliDel(args string) {
	fmt.Printf("Delete key-value pair args = %v\n", args)
}

func CliHelp() {
	fmt.Println("Help (add something useful here)")
}
