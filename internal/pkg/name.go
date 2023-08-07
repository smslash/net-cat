package pkg

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/smslash/net-cat/internal/entity"
)

func GetName(conn net.Conn) string {
	fmt.Fprint(conn, "[ENTER YOUR NAME]: ")
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Error during reading: %s\n", err)
		return ""
	}

	name = strings.TrimSpace(name)
	if len(name) == 0 {
		fmt.Fprintf(conn, "Invalid name. Try again...\n")
		return ""
	}

	if _, ok := entity.Users[name]; ok {
		fmt.Fprintf(conn, "User with this name already exists. Try again...\n")
		return ""
	}

	return name
}
