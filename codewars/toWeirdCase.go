package codewars

import "unicode"

// ToWeirdCase accepts a string, and returns the same string with all even indexed characters
// in each word upper cased,and all odd indexed characters in each word lower cased.
func ToWeirdCase(str string) string {
	r := []rune(str)
	index := 0
	for i, v := range r {
		if v == ' ' {
			index = 0
			continue
		}
		if index%2 == 0 {
			r[i] = unicode.ToUpper(r[i])
			index++
			continue
		}
		r[i] = unicode.ToLower(r[i])
		index++
	}
	return string(r)
}
