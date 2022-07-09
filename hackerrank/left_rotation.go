package hackerrank

func rotateLeft(d int32, arr []int32) []int32 {

	l := int32(len(arr))

	if d == l {
		return arr
	}

	stepsToRight := l - d

	res := make([]int32, 0)

	res = append(res, arr[l-stepsToRight:]...)
	res = append(res, arr[:d]...)

	return res
}
