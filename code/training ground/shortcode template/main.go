package main

import (
	"fmt"
)

func main() {
	//controller.AddUser()

	port := 3000
	port, err := StartServer(port)
	fmt.Println("Port", port, "has started successfully", err)
}

func StartServer(port int) (int, error) {
	fmt.Println("THe server has started on port", port)
	return port, nil
}
