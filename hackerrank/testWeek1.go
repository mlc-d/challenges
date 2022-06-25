package hackerrank

func getMax(arr []int32) (res int32) {
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return
}

func countSort(n, exp int32, arr []int32) {
	var i int32
	output := make([]int32, n)
	count := make([]int32, 10)

	for _, v := range arr {
		count[(v/exp)%10]++
	}

	for i = 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i = n - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}

	for i = 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func radixsort(arr []int32) {
	n := int32(len(arr))
	m := getMax(arr)

	var exp int32
	
	for exp = 1; m/exp > 0; exp *= 10 {
		countSort(n, exp, arr)
	}
}

func FindMedian(arr []int32) int32 {
	radixsort(arr)
	return arr[len(arr)/2]
}
