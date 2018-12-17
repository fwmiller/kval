package main

import (
	"fmt"
	"os"
)

func Create(dbpath string) {
	/* Assume dbpath is a valid database path */
	fmt.Printf("Create database %v", dbpath)
	err := os.Mkdir(dbpath, 0777)
	if err != nil {
		fmt.Printf(" failed")
	}
	fmt.Printf("\n")
}
