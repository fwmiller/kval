package main

import (
	"fmt"
)

func CliCreate(args []string) {
	if len(args) != 2 {
		fmt.Println("Syntax: create|c <database>")
		return
	}
	Create(args[1])
}

func CliRemove(args []string) {
	if len(args) != 2 {
		fmt.Println("Syntax: remove|r <database>")
		return
	}
}

func CliSet(args []string) {
	fmt.Printf("Set key-value pair %v\n", args)

	if len(args) != 3 {
		fmt.Println("Syntax: set|s <key> <value>")
		return
	}
}

func CliGet(args []string) {
	fmt.Printf("Get value for key %v\n", args)

	if len(args) != 2 {
		fmt.Println("Syntax: get|g <key>")
		return
	}
}

func CliDel(args []string) {
	fmt.Printf("Delete key-value pair %v\n", args)

	if len(args) != 2 {
		fmt.Println("Syntax: del|d <key>")
		return
	}
}

func CliHelp(args []string) {
	fmt.Printf("Unknown command %v (add something helpful here)\n", args)
}
