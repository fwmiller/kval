package main

import (
	"fmt"
)

func Create(args []string) {
	fmt.Printf("Create database %v\n", args)

	if len(args) != 2 {
		fmt.Println("Syntax: create|c <database>")
		return
	}
}

func Remove(args []string) {
	fmt.Printf("Remove database %v\n", args)

	if len(args) != 2 {
		fmt.Println("Syntax: remove|r <database>")
		return
	}
}

func Set(args []string) {
	fmt.Printf("Set key-value pair %v\n", args)

	if len(args) != 3 {
		fmt.Println("Syntax: set|s <key> <value>")
		return
	}
}

func Get(args []string) {
	fmt.Printf("Get value for key %v\n", args)

	if len(args) != 2 {
		fmt.Println("Syntax: get|g <key>")
		return
	}
}

func Del(args []string) {
	fmt.Printf("Delete key-value pair %v\n", args)

	if len(args) != 2 {
		fmt.Println("Syntax: del|d <key>")
		return
	}
}

func Help(args []string) {
	fmt.Printf("Unknown command %v (add something helpful here", args)
}
