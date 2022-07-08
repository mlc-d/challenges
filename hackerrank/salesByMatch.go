package hackerrank

func sockMerchant(n int32, ar []int32) int32 {
	values := make(map[int32]int32, n)
	for _, v := range ar {
		values[v]++
	}
	var res int32
	for _, v := range values {
		res += v / 2
	}
	return res
}
