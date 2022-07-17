package hackerrank

import (
	"fmt"
)

func anagram(s string) int32 {
	l := len(s)

	if l%2 != 0 {
		return -1
	}

	mid := (l / 2)

	m1 := make(map[uint8]int32, l/2)
	m2 := make(map[uint8]int32, l/2)

	for i := 0; i < mid; i++ {
		m1[s[i]]++
		m2[s[i+mid]]++
	}
	var r int32
	for i, v := range m1 {
		fmt.Printf("Testing: %s (%d).\n", string(i), v)
		var ok bool
		if _, ok = m2[i]; !ok {
			r += v
		} else {
			if v > m2[i] {
				r += v - m2[i]
			}
		}
	}
	return r
}
