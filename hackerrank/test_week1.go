package hackerrank

import (
	"retos/utils"
)

func FindMedian(arr []int32) int32 {
	utils.RadixSort(arr)
	return arr[len(arr)/2]
}
