package main

import (
    "fmt"
    "net"
    "flag"
)

const (
    SERVER_ADDR = "localhost"
    SERVER_PORT = "8090"
)


func main() {
    var message string
    parseArguments(&message)

    conn, _ := net.Dial("tcp", SERVER_ADDR + ":" + SERVER_PORT)
    fmt.Fprintf(conn, message)

    return
}


func parseArguments(message *string) {
    flag.StringVar(message, "message", "", "tell a ducktern")
    flag.Parse()
}
