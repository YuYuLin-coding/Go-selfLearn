package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
	}
	users["073"] = "Tracy"
	return users
}

func getUser(id string) (string, bool) {
	users := getUsers()
	user, exits := users[id]
	return user, exits
}

var users = getUsers()

func deleteUser(id string) {
	delete(users, id)
}

func main() {
	fmt.Println("Users:", getUsers())
	if len(os.Args) < 2 {
		fmt.Println("未傳入使用者 ID")
		os.Exit(1)
	}
	userID := os.Args[1]
	name, exits := getUser(userID)
	if !exits {
		fmt.Printf("查無傳入的使用者 ID (%v).\n使用者列表:\n", userID)
		for key, value := range getUsers() {
			fmt.Println("使用者 ID:", key, "名字:", value)
		}
		os.Exit(1)
	}
	fmt.Println("查得名字:", name)

	deleteUser(userID)
	fmt.Println("使用者列表:", users)
}
