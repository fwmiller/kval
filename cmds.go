package main

import (
	"fmt"
	"regexp"
	"strings"
)

var dbNameCheck = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`).MatchString

func CliDb(args string) {
	/* Check for valid dbname */
	n := strings.TrimSpace(args)
	if !dbNameCheck(n) {
		fmt.Printf("Illegal characters in %s\n", n)
		return
	}
	/* Check for valid database */
	name, err := db.IsDB(n)
	if err != nil {
		fmt.Println(err)
		return
	}

	if name != "" {
		currdb = name
	}
}

func CliCreate(args string) {
	/* Check for valid dbname */
	dbname := strings.TrimSpace(args)
	if !dbNameCheck(dbname) {
		fmt.Printf("Illegal characters in %s\n", dbname)
		return
	}
	/* Create new database */
	err := db.Create(dbname)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Set current to new database if it was clear */
	if currdb == "" {
		currdb = dbname
	}
}

func CliRemove(args string) {
	/* Check for valid dbname */
	dbname := strings.TrimSpace(args)
	if !dbNameCheck(dbname) {
		fmt.Printf("Illegal characters in %s\n", dbname)
		return
	}
	/* Remove existing database */
	err := db.Remove(dbname)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Clear current database if it was just removed */
	if dbname == currdb {
		currdb = ""
	}
}

func CliKeys() {
	if currdb == "" {
		fmt.Println("Current database not set")
		return
	}

	keys, err := db.Keys(currdb)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, k := range keys {
		fmt.Println(k)
	}
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

	/* Set key-value pair in current database */
	if err := db.Set(currdb, key, value); err != nil {
		fmt.Println(err)
	}
}

func CliGet(args string) {
	if currdb == "" {
		fmt.Println("Current database not set")
		return
	}
	key := strings.TrimSpace(args)

	/* Get value associated with key in current database */
	value, err := db.Get(currdb, key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

func CliDel(args string) {
	if currdb == "" {
		fmt.Println("Current database not set")
		return
	}
	key := strings.TrimSpace(args)

	/* Delete key-value pair in current database */
	if err := db.Del(currdb, key); err != nil {
		fmt.Println(err)
	}
}

func CliHelp() {
	fmt.Println("Help (add something useful here)")
}
