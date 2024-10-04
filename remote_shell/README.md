# Remote shell

## What it is

This folder contains two programs; A client which will connect to a server and can execute arbitrary commands on the server-side and read the results


## How it works

Given the address of a running server, the client connects to it via TCP. Using pipes and the `io.Copy` function, input and output streams are redirected on the TCP connection, so the client can send commands and receive their output.

## Demo

1. Run the server, specifying the port 

    ```
    go run server.go 8081
    ```

2. Run the client to connect to it:

    ```
    go run client.go 127.0.0.1:8081
    ```

3. You can now run commands on the client an see the results:

    ```sh
    ls # listing of the server folder
    go.mod    server.go
    ```