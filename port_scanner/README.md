# Port scanner

## What it is

This tool provides the functionality to scan a port, range of ports or multiple specific ports and a remote address.
The program will then print out all open ports

## Demo

The demo is hardcoded to scan `scanme.nmap.org`, so if you want to scan anything else, change it.

You can run it three different ways:

```sh
# Port range
go run main.go 3500-4851
# Single port 
go run main.go 22
# Multiple ports
go run main.go 8080,22,5432
```