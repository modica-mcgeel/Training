package main

import "fmt"

type User struct {
	FName string
	LName string
	ID    int
}

var (
	user   []*User
	NextID = 1
)

func AddUser(u User) (User, error) {
	//u.FName = "Lia"
	//u.LName = "McGee"
	//u.ID = 10
	return u, nil
}

func main() {
	User.FName="Lia"

	AddUser(U)
	fmt.Println("Hello from the controller")

}
