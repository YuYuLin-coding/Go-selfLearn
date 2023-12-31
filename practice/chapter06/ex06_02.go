package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate   = errors.New("無效的時薪")
	ErrHourlyWorked = errors.New("無效的一週工時")
)

func main() {
	pay, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	pay, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	pay, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 {
		// return 0, ErrHourlyRate
		// return 0, fmt.Errorf("無效的時薪: %d", hourlyRate)
		// return 0, fmt.Errorf("payDay 錯誤: %s", ErrHourlyRate.Error())
		return 0, fmt.Errorf("payDay 錯誤: %w", ErrHourlyRate)

	}

	if hoursWorked < 0 || hoursWorked > 80 {
		// return 0, ErrHourlyWorked
		// return 0, fmt.Errorf("無效的一週工時: %d", hoursWorked)
		// return 0, fmt.Errorf("payDay 錯誤: %s", ErrHourlyWorked.Error())
		return 0, fmt.Errorf("payDay 錯誤: %w", ErrHourlyWorked)

	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		hoursRegular := hoursWorked - hoursOver
		return hoursRegular*hourlyRate + hoursOver*hourlyRate*2, nil
	}
	return hoursWorked * hourlyRate, nil
}
