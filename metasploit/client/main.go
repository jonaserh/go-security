package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jonaserh/go-security/metasploit/rpc"
)

func main() {
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"

	if host == "" || pass == "" {
		log.Fatalln("Missing required env variable MSFHOST or MSFPASS")
	}

	msf, err := rpc.New(host, user, pass)
	if err != nil {
		log.Panicln(err)
	}

	defer msf.Logout()

	sessions, err := msf.ListSessions()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Sessions:")
	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}
}
