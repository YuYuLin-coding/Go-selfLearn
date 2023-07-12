package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
	name string
	age  int
}

// func main() {
// 	c := cat{}
// 	fmt.Println(c.Speak())
// 	c.Greeting()
// }

func (c *cat) Speak() string {
	return "Purrrr Meow"
}

func (c *cat) Greeting() {
	fmt.Println("Meow, mmmmmeeeow")
}

func (c *cat) String() string {
	return fmt.Sprintf("%v (%v years old)", c.name, c.age)
}

func main() {
	c := &cat{name: "Oreo", age: 9}
	fmt.Println(c.Speak())
	fmt.Println(c)
}
