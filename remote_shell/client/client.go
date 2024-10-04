package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 1 {
		log.Fatal("Usage: ./client <address:port>")
		os.Exit(1)
	}

	address := argsWithoutProg[0]

	conn, err := net.Dial("tcp", address)

	if err != nil {
		log.Fatalf("Could not establish a connection: %v\n", err)
	}

	defer conn.Close()

	go io.Copy(os.Stdout, conn)
	io.Copy(conn, os.Stdin)
}
