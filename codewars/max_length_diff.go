package codewars

type Number interface {
	~float64 | ~float32 | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Abs implementation using generics to return the absolute value of a number
func Abs[T Number](number T) T {
	if number < 0 {
		return number * -1
	}
	return number
}

// MaxLengthDiff given two slices of strings, MaxLengthDiff() returns the difference between
// the longest string in one slice and the shortest in the other slice
func MaxLengthDiff(a1, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	var diff int
	for _, x := range a1 {
		for _, y := range a2 {
			if Abs(len(x)-len(y)) > diff {
				diff = Abs(len(x) - len(y))
			}
		}
	}
	return diff
}
