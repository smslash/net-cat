package pkg

import (
	"fmt"
	"net"
	"os"
)

func Greet(conn net.Conn) error {
	data, err := os.ReadFile("config/welcome.txt")
	if err != nil {
		return err
	}
	fmt.Fprint(conn, string(data))
	return nil
}
