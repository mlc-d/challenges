package hackerrank

func pickingNumbers(a []int32) int32 {
	RadixSort(a)
	var res, max int32

	for i := 0; i < len(a)-1; i++ {
		max = 1
		for j := i + 1; j < len(a); j++ {
			if a[j] == a[i] || a[j] == a[i]+1 {
				max++
				continue
			}
			i = j - 1
			break
		}
		if max > res {
			res = max
		}
	}
	return res
}
