package hackerrank

import "retos/utils"

func maximumPerimeterTriangle(sticks []int32) []int32 {
	utils.RadixSort(sticks)

	l := len(sticks)

	if l < 3 {
		return []int32{-1}
	}

	res := make([]int32, 3)

	for i := l - 1; i >= 2; i-- {
		if sticks[i-1] <= sticks[i]/2 {
			continue
		}
		if sticks[i-2]+sticks[i-1] <= sticks[i] {
			continue
		}
		if res[2] >= sticks[i] {
			continue
		}
		res[0] = sticks[i-2]
		res[1] = sticks[i-1]
		res[2] = sticks[i]
	}

	if res[2] != 0 {
		return res
	}
	return []int32{-1}
}
