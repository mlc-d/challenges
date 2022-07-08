package hackerrank

func pageCount(n int32, p int32) int32 {
	left := p / 2
	right := (n - p) / 2
	if n%2 == 0 && n-p == 1 {
		right++
	}
	if left > right {
		return right
	}
	return left
}
