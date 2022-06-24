package codewars

import "strings"

// Accum Set the character as uppercase, and repeat its lowercase equivalent
// n times, where n is the character's original index
func Accum(s string) string {
	var res string
	for i, v := range s {
		if i == 0 {
			res += strings.ToUpper(string(v))
		}
		for j := 0; j < i; j++ {
			if j == 0 {
				res += strings.ToUpper(string(v))
			}
			res += strings.ToLower(string(v))
		}
		if i != len(s)-1 {
			res += "-"
		}
	}
	return res
}

// AccumAlt clever solution from another users, notice the strings methods
// func AccumAlt(s string) string {
// 	parts := make([]string, len(s))
// 	for i := 0; i < len(s); i++ {
// 		parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
// 	}
// 	return strings.Join(parts, "-")
// }
