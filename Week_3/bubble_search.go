// bubble_search.go

package week_3

import "fmt"

func Merge(nums1 []int, len1 int, nums2 []int, len2 int) []int {

	nums1 = append(nums1, nums2...)

	fmt.Println(nums1)

	nums4 := []int{}
	for _, v := range nums1 {
		if v != 0 {
			nums4 = append(nums4, v)
		}
	}
	nums1 = nums4

	// loop until full sorted
	for {
		//initalise swap checker
		swapped := false
		// loop through array
		for i := 0; i < len(nums4)-1; i++ {
			comp1 := nums4[i]
			comp2 := nums4[(i + 1)]

			if comp1 > comp2 {
				// rearrange in slice
				nums4[i+1] = comp1
				nums4[i] = comp2
				// swap flag
				swapped = true
			}
		}
		// swap check
		if !swapped {
			break
		}
	}
	return nums1
}
