package apply_deltas_trace

// Trace applies each delta to start in order and returns the value after each step.
func Trace(start int, deltas []int) []int {
	// start = 10, deltas = [+3, -5, +2]
	// Output: [13, 8, 10]

	// loop the delta
	// mutate start on each delta itorate
	// append mutated value to new slice
	result := make([]int, 0, len(deltas))

	for _, num := range deltas {
		start += num
		result = append(result, start)
	}

	return result
}
