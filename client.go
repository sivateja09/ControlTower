package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const clientAuthKey = "mysecret123" // must match server

func startClient() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter server address (e.g. 192.168.1.100:5000): ")
	if !scanner.Scan() {
		fmt.Println("Failed to read input")
		return
	}
	serverAddress := scanner.Text()

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send auth key
	_, err = conn.Write([]byte(clientAuthKey))
	if err != nil {
		fmt.Println("Auth failed:", err)
		return
	}

	fmt.Println("Connected to server.")

	reader := bufio.NewReader(conn)
	for {
		cmdLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed.")
			return
		}
		cmdLine = strings.TrimSpace(cmdLine)
		fmt.Println("Executing:", cmdLine)

		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/C", cmdLine)
		} else {
			cmd = exec.Command("bash", "-c", cmdLine)
		}

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(string(out))
	}
}
