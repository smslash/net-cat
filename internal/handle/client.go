package handle

import (
	"fmt"
	"net"
	"sync"

	"github.com/smslash/net-cat/internal/entity"
	"github.com/smslash/net-cat/internal/pkg"
)

var mu sync.Mutex

func Client(conn net.Conn) error {
	if err := pkg.Greet(conn); err != nil {
		return err
	}

	mu.Lock()
	for {
		name := pkg.GetName(conn)
		if name == "" {
			continue
		}
		entity.Users[name] = conn
		break
	}
	mu.Unlock()

	for i := range entity.Users {
		fmt.Printf("Name: %s\n", i)
	}
	return nil
}
