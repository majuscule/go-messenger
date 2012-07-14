package main

import (
    "fmt"
    "net"
)

const (
    SERVER_ADDR = "localhost"
    SERVER_PORT = "8090"
)

func main() {

    conn, _ := net.Dial("tcp", SERVER_ADDR + ":" + SERVER_PORT)
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
}
