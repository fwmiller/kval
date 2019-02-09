package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fwmiller/kval/internal/cli"
	"github.com/fwmiller/kval/internal/kval"
)

func main() {
	fmt.Println("kval (C) Frank W Miller")

	go startTcpServer()

	db, err := kval.New()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	cli := cli.New(db)
	cli.Loop()
}

const (
	ConnHost = "localhost"
	ConnPort = "6380"
	ConnType = "tcp"
)

func startTcpServer() {
	l, err := net.Listen(ConnType, ConnHost + ":" + ConnPort)
	if err != nil {
		fmt.Println("Listen failed (", err.Error, ")")
		return
	}
	defer l.Close()

	fmt.Println("Listening on port " + ConnPort)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Accept failed (", err.Error, ")")
			continue
		}
		go handleTcpConnection(c)
	}
}

func handleTcpConnection(c net.Conn) {
	buf := make([]byte, 1024)
	_, err := c.Read(buf)
	if (err != nil) {
		fmt.Println("Read failed (", err.Error, ")")
	} else {
		c.Write(buf)
	}
	c.Close()
}
