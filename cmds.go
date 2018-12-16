package main

import (
	"fmt"
)

func CliCreate(args string) {
	fmt.Printf("Create args = %v\n", args)
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
