package handle

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/smslash/net-cat/internal/entity"
	"github.com/smslash/net-cat/internal/pkg"
)

var mu sync.Mutex

func Client(conn net.Conn, mu *sync.Mutex) error {
	if err := pkg.Greet(conn); err != nil {
		return err
	}

	name := pkg.GetName(conn)
	user := entity.User{
		Name: name,
		Conn: conn,
	}
	pkg.PrintHistory(conn)

	mu.Lock()
	entity.Users[name] = conn
	Join <- user
	mu.Unlock()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if len(strings.TrimSpace(text)) == 0 {
			fmt.Fprintln(conn, "Message can not be empty")
			pkg.PrintTime(conn, name)
			continue
		}

		message := entity.Message{
			Author: name,
			Text:   text,
		}

		mu.Lock()
		Message <- message
		mu.Unlock()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error during scanning: %s\n", err)
	}

	mu.Lock()
	Leave <- user
	delete(entity.Users, user.Name)
	conn.Close()
	mu.Unlock()

	return nil
}
