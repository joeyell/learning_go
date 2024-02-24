// math.go

package week_1

import (
	"fmt"
	"math"
)

func MathEquation() {
	a, vo, so, time := getUserValues()
	fn := GenDisplaceFn(a, vo, so)
	fmt.Printf("Your displacement after %f seconds is %f\n", math.Round(time), fn(time))
}

// Function that gets user values and returns each

func getUserValues() (float64, float64, float64, float64) {

	// delcare int values
	var a, vo, so, time float64

	// get user input
	fmt.Println("Enter acceleration value: ")
	fmt.Scan(&a)
	fmt.Println("Enter initial velocity value: ")
	fmt.Scan(&vo)
	fmt.Println("Enter initial displacement value: ")
	fmt.Scan(&so)
	fmt.Println("Enter initial time value: ")
	fmt.Scan(&time)

	return a, vo, so, time
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	fx := func(time float64) float64 {
		resultSo := (0.5 * a * (math.Pow(time, 2))) + (vo * time) + (so)
		return resultSo
	}
	return fx
}
