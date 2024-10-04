# Subdomain guesser

## What it is

This program can find subdomains in a wordlist and expose unknown/overlooked and possibly vulnerable subdomains.

## How it works

Using the [github.com/miekg/dns](github.com/miekg/dns) package in combination with go routines, this program can work it's way through a wordlist very fast and check for any subdomains. The user can also specify how many workers should be used in parallel.

## Demo 

Run with all arguments: 

```
go run main.go -domain=example.com -wordlist=domains.txt -c=1000 -server=8.8.8.8
```

Default arguments for `c` = 100 and `server` = '8.8.8.8'