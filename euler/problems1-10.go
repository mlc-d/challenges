package euler

// p. 1. If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
func Multiples3or5() (res int) {
	for i := 3; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			res += i
		}
	}
	return
}

// Problem 2. Finds the sum of even fibonacci numbers below 4 000 000
func EvenFibNumbers() (res int) {
	slice := []int{1, 2}
	x := slice[len(slice)-2] + slice[len(slice)-1]
	for x < 4000000 {
		slice = append(slice, x)
		if x%2 == 0 {
			res += x
		}
		x = slice[len(slice)-2] + slice[len(slice)-1]
	}
	return res + 2
}
