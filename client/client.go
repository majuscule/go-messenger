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
    conn, _ := net.Dial("tcp", SERVER)
    fmt.Fprintf(conn, message)
    status, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Println(status)
}

func parseArguments(message *string) {
    flag.StringVar(message, "message", "", "tell a ducktern")
    flag.Parse()
}
