package main

import (
	"fmt"
	"time"
)

func main() {
	// start := time.Now()
	// fmt.Println("start:", start)
	// fmt.Println("Running...")
	// time.Sleep(2 * time.Second)
	// end := time.Now()
	// fmt.Println("end:", end)

	now := time.Now()
	day := now.Weekday()
	hour := now.Hour()
	fmt.Println("Day:", day, "/ hour:", hour)
	if day.String() == "Monday" && (hour >= 0 && hour < 2) {
		fmt.Println("full test")
	} else {
		fmt.Println("simple test")
	}
}
