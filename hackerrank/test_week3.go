package hackerrank

func getTotalX(a []int32, b []int32) int32 {

	// After submitting the solution without these two lines, it worked,
	// so the two arrays are presumably sorted.
	// RadixSort(a)
	// RadixSort(b)

	la := len(a)

	greatestFactor := a[la-1]

	var res int32
	var flag bool

	candidates := make([]int32, 0)

	for i := greatestFactor; i <= b[0]; i += greatestFactor {
		flag = true
		for j := 0; j < la; j++ {
			if i%a[j] != 0 {
				flag = false
				break
			}
		}
		if flag {
			candidates = append(candidates, i)
		}
	}

	for _, u := range candidates {
		flag = true
		for _, v := range b {
			if v%u != 0 {
				flag = false
				break
			}
		}
		if flag {
			res++
		}
	}
	return res
}
