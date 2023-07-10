package main

import "fmt"

type calc func(int, int) string

func main() {
	calculator(add, 5, 6)
	calculator(substract, 10, 5)
}

func add(i, j int) int {
	result := i + j
	fmt.Printf("%d + %d = %d\n", i, j, result)
	return i + j
}

func substract(i, j int) int {
	return i - j
}

// func calculator(f calc, i, j int) {
// 	fmt.Println(f(i, j))
// }

func calculator(f func(int, int) int, i, j int) {
	fmt.Println(f(i, j))
}
