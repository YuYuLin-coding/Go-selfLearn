package main

import (
	"fmt"
)

func greeting(fname, lname string) string {
	return fmt.Sprintln("哈囉:", fname, lname)
}

func main() {
	fname := "Edward"
	lname := "Scissorhands"

	fmt.Println("哈囉:", fname, lname)
	fmt.Printf("哈囉: %v %v\n", fname, lname)
	fmt.Print(greeting(fname, lname))

	fname = "Joe"
	gpa := 3.75
	hasJob := true
	age := 24
	hourlyWage := 45.53

	fmt.Printf("%s 的GPA: %f\n", fname, gpa)
	fmt.Printf("有工作: %t\n", hasJob)
	fmt.Printf("年齡: %d, 時薪: %v\n", age, hourlyWage)

	for i := 1; i <= 255; i++ {
		fmt.Printf("10進位: %3.d | ", i)
		fmt.Printf("2進位: %8.b | ", i)
		fmt.Printf("16進位: %2.x\n ", i)
	}
}
