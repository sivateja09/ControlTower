package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

var clients []net.Conn
var mu sync.Mutex

const authKey = "mysecret123"

func startServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started on port 5000")

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}

			// Check auth
			buf := make([]byte, len(clientAuthKey))
			_, err = conn.Read(buf)
			if err != nil || string(buf) != clientAuthKey {
				fmt.Println("Client failed authentication:", conn.RemoteAddr())
				conn.Close()
				continue
			}

			mu.Lock()
			clients = append(clients, conn)
			mu.Unlock()
			fmt.Println("Client connected:", conn.RemoteAddr())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("cmd> ")
		if scanner.Scan() {
			cmd := scanner.Text()
			mu.Lock()
			for _, conn := range clients {
				conn.Write([]byte(cmd + "\n"))
			}
			mu.Unlock()
		}
	}
}
