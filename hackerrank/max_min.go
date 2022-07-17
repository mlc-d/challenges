package hackerrank

import (
	"fmt"
	"math"
	"retos/utils"
)

func maxMin(k int32, arr []int32) int32 {
	utils.RadixSort(arr)
	l := len(arr)

	if k == int32(l) {
		return arr[l] - arr[0]
	}

	var min = int32(math.MaxInt32)

	for i := 0; i <= l-int(k); i++ {
		fmt.Printf("Testing: %v.\n", arr[i:(int(k)-1)+i])
		if min > arr[int(k-1)+i]-arr[i] {
			min = arr[int(k-1)+i] - arr[i]
		}
	}
	return min
}
