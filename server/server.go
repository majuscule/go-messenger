package main

import (
    "fmt"
    "net"
    "io"
    "os"
)

const (
    OUR_PORT = "8090"
)

func main() {

    ln, err := net.Listen("tcp", ":" + OUR_PORT)

    if err != nil {
        fmt.Println("[ERROR] Failed to listen on port " + OUR_PORT + ".")
    }

    fmt.Println("Listening on port " + OUR_PORT + "...")

    for {
        conn, err := ln.Accept()

        if err != nil {
            fmt.Println("[ERROR] Connection failed.")

            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    fmt.Println("[SUCCESS] Connection received.")

    // message  := conn.Read()

    fmt.Println(io.Copy(os.Stdout, conn))

    go dispatchMessage("message")
}

func dispatchMessage(message string) {
    fmt.Println("Message dispatch called.")
    conn, err := net.Dial("tcp", "localhost:8091")

    if err != nil {
        fmt.Println("[ERROR] Failed to connect to client")
    }

    fmt.Fprintf(conn, message)
}
