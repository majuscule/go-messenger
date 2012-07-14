package main

import (
    "fmt"
    "net"
    "bufio"
)

const (
    SERVER = "0:8090"
)

func main() {
    conn, err := net.Dial("tcp", SERVER)
    if err != nil {
        // handle error
    }
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    status, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        // handle error
    }
    fmt.Println(status)
}
