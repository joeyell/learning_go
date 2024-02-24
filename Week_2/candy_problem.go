// candy_problem.go

// There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

// You are giving candies to these children subjected to the following requirements:

// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
// Return the minimum number of candies you need to have to distribute the candies to the children.

// Input: ratings = [1,0,2]
// Output: 5
// Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

package week_2

// {1, 0, 0, 0, 0, 7} == 8 (y)
// {1, 0, 1, 0, 1} == 8 (n)
// {0, 0, 1, 1, 1, 1} == 6 (y)
// {5, 4, 3, 2, 3, 4, 5} = 19 (n)
// {7, 6, 5, 4, 3, 2, 1} == 28 (y)
// {1, 2, 3, 4, 5, 6, 7} == 28 (y)

func Candy_problem(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}

func Candy_problem_2(childRatings []int) int {

	line := len(childRatings)
	finalCandyNum := 0
	finalCandyNum += line

	candyCounter := 0
	for i := 0; i < line-1; i++ {

		comp1 := childRatings[i]
		comp2 := childRatings[i+1]

		if comp1 == comp2 {
			candyCounter = 0
		} else {
			candyCounter++
		}

		finalCandyNum += candyCounter

	}
	return finalCandyNum
}

// 5 4 = +1
// 4 3 = +1
// 3 2 = +1
// 2 3 = +1
// 3 4 = +1
// 4 5 = +1
