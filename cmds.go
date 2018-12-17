package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var dbNameCheck = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`).MatchString

func CliDb(args string) {
	/* Check for valid dbname */
	dbname := strings.TrimSpace(args)
	if !dbNameCheck(dbname) {
		fmt.Printf("Illegal characters in %v\n", dbname)
		return
	}
	dbpath := kvaldir + "/" + dbname

	/* Check whether dbpath exists */
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		fmt.Printf("Database %v does not exist\n", dbpath)
		return
	}
	/* Check whether dbpath is a directory */
	fd, err := os.Open(dbpath)
	if err != nil {
		fmt.Printf("Open database %v failed\n", dbpath)
		return
	}
	stat, err := fd.Stat()
	if err != nil {
		fmt.Printf("Stat database %v failed\n", dbpath)
		return
	}
	if !stat.IsDir() {
		fmt.Printf("Database %v is not a directory\n", dbpath)
		return
	}
	currdb = dbname
}

func CliCreate(args string) {
	/* Check for valid dbname */
	dbname := strings.TrimSpace(args)
	if !dbNameCheck(dbname) {
		fmt.Printf("Illegal characters in %v\n", dbname)
		return
	}
	dbpath := kvaldir + "/" + dbname

	/* Create new database */
	Create(dbpath)
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
