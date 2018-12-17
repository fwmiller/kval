package main

import (
	"fmt"
	"regexp"
	"strings"
)

var dbNameCheck = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`).MatchString

func CliDb(args string) string {
	/* Check for valid dbname */
	dbname := strings.TrimSpace(args)
	if !dbNameCheck(dbname) {
		fmt.Printf("Illegal characters in %v\n", dbname)
		return ""
	}
	/* Check for valid database */
	return KvalIsDb(dbname)
}

func CliCreate(args string) {
	/* Check for valid dbname */
	dbname := strings.TrimSpace(args)
	if !dbNameCheck(dbname) {
		fmt.Printf("Illegal characters in %v\n", dbname)
		return
	}
	/* Create new database */
	KvalCreateDb(dbname)
}

func CliRemove(args string) {
	fmt.Printf("Remove args = %v\n", args)
}

func CliSet(args string) {
	if currdb == "" {
		fmt.Println("Current database not set")
		return
	}
	s := strings.SplitAfterN(args, " ", 2)
	if len(s) != 2 {
		fmt.Println("Missing value")
		return
	}
	key := strings.TrimSpace(s[0])
	value := strings.TrimSpace(s[1])
	fmt.Printf("Set key = %v value = %v\n", key, value)
	KvalSet(currdb, key, value)
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
