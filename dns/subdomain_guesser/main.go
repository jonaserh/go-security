package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/miekg/dns"
)

func usageAndExit() {
	log.Fatalln("usage: subdomain_guesser <domain> <wordlist> [-c=<worker count>] [-server=<dns server to use>]")
}

type empty struct{}

type result struct {
	IPAddress string
	Hostname  string
}

func lookupA(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return ips, err
	}
	if len(in.Answer) < 1 {
		return ips, errors.New("no answer")
	}

	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}

	return ips, nil
}

func lookupCNAME(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var fqdns []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return fqdns, err
	}
	if len(in.Answer) < 1 {
		return fqdns, errors.New("no answer")
	}
	for _, answer := range in.Answer {
		if c, ok := answer.(*dns.CNAME); ok {
			fqdns = append(fqdns, c.Target)
		}
	}
	return fqdns, nil
}

func lookup(fqdn, serverAddr string) []result {
	var results []result
	var cfqdn = fqdn
	for {
		cnames, err := lookupCNAME(cfqdn, serverAddr)
		if err == nil && len(cnames) > 0 {
			cfqdn = cnames[0]
			continue
		}
		ips, err := lookupA(cfqdn, serverAddr)
		if err != nil {
			break
		}
		for _, ip := range ips {
			results = append(results, result{IPAddress: ip, Hostname: fqdn})
		}
		break
	}
	return results
}

func worker(id int, tracker chan empty, fqdns chan string, gather chan []result, serverAddr string) {
	log.Printf("Worker #%d started\n", id)
	idx := 0
	hits := 0
	for fqdn := range fqdns {
		idx++
		results := lookup(fqdn, serverAddr)
		hasResult := len(results) > 0
		if hasResult {
			gather <- results
			hits++
		}

		if ((idx) % 10) == 0 {
			log.Printf("[Worker #%d]: %d / %d\n", id, hits, idx)
		}

	}

	log.Printf("[Worker #%d]: finished with %d / %d\n", id, hits, idx)

	var e empty
	tracker <- e
}

func main() {

	if len(os.Args) < 3 {
		usageAndExit()
	}

	var (
		domain        = flag.String("domain", "", "The domain")
		wordlist      = flag.String("wordlist", "", "The wordlist to use")
		flWorkerCount = flag.Int("c", 100, "The amount of workers to use")
		flServerAddr  = flag.String("server", "8.8.8.8:53", "The DNS server to use")
	)

	flag.Parse()

	if *domain == "" || *wordlist == "" {
		log.Println("domain and wordlist are required")
		usageAndExit()
	}

	log.Println(*flWorkerCount, *flServerAddr)

	var results []result
	fqdns := make(chan string, *flWorkerCount)
	gather := make(chan []result)
	tracker := make(chan empty)

	fh, err := os.Open(*wordlist)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	for i := 0; i < *flWorkerCount; i++ {
		go worker(i, tracker, fqdns, gather, *flServerAddr)
	}

	go func() {
		for r := range gather {
			log.Println("found results")
			results = append(results, r...)
		}
		var e empty
		tracker <- e
	}()

	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			log.Printf("Error scanning line: %v\n", err.Error())
			continue
		}
		fqdns <- fmt.Sprintf("%s.%s", scanner.Text(), *domain)
	}

	close(fqdns)
	for i := 0; i < *flWorkerCount; i++ {
		fmt.Println("waiting for worker to complete...")
		<-tracker
	}
	close(gather)
	<-tracker

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
	fmt.Fprintf(w, "Hostname\tIP address\n")
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\n", r.Hostname, r.IPAddress)
	}

	fmt.Println("\nResults:")
	w.Flush()
}
