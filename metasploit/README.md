# Metasploit interaction Demo

## What it is

Metasploit is a popular framwork used by penetration testers for reconnaisance, exploitation, command and control, persistence, payload creation, privelege escalation and more.
It also exposes as `remote procedure protocol (RPC)` API to allow remote interaction with it's functionality.

This demo includes a small client that interacts with the Metasploit RPC API.

## How it works

Using the package [gopkg.in/vmihailenco/msgpack.v2](gopkg.in/vmihailenco/msgpack.v2) and some custom data types, this program will list exploit sessions.

## Demo

1. Make sure you have the metasploit community edition installed

2. Open the metasploit console and then set the server host and password for RPC:

    ```sh
    $ msfconsole

    msf > load msgrpc Pass=password ServerHost=10.0.1.6
    ```

3. Set the same parameters as env variables for the go program

    ```sh
    export MSFHOST=10.0.1.6:55552
    export MSFPASS=password
    ```

4. Run the program

    ```
    go run main.go
    ```