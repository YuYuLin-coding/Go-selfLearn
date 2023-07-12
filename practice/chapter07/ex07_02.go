package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
}

type dog string

type person struct {
	name string
}

func main() {
	c := cat{}
	d := dog("")
	p := person{name: "Heather"}
	// thingSpeak(c)
	// thingSpeak(d)
	// thingSpeak(p)
	thingSpeak(c, d, p)

}

func (c cat) Speak() string {
	return "Purr Meow"
}

func (d dog) Speak() string {
	return "Woof"
}

func (p person) Speak() string {
	return "Hi, my name is " + p.name + "."
}

// func thingSpeak(s Speaker) {
// 	fmt.Println(s.Speak())
// }

func thingSpeak(speakers ...Speaker) {
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}
}
