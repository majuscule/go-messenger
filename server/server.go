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

func handleConnection(conn net.Conn) {
    fmt.Println(io.Copy(os.Stdout, conn))
}
