package hackerrank

func sum(args ...int32) (res int32) {
	for _, v := range args {
		res += v
	}
	return
}

func birthday(s []int32, d int32, m int32) (res int32) {
	for i := 0; i <= len(s)-int(m); i++ {
		if sum(s[i:i+int(m)]...) == d {
			res++
		}
	}
	return
}
