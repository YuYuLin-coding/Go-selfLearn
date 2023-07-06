package main

import "fmt"

func main() {
	comment1 := `This is the BEST
thing ever`
	comment2 := `This is the BEST\nthing ever!`
	comment3 := "This is the BEST\nthing ever!"

	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")

	// comment4 := 'In "Windows" the user directory is "C:\Users\"'
	comment5 := "In \"Windows\" the user directory is \"C:\\Users\\\""

	// fmt.Println(comment4)
	fmt.Println(comment5)
}
