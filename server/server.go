package main

import (
    "fmt"
    "net"
)

func main() {
    ln, err := net.Listen("tcp", ":8090")

    if err ~= nil {
        fmt.Println("Listen error")
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Connection error.")
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    fmt.Println("success");
}
