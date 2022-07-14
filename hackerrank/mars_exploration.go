package hackerrank

var message = [3]byte{83, 79, 83}

func countDiff(a, b [3]byte, counter *int32) {
	for i := 0; i < 3; i++ {
		if a[i] != b[i] {
			*counter++
		}
	}
}

func marsExploration(s string) (res int32) {
	for i := 0; i < len(s)-2; i += 3 {
		countDiff([3]byte{s[i], s[i+1], s[i+2]}, message, &res)
	}
	return
}
