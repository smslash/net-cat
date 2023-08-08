package handle

import (
	"fmt"
	"os"
	"time"

	"github.com/smslash/net-cat/internal/entity"
	"github.com/smslash/net-cat/internal/pkg"
)

var (
	Join    = make(chan entity.User)
	Message = make(chan entity.Message)
	Leave   = make(chan entity.User)
)

func Broadcast() {
	for {
		select {
		case ans := <-Join:
			mu.Lock()
			message := fmt.Sprintf("\n%s has joined our chat...\n", ans.Name)
			if len(entity.Users) != 1 {
				file, err := os.OpenFile("config/history.txt", os.O_APPEND|os.O_WRONLY, 0o644)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}

				_, err = file.WriteString(message[1:])
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
			}

			for i, v := range entity.Users {
				if i != ans.Name {
					fmt.Fprintln(v, message[:len(message)-1])
				}
				pkg.PrintTime(v, i)
			}
			mu.Unlock()
		case ans := <-Message:
			mu.Lock()
			time := time.Now().Format("2006-01-02 15:04:05")
			message := fmt.Sprintf("\n[%s][%s]:%s\n", time, ans.Author, ans.Text)
			file, err := os.OpenFile("config/history.txt", os.O_APPEND|os.O_WRONLY, 0o644)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			_, err = file.WriteString(message[1:])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			for i, v := range entity.Users {
				if i != ans.Author {
					fmt.Fprintln(v, message[:len(message)-1])
				}
				pkg.PrintTime(v, i)
			}
			mu.Unlock()
		case ans := <-Leave:
			mu.Lock()
			message := fmt.Sprintf("\n%s has left our chat...\n", ans.Name)
			file, err := os.OpenFile("config/history.txt", os.O_APPEND|os.O_WRONLY, 0o644)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			_, err = file.WriteString(message[1:])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			for i, v := range entity.Users {
				if i != ans.Name {
					fmt.Fprintln(v, message[:len(message)-1])
				}
				pkg.PrintTime(v, i)
			}
			mu.Unlock()
		}
	}
}
