package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	pwR := []rune(pw)
	if len(pwR) < 8 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("") {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
	if passwordChecker("This!I5A") {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}

}
