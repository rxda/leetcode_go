package main

import "fmt"

type User interface {
	Name() string
	SetName(name string)
}

type Admin struct {
	name string
}

func (a *Admin) Name() string {
	return a.name
}

func (a *Admin) SetName(name string) {
	a.name = name
}

func main() {
	var user1 User
	user1 = &Admin{name:"user1"}

	fmt.Printf("User1's name: %s\n", user1.Name())

	//var user2 User
	//user2 = user1
	//user2.SetName("user2")

	var user2 User
	padmin :=user1.(*Admin)

	fmt.Printf("User2's name: %s\n", user2.Name()) // The name will be changed as "user2"
	fmt.Printf("User1's name: %s\n", user1.Name())  // The name will be changed as "user2" too, How to make the user1 name does not change？
}
