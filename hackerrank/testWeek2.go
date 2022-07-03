package hackerrank

import (
	"fmt"
	"math"
)

func greatestNumber(args ...int32) int32 {
	var max int32 = math.MinInt32
	for _, v := range args {
		fmt.Println("eval:", v)
		if v > max {
			max = v
		}
	}
	fmt.Println()
	return max
}

func flippingMatrix(matrix [][]int32) (res int32) {
	length := len(matrix[0])

	t := length - 1

	for h := 0; h < length/2; h++ {
		for i := 0; i < length/2; i++ {
			x := t - i
			res += greatestNumber(matrix[h][i], matrix[h][x], matrix[t-h][i], matrix[t-h][x])
		}
	}
	return res
}
