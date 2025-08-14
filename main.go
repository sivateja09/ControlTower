package main

import (
	"fmt"
)

func main() {
	var role string
	fmt.Print("Run as (server/client): ")
	fmt.Scanln(&role)

	if role == "server" {
		startServer()
	} else {
		startClient()
	}
}
