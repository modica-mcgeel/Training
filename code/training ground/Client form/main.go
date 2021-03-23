package main

import (
	"clientForms/models"
	"fmt"
)

func main() {
	clientName := "SMSGLobal"
	models.AddClient(clientName)
	fmt.Println(clientName)
}
