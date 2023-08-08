package pkg

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/smslash/net-cat/internal/entity"
)

func GetName(conn net.Conn) string {
	fmt.Fprint(conn, "[ENTER YOUR NAME]: ")
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	name = strings.TrimSpace(name)
	if len(name) == 0 {
		fmt.Fprintf(conn, "Invalid name. Try again...\n")
		return GetName(conn)
	}

	if _, ok := entity.Users[name]; ok {
		fmt.Fprintf(conn, "User with this name already exists. Try again...\n")
		return GetName(conn)
	}

	return name
}
