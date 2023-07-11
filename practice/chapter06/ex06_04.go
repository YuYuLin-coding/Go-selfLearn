package main

import (
	"errors"
	"fmt"
)

func main() {
	a()
	fmt.Println("這一行現在會印出")
}

func a() {
	b("goodbye")
	fmt.Println("返回 a()")

}

func b(msg string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("b() 發生錯誤:", r)
		}
	}()
	if msg == "goodbye" {
		panic(errors.New("出事了"))
	}
	fmt.Print(msg)
}
