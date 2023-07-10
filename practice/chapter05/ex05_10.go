package main

import "fmt"

func main() {
	defer done()
	fmt.Println("main() 開始")
	fmt.Println("main() 結束")

	defer func() {
		fmt.Println("1 claim")
	}()

	defer func() {
		fmt.Println("2 claim")
	}()
	defer func() {
		fmt.Println("3 claim")
	}()

	f1 := func() {
		fmt.Println("f1 start")
	}
	f2 := func() {
		fmt.Println("f2 end")
	}

	f1()
	f2()
	fmt.Println("main() end")
}

func done() {
	fmt.Println("換我結束了!")
}
