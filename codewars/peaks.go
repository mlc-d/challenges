package codewars

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// PickPeaks returns the values and positions of
// partial maxima within an int slice.
func PickPeaks(input []int) PosPeaks {
	length := len(input)
	pos := make([]int, 0)
	peaks := make([]int, 0)

	if length < 3 {
		return PosPeaks{pos, peaks}
	}

	for i := 1; i < length-1; i++ {
		if input[i] > input[i-1] && input[i] > input[i+1] {
			pos = append(pos, i)
			peaks = append(peaks, input[i])
		}
		if input[i] > input[i-1] && input[i] == input[i+1] {
			for j := i + 1; j < length; j++ {
				if input[j] > input[i] {
					break
				}
				if input[j] < input[i] {
					pos = append(pos, i)
					peaks = append(peaks, input[i])
					i = j
					break
				}
			}

		}
	}
	return PosPeaks{pos, peaks}
}
