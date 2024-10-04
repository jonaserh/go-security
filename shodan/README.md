# Shodan interaction

## What it is

This demo shows how to use the shodan go package to interact with [shodan.io](https://shodan.io), the `search engine for the internet of everything`.

## How it works

Using [github.com/jonaserhart/security/go/shodan/shodan](github.com/jonaserhart/security/go/shodan/shodan) and custom defined data types in the [shodan](./shodan/) folder, this program will perform a host search.

## Demo

1. Get an API Key from shodan by signing up on [shodan.io](https://shodan.io)

2. Run the program like this:
    
    ```
    SHODAN_API_KEY=<your key> go run main.go tomcat
    Query Credits: 100
    Scan Credits: 100

        185.23.138.141   8081
       218.103.124.239   8080
       ...
    ```