package main

import (
	"fmt"
	"os"
)

func Create(database string) {
	dbpath := kvaldir + "/" + database

	/* Create database directory */
	fmt.Printf("Create database %v", dbpath)
	err := os.Mkdir(dbpath, 0777)
	if err != nil {
		fmt.Printf(" failed")
	}
	fmt.Printf("\n")
}
