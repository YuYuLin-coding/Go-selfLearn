package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}
	for i := 0; i < len(username); i++ {
		fmt.Print(string(username[i]), " ")
	}

	runes := []rune(username)

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}

	//================================

	for _, v := range runes {
		fmt.Print(string(v), " ")
	}
}
