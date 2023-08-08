package main

import (
	"fmt"
	"net"
	"os"

	"github.com/smslash/net-cat/internal/handle"
)

func main() {
	host := "localhost"
	protocol := "tcp"
	port := "8989"

	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	address := net.JoinHostPort(host, port)
	if err := handle.Run(protocol, address); err != nil {
		fmt.Printf("Error durig runnig server: %s\n", err)
		return
	}
}
