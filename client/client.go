package main

import (
    "fmt"
    "net"
    "bufio"
    "flag"
)

const (
    SERVER = "0:8090"
)

func main() {
    var message string
    parseArguments(&message);
    conn, err := net.Dial("tcp", SERVER)
    if err != nil {
        // handle error
    }
    fmt.Fprintf(conn, message)
    status, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        // handle error
    }
    fmt.Println(status)
}

func parseArguments(message *string) {
    flag.StringVar(message, "message", "", "tell a ducktern")
    flag.Parse()
}
