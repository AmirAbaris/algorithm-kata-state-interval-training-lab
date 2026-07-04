package running_max

// RunningMax returns, for each index i, the maximum value in nums[0:i+1].
func RunningMax(nums []int) []int {
	result := make([]int, 0, len(nums))
	best := nums[0]

	for _, num := range nums {
		if best < num {
			best = num
			result = append(result, best)
		} else {
			result = append(result, best)
		}
	}

	return result
}
