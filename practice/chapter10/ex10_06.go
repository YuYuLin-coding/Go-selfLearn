package main

import (
	"fmt"
	"time"
)

func main() {
	// date := time.Date(2050, 12, 31, 0, 0, 0, 0, time.Local)

	// fmt.Println("Equal:", time.Now().Equal(date))
	// fmt.Println("Before:", time.Now().Before(date))
	// fmt.Println("After:", time.Now().After(date))

	now := time.Now()
	duration1 := time.Duration(time.Second * 360)
	duration2 := time.Duration(time.Hour*1 + time.Minute*30)
	fmt.Println("Dur1:", duration1.Nanoseconds(), "ns")
	fmt.Println("Dur2:", duration2.Nanoseconds(), "ns")

	date1 := now.Add(duration1)
	date2 := now.Add(duration2)

	fmt.Println("Now:", now)
	fmt.Println("Date1:", date1)
	fmt.Println("Date2:", date2)
}
