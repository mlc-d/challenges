package hackerrank

func migratoryBirds(arr []int32) int32 {
	values := make(map[int32]int32, 5)
	for _, v := range arr {
		values[v]++
	}

	var max, i, res int32

	for i = 5; i > 0; i-- {
		if values[i] >= max {
			max = values[i]
			res = i
		}
	}

	return res
}
