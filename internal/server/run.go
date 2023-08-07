package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/smslash/net-cat/internal/entity"
	"github.com/smslash/net-cat/internal/handle"
)

func Run(protocol, address string) error {
	listener, err := net.Listen(protocol, address)
	if err != nil {
		return err
	}
	c := make(chan bool)

	go func() {
		sigInt := make(chan os.Signal, 1)
		signal.Notify(sigInt, syscall.SIGINT, syscall.SIGTERM)
		<-sigInt

		if err = listener.Close(); err != nil {
			fmt.Printf("Error during closing server: %s\n", err)
			c <- false
		}

		fmt.Println("\nServer is closed")
		c <- true
	}()

	fmt.Printf("The server is running on address %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}

		if len(entity.Users) >= 10 {
			fmt.Fprintln(conn, "Apologize, chat is full. Come back later...")
			if err = conn.Close(); err != nil {
				fmt.Printf("Error during closing conn: %s\n", err)
				break
			}
		}

		go func() {
			if err = handle.Client(conn); err != nil {
				fmt.Printf("Error during handling conn: %s\n", err)
				os.Exit(1)
			}
		}()
	}

	<-c
	return nil
}
