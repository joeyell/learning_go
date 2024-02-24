// main.go

package week_2

func BubbleSort(integerSlice []int) []int {

	var swapped bool
	for {
		swapped = false
		for index := 0; index < len(integerSlice)-1; index++ {

			if integerSlice[index] > integerSlice[index+1] {
				Swap(integerSlice, index)
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

func Swap(integerSlice []int, index int) {

	comp1 := integerSlice[index]
	comp2 := integerSlice[index+1]

	integerSlice[index+1] = comp1
	integerSlice[index] = comp2
}
