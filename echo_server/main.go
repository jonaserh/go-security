package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func echo(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)

		s, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Client has been disconnected")
			break
		}

		log.Printf("Received %d bytes: %s\n", len(s), s)

		writer := bufio.NewWriter(conn)
		if _, err := writer.WriteString(s); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 || len(argsWithoutProg) > 1 {
		log.Fatal("Usage: ./go_echo_server <port>")
		os.Exit(1)
	}

	port, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		log.Fatalf("Unable to parse port from %s", argsWithoutProg[0])
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("Unable to bind port: %v\n", err)
	}

	log.Printf("Listening on port %d\n", port)

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}

}
