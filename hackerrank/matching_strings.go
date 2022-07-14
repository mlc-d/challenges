package hackerrank

func matchingStrings(data []string, queries []string) []int32 {
	res := make([]int32, len(queries))
	for i, v := range queries {
		for _, x := range data {
			if x == v {
				res[i]++
			}
		}
	}
	return res
}
