//trunc.go

package week_2

import (
	"fmt"
	"strconv"
	"strings"
)

func Trunc() {
	//Display output in the next line
	fmt.Println("Input a floating point number: ")

	// var then variable name then variable type
	var userInput float64

	// Taking input from user
	fmt.Scanln(&userInput)

	// String conversion and manipulation
	floatString := strconv.FormatFloat(userInput, 'f', -1, 64)
	decimalFinder := strings.Index(floatString, ".")

	// Make sure a decimal is present
	if decimalFinder == -1 {
		fmt.Println(floatString)
	} else {
		fmt.Println(floatString[:decimalFinder])
	}

}
