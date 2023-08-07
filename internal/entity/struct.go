package entity

import (
	"net"
)

var Users = make(map[string]net.Conn)

type User struct {
	name string
	conn net.Conn
}

type Message struct {
	author string
	text   string
}
