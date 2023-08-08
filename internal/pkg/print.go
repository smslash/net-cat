package pkg

import (
	"fmt"
	"net"
	"os"
	"time"
)

func PrintHistory(conn net.Conn) {
	file, err := os.ReadFile("config/history.txt")
	if err != nil {
		fmt.Println("Error during reading history.txt")
		return
	}

	fmt.Fprint(conn, string(file))
}

func PrintTime(conn net.Conn, name string) {
	time := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(conn, fmt.Sprintf("[%s][%s]:", time, name))
}
