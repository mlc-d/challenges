package hackerrank

func countingSort(arr []int32) []int32 {
	count := make([]int32, 100)

	for _, v := range arr {
		count[v]++
	}
	return count
}
