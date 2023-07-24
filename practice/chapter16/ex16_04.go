package main

import (
	"fmt"
)

// func push(from, to int, out chan int) {
// 	for i := from; i <= to; i++ {
// 		out <- i
// 		time.Sleep(time.Microsecond)
// 	}
// }

func push(from, to int, in chan bool, out chan int) {
	for i := from; i <= to; i++ {
		<-in
		out <- i
	}
}

func main() {
	s1 := 0
	// ch := make(chan int, 100)

	// go push(1, 25, ch)
	// go push(26, 50, ch)
	// go push(51, 75, ch)
	// go push(76, 100, ch)

	// for c := 0; c < 100; c++ {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	s1 += i

	// }

	out := make(chan int, 100)
	in := make(chan bool, 100)

	go push(1, 25, in, out)
	go push(26, 50, in, out)
	go push(51, 75, in, out)
	go push(76, 100, in, out)

	for c := 0; c < 100; c++ {
		in <- true
		i := <-out
		fmt.Println(i)
		s1 += i
	}

	fmt.Println(s1)
}
