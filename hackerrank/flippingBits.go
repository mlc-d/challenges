package hackerrank

import (
	"strconv"
)

func decimalToBinaryLiteral(n uint32) []rune {
	res := make([]rune, 32)
	x := []rune(strconv.FormatInt(int64(n), 2))
	for i := 0; i < 32-len(x); i++ {
		res[i] = '0'
	}
	return append(res[:32-len(x)], x...)
}

func binaryStringToDecimal(arr []rune) (res int64) {
	res, err := strconv.ParseInt(string(arr), 2, 64)
	if err != nil {
		panic(err.Error())
	}
	return res
}

func flipBits(arr []rune) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == '0' {
			arr[i] = '1'
			continue
		}
		arr[i] = '0'
	}
}

func FlippingBits(m int64) int64 {
	n := uint32(m)
	binaryNumber := decimalToBinaryLiteral(n)
	flipBits(binaryNumber)
	return binaryStringToDecimal(binaryNumber)
}
