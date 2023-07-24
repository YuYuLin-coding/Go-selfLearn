// package main

// import "fmt"

// func greet(ch chan string) {
// 	ch <- "Hello"
// }

// func main() {
// 	// ch := make(chan int, 1)
// 	ch := make(chan string)
// 	// ch <- 1
// 	// i := <-ch
// 	// fmt.Println(i)
// 	go greet(ch)
// 	fmt.Println(<-ch)

// }

package main

import "fmt"

func greet(ch chan string) {
	msg := <-ch
	ch <- fmt.Sprintf("msg: %s", msg)
	ch <- "Hello David"
}

func main() {
	ch := make(chan string)
	go greet(ch)

	ch <- "Hello John"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
