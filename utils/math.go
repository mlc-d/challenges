package utils

type Numeric interface {
	~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func GetGreatestNumber[T Numeric](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func Abs[T Numeric](a T) T {
	if a >= 0 {
		return a
	}
	return a - a - a
}

func getMax(arr []int32) (res int32) {
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return
}
