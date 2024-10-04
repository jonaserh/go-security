package main

import (
	"fmt"
	"net"
	"os"
	"sort"
)

func worker(ports, results chan int, verbose bool) {
	for p := range ports {
		if verbose {
			fmt.Printf("Scanning %d...\n", p)
		}
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			if verbose {
				fmt.Printf("No connection on port %d\n", p)
			}
			results <- 0
			continue
		}
		if verbose {
			fmt.Printf("Got connection on port %d\n", p)
		}
		conn.Close()
		results <- p
	}
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 || len(argsWithoutProg) > 2 {
		fmt.Fprint(os.Stderr, "Usage: ./go_port_scanner <port-range | port1,port2 | port> [-v (verbose)]")
		os.Exit(1)
	}

	portsToScan, err := Parse(argsWithoutProg[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse port or port range: %s", argsWithoutProg[0])
		os.Exit(1)
	}

	verbose := false

	if len(argsWithoutProg) == 2 && argsWithoutProg[1] == "-v" {
		verbose = true
	}

	ports := make(chan int, len(portsToScan))
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, verbose)
	}

	go func() {
		for _, p := range portsToScan {
			ports <- p
		}
	}()

	for i := 0; i < len(portsToScan); i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	fmt.Print("\nOpen ports:\n")
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
