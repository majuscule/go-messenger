package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

const (
    SERVER = "0:8090"
)

func main() {
    conn, err := net.Dial("tcp", SERVER)
    if err != nil {
        // handle error
    }
    fmt.Fprintf(conn, os.Args[1])
    status, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        // handle error
    }
    fmt.Println(status)
}
