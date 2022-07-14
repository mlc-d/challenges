package hackerrank

import "retos/utils"

func twoArrays(k int32, A []int32, B []int32) string {
	length := len(A)

	utils.RadixSort(A)
	utils.RadixSort(B)
	for i := 0; i < length; i++ {
		if A[i]+B[length-(i+1)] < k {
			return "NO"
		}
	}
	return "YES"
}
