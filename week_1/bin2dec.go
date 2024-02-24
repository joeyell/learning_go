// bin2dec.go
// 8 digit binary to decimal converter

package week_1

import (
	"fmt"
	"math"
	"strconv"
)

func Bin2dec() {

	// Display output in the next line
	fmt.Println("Enter an eight digit binary number: ")
	// var then variable name then variable type
	var first string
	// Taking input from user
	fmt.Scanln(&first)

	for {
		if !lengthCheck(first) || !numberCheck(first) {
			fmt.Println("Digits must be binary and 8 characters long, try again: ")
			fmt.Scanln(&first)
		} else {
			fmt.Println(binaryBase10Convert(first))
			break
		}
	}
}

func binaryBase10Convert(binary string) float64 {

	var exp float64 = 7
	var decimal float64 = 0

	for j := 0; j <= 7; j++ {
		i, _ := strconv.ParseFloat(string(binary[j]), 64)

		//fmt.Println(i, exp)
		result := math.Pow(2, exp)
		decimal += i * result

		exp -= 1
	}

	return decimal
}

func lengthCheck(binary string) bool {
	var returnValue bool
	if len(binary) == 8 {
		returnValue = true
	} else {
		returnValue = false
	}
	return returnValue
}

func numberCheck(binary string) bool {
	var sum int = 0
	for i := 1; i < 7; i++ {
		checkString := string(binary[i])
		if checkString != "1" && checkString != "0" {
			return false
		}
		sum += i
	}
	return true
}
