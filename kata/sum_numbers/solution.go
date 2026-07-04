package sum_numbers

// Sum returns the total of all numbers in nums.
// Empty slice returns 0.
func Sum(nums []int) int {
	var total int 

	for _, num := range nums {
		total += num
	}

	return total
}
