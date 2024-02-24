// main.go

package week_2

import (
	"fmt"
	"sort"
	"strconv"
)

func Slice() {

	// Initialse a slice of length 3
	sli_3 := make([]int, 3)

	// Infinite loop
	for {
		// Get user input
		var userInput string
		fmt.Printf("Please enter an integer, type 'X' to exit: ")
		fmt.Scan(&userInput)

		// Check if input was X to exit loop
		if userInput == "X" {
			fmt.Println("Exiting")
			break
		}

		// Convert user string input to an integer
		userInt, err := strconv.Atoi(userInput)

		// Handle error
		if err != nil {
			panic("Only integers, dummy!")
		}

		// Append converted int to slice
		sli_3 = append(sli_3, userInt)
		//Sort slice
		sort.Ints(sli_3)

		// Print slice
		fmt.Println(sli_3)
	}
}
