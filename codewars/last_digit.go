package codewars

import "fmt"

func lastDigit(b, exp int) int {
	x := b % 10
	y := exp % 2
	z := exp % 4
	if exp == 1 {
		return x
	}
	if exp == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	if x == 1 || x == 5 || x == 6 {
		return x
	}
	if x == 4 {
		if y == 0 {
			return 6
		}
		return x
	}
	if x == 9 {
		if y == 0 {
			return 1
		}
		return x
	}
	if z == 0 {
		if x%2 == 0 {
			return 6
		}
		return 1
	}

	if z == 1 {
		return x
	}
	if z == 2 {
		if x%2 == 0 {
			return 4
		}
		return 9
	}
	if x == 2 {
		return 8
	}
	if x == 3 {
		return 7
	}
	if x == 7 {
		return 3
	}
	return 2
}

func intPow(b, exp int) int {
	if exp == 0 {
		return 1
	}
	res := 1
	for exp > 0 {
		res *= b
		exp--
	}
	return res
}

func LastDigit(as []int) int {
	l := len(as)
	if l < 1 {
		return 1
	}
	if l == 1 {
		return as[0] % 10
	}
	if l >= 2 {
		for i := l - 2; i >= 0; i-- {
			as[i] = lastDigit(as[i], as[i+1]) % 4
			fmt.Printf("as[i]: %d.\n", as[i])
		}
	}
	return as[0]
}
