//goroutine_sort.go

package week_3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Goroutine_sort() {
	printInstructions()
	userInput := getUserInput()
	w, x, y, z := splitIntSlice(userInput)
	fmt.Println("You integer slice: ", userInput)
	go bubbleSort(w)
	go bubbleSort(x)
	go bubbleSort(y)
	go bubbleSort(z)
	fmt.Println("========================")
	fmt.Println("Go routine 1 sorted: ", w)
	fmt.Println("Go routine 2 sorted: ", x)
	fmt.Println("Go routine 3 sorted: ", y)
	fmt.Println("Go routine 4 sorted: ", z)
	fmt.Println("========================")
	fmt.Println(bubbleSort(merge(w, x, y, z)))
	fmt.Println("========================")
}

func printInstructions() {
	fmt.Println("Please enter a list of intergers.")
}

func getUserInput() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter integers separated by spaces:")
	inputStr, _ := reader.ReadString('\n')
	numStr := strings.Split(strings.TrimSpace(inputStr), " ")
	var nums []int
	for _, k := range numStr {
		num, err := strconv.Atoi(k)
		if err != nil {
			panic("Error: Unable to parse the non integer value")
		}
		nums = append(nums, num)
	}
	return nums
}

func splitIntSlice(intSlice []int) ([]int, []int, []int, []int) {
	var w, x, y, z []int
	var split int = len(intSlice) / 4

	w = append(w, intSlice[:split]...)
	x = append(x, intSlice[split:split*2]...)
	y = append(y, intSlice[split*2:split*3]...)
	z = append(z, intSlice[split*3:split*4]...)

	return w, x, y, z
}

func bubbleSort(integerSlice []int) []int {

	var swapped bool
	for {
		swapped = false
		for index := 0; index < len(integerSlice)-1; index++ {

			if integerSlice[index] > integerSlice[index+1] {
				swap(integerSlice, index)
				// swap flag
				swapped = true
			}
		}
		// swap check
		if !swapped {
			break
		}
	}
	return integerSlice
}

func swap(integerSlice []int, index int) {

	comp1 := integerSlice[index]
	comp2 := integerSlice[index+1]

	integerSlice[index+1] = comp1
	integerSlice[index] = comp2
}

func merge(slices ...[]int) []int {
	var sourceSlice []int
	for _, i := range slices {
		sourceSlice = append(sourceSlice, i...)
	}
	return sourceSlice
}
