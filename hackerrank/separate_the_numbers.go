package hackerrank

import (
	"strconv"
)

func separateNumbers(s string) bool {
	// vars used through the function
	l := len(s)
	var flag bool
	var altString string

	// base case
	if s[0] == '0' {
		altString = "0"
		next := 1
		for len(altString) < l {
			altString += strconv.Itoa(next) // fill the alternative string with consecutive numbers
			next++
		}
		if altString == s {
			flag = true
		}
	} else {
		res := 0
		for i := 1; i <= l/2; i++ {
			res *= 10
			res += int(s[i-1] - '0') // removes any '0' from the rune, and convert it to int
			next := res + 1
			altString = s[:i]
			for len(altString) < l {
				altString += strconv.Itoa(next)
				next++
			}
			if altString == s {
				flag = true
				break
			}
		}
	}
	if flag && l > 1 {
		return true
	}
	return false
}
