package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jonaserh/go-security/shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}

	apiKey := os.Getenv("SHODAN_API_KEY")

	client := shodan.New(apiKey)

	info, err := client.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Query credits: %d\nScan credits: %d\n", info.QueryCredits, info.ScanCredits)

	hostSearch, err := client.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d", host.IPString, host.Port)
	}

}
