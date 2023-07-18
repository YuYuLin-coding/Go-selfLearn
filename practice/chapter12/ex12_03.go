package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// content, err := io.ReadAll(f)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(content))

	reader := bufio.NewReaderSize(f, 10)
	for {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			break
		}
	}
}
