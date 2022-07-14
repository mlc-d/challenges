package hackerrank

import "retos/codewars"

func diagonalDifference(arr [][]int32) int32 {
	length := len(arr[0])
	var diagonal1, diagonal2 int32

	for i := 0; i < length; i++ {
		diagonal1 += arr[i][i]
		diagonal2 += arr[i][length-i]
	}

	return codewars.Abs(diagonal1 - diagonal2)
}
