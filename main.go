package main

func main() {
  port := "8989"
  if len(os.Args) > 2 {
    return
  } else if len(os.Args) == 2 {
    port = os.Args[1]
  }
}
