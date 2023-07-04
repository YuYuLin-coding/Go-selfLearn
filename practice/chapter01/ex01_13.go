package main

import (
	"fmt"
	"time"
)

func main() {
	// var count1 *int
	// count2 := new(int)
	// countTemp := 5
	// count3 := &countTemp
	// t := &time.Timer{}

	// fmt.Printf("count1: %#v\n", count1)
	// fmt.Printf("count2: %#v\n", count2)
	// fmt.Printf("count3: %#v\n", count3)
	// fmt.Printf("time: %#v\n", t)

	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("time: %#v\n", *t)
		fmt.Printf("time: %#v\n", t.String())
	}
}