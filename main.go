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
  if err != nil {
    return
  }

  ch := make(chan string)
  for {
    conn, err := listen.Accept()
    if err != nil {
      return
    }
    go live()
    fmt.Println(<-ch)
    conn.Close()
  }
}
