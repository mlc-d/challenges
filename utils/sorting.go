package utils

// countSort sorts elements using a hashmap-like structure.
// Works only for positive integer values, within O(n),
// where n is the number of elements in the array.
func countSort(n, exp int32, arr []int32) {
	// create variables
	var i int32
	output := make([]int32, n)
	count := make([]int32, 10)
	// fills the count array according to each value's digit (value/exponent)%10
	for _, v := range arr {
		count[(v/exp)%10]++
	}
	// adds the previous value to each position, thus giving the starting index
	// for the sorted array
	for i = 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	// takes the element from the original array, and finds its target position ( count[(arr[i] / exp) % 10 )
	// using the count array. Then, stores the value into output[target position - 1]
	for i = n - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}
	// copies the sorted array to the original one
	for i = 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// RadixSort sorts an array of positive integer values based on their digits,
// starting with the less significant digit. Since countSort() is a stable
// sorting algorithm, it's guaranteed that after sorting one order of magnitude,
// the next one will be sorted correctly since the smaller magnitude is already sorted.
func RadixSort(arr []int32) {
	n := int32(len(arr))
	m := getMax(arr)

	var exp int32

	for exp = 1; m/exp > 0; exp *= 10 {
		countSort(n, exp, arr)
	}
}
