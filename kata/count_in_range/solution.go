package count_in_range

// Count returns how many values in nums satisfy low <= x <= high.
func Count(nums []int, low, high int) int {
	// nums = [1, 5, 3, 8, 2], low=2, high=5
	// output = 3

	// loop the nums
	// check condition of: low <= x <= high
	// if met, increase else nothing (default 0)
	// return counter

	counter := 0

	for _, num := range nums {
		if low <= num && num <= high {
			counter++
		}
	}

	return counter
}
