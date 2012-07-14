package main

import "fmt"
import "net"

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        // handle error
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            // handle error
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    fmt.Println("success");
}
