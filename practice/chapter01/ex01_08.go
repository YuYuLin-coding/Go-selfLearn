package main

import (
	"fmt"
)

var defaultOffset = 10

func main() {
	_A1_Me := ""
	fmt.Println(_A1_Me)
	offset := 5
	fmt.Println(offset)
	offset = 10
	fmt.Println(offset)
	offset1 := defaultOffset
	fmt.Println(offset1)
	offset1 = offset1 + defaultOffset
	fmt.Println(offset1)
}
