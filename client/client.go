package main

import (
    "fmt"
    "net"
    "bufio"
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
    status, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Println(status)
}

func parseArguments(message *string) {
    flag.StringVar(message, "message", "", "tell a ducktern")
    flag.Parse()
}
