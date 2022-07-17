package hackerrank

func transform(c int32, k int32) int32 {
	k = k % 26
	if c <= 90 {
		c += k
		if c > 90 {
			c -= 26
		}
		return c
	}
	c += k

	if c > 122 {
		c -= 26
	}

	return c
}

func caesarCipher(s string, k int32) string {
	var cypheredString string

	for _, v := range s {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			cypheredString += string(transform(v, k))
		} else {
			cypheredString += string(v)
		}
	}

	return cypheredString
}
