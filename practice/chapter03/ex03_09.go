package main

import "fmt"

func main() {
	var message *string

	if message == nil {
		fmt.Println("錯誤")
		return
	}
	fmt.Println(&message)
}
