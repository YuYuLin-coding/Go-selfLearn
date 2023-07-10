package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	csvHdrCol(hdr)

	hdr2 := []string{"Employee", "Empid", "Hours worked", "Address", "Manager", "Hourly rate"}
	csvHdrCol(hdr2)
}

func csvHdrCol(header []string) {
	csvHeadersColumnIndex := make(map[int]string)
	for i, v := range header {
		switch v := strings.ToLower(strings.TrimSpace(v)); v {
		case "employee":
			csvHeadersColumnIndex[i] = v
		case "hours worked":
			csvHeadersColumnIndex[i] = v
		case "hourly rate":
			csvHeadersColumnIndex[i] = v
		}
	}
	fmt.Println(csvHeadersColumnIndex)
}
