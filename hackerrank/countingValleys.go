package hackerrank

func moveUpOrDown(height *int, input byte) {
	switch input {
	case 'U':
		*height++
	case 'D':
		*height--
	}
}

func countingValleys(steps int32, path string) (res int32) {
	height := 0
	length := len(path)

	for i := 0; i < length; i++ {

		moveUpOrDown(&height, path[i])

		if height < 0 {
			for height < 0 {
				i++
				moveUpOrDown(&height, path[i])
			}
			res++
		}
	}
	return res
}
