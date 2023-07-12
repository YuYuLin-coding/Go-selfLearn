package main

import (
	"fmt"
	// "strconv"
)

// func main() {
// 	var s interface{} = "42"
// 	v := s.(string)
// 	fmt.Println(strconv.Atoi(v))
// }

// func main() {
// 	var str interface{} = 42
// 	if s, ok := str.(string); ok {
// 		fmt.Println(strconv.Atoi((s)))
// 	} else {
// 		fmt.Println("Type assertion failed")
// 	}
// }

type cat struct {
	name string
}

func main() {
	i := []interface{}{42, "The book club", true, cat{name: "oreo"}}
	typeExample(i)
}

func typeExample(i []interface{}) {
	for _, x := range i {
		switch v := x.(type) {
		case int:
			fmt.Printf("%v is int\n", v)
		case string:
			fmt.Printf("%v is a string\n", v)
		case bool:
			fmt.Printf("%v is a bool\n", v)
		default:
			fmt.Printf("%T is unknown type\n", v)
		}
	}
}
