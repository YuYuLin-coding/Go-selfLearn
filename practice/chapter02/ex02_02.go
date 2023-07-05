package main

import (
	"fmt"
	"time"
)

func main() {
	// day := time.Monday
	// switch day {
	// case time.Monday:
	// 	fmt.Println("星期一")
	// case time.Tuesday:
	// 	fmt.Println("星期二")
	// case time.Wednesday:
	// 	fmt.Println("星期三")
	// case time.Thursday:
	// 	fmt.Println("星期四")
	// case time.Friday:
	// 	fmt.Println("星期五")
	// case time.Saturday:
	// 	fmt.Println("星期六")
	// case time.Sunday:
	// 	fmt.Println("星期日")
	// default:
	// 	fmt.Println("日期不正確")
	// }

	// =====================

	// dayBorn := time.Sunday
	// switch dayBorn {
	// case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
	// 	fmt.Println("為平日")
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("為週末")
	// default:
	// 	fmt.Println("日期錯誤")
	// }

	// =====================

	switch dayBorn := time.Sunday; {
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("為週末")
	default:
		fmt.Println("非週末")
	}
}
