package hackerrank

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

// splitCamelCase splits a camel case identifier into space-separated words
func splitCamelCase(s string) (res string) {
	for _, character := range s {
		if character == '(' {
			break
		}
		switch unicode.IsUpper(character) {
		case true:
			res += " " + string(character)
		default:
			res += string(character)
		}
	}
	return strings.TrimSpace(strings.ToLower(res))
}

// combineCamelCase creates a camel case identifier from space-separated words
func combineCamelCase(s string) (res string) {
	var mustUppercase bool
	for _, v := range s {
		if v == ' ' {
			mustUppercase = true
			continue
		}
		if mustUppercase {
			res += string(unicode.ToUpper(v))
		} else {
			res += string(v)
		}
		mustUppercase = false
	}
	return strings.Trim(res, "\n")
}

// CamelCase Create an camel case identifier, or splits one into its words
// separated by space and in lowercase
func CamelCase(s1 string) (res string) {
	s := strings.Trim(s1, "\n")
	fmt.Printf("from input's length: %d, value: %s.\n", len(s1), s1)
	fmt.Printf("made input's length: %d, value: %s.\n\n", len(s), s)
	switch s[0] == 'C' {
	case true:
		switch s[2] {
		case 'C':
			res = string(bytes.ToUpper([]byte{s[4]})) + combineCamelCase(s[5:]) + "\n"
		case 'M':
			res = combineCamelCase(s[4:]) + "()\n"
		case 'V':
			res = combineCamelCase(s[4:]) + "\n"
		}
	case false:
		res = splitCamelCase(s[4:]) + "\n"
	}
	return
}
