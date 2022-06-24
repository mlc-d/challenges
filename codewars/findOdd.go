package codewars

// FindOdd returns the element in the slice that appears an odd amount of times.
// It's guaranteed that there's always going to be 1 element meeting that condition.
func FindOdd(seq []int) int {
	numbers := make(map[int]int)
	for _, v := range seq {
		numbers[v] += 1
	}
	for i, v := range numbers {
		if v%2 != 0 {
			return i
		}
	}
	return 0
}
