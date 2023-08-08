package entity

import (
	"net"
)

var Users = make(map[string]net.Conn)

type User struct {
	Name string
	Conn net.Conn
}

type Message struct {
	Author string
	Text   string
}
