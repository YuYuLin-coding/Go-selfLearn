package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second * 2)
	end := time.Now()
	duration1 := end.Sub(start)
	duration2 := time.Since(start)

	fmt.Println("Duration1:", duration1)
	fmt.Println("Duration2:", duration2)
	if duration1 < time.Duration(time.Millisecond*2500) {
		fmt.Println("In time")
	} else {
		fmt.Println("Exceed time")
	}
}
