package hackerrank

import (
	"strings"
)

var (
	responses = map[bool]string{true: "pangram", false: "not pangram"}
)

func pangrams(s string) string {
	s1 := strings.ToUpper(s)
	letters := make(map[uint8]struct{}, 26)
	var i uint8
	for i = 65; i <= 90; i++ {
		letters[i] = struct{}{}
	}
	for j := 0; j < len(s1); j++ {
		delete(letters, s1[j])
		if len(letters) == 0 {
			return responses[true]
		}
	}
	return responses[false]
}
