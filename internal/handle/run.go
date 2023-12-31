package handle

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/smslash/net-cat/internal/entity"
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

		file, err := os.OpenFile("config/history.txt", os.O_WRONLY|os.O_TRUNC, 0o644)
		if err != nil {
			fmt.Printf("Error during openning history.txt: %v\n", err)
			return
		}
		defer file.Close()

		fmt.Println("\nServer is closed")
		c <- true
	}()

	fmt.Printf("The server is running on address %s\n", address)
	go Broadcast()

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
			if err = Client(conn, &mu); err != nil {
				fmt.Printf("Error during handling conn: %s\n", err)
				os.Exit(1)
			}
		}()
	}

	<-c
	return nil
}
