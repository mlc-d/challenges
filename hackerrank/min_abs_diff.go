package hackerrank

import (
	"fmt"
	"math"
	"retos/utils"
)

func minimumAbsoluteDifference(arr []int32) int32 {
	l := len(arr)

	var min = int32(math.MaxInt32)
	for i := 0; i < l; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	min = utils.Abs(min)

	for i := 0; i < l; i++ {
		/*		if arr[i] < 0 {
				arr[i] -= min
				continue
			}*/
		arr[i] += min
	}

	utils.RadixSort(arr)
	for i := 1; i < l; i++ {
		x := arr[i] - arr[i-1]
		fmt.Println("testing:", arr[i-1], arr[i], x)
		if utils.Abs(x) < min {
			min = x
		}
	}
	return min
}
