package main

import "os"
import "net"

func main() {
  port := "8989"
  if len(os.Args) > 2 {
    return
  } else if len(os.Args) == 2 {
    port = os.Args[1]
  }
  listen, err := net.Listen("tcp", "localhost:"+port)
}
