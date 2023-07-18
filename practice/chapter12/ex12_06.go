package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

const (
	firstName = iota
	lastname
	age
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := true
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if header {
			header = false
			continue
		}
		// fmt.Println(record)
		fmt.Println("First name:", record[firstName])
	}

}
