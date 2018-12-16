package main

import (
	"fmt"
	"os"
)

func Create(database string) {
	dbdir := kvaldir + "/" + database

	/* Create database directory */
	fmt.Printf("Create database %v", dbdir)
	err := os.Mkdir(dbdir, 0777)
	if err != nil {
		fmt.Printf(" failed\n")
		return
	}
	fmt.Printf("\n")
}
