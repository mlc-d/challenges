package codewars

import (
	"strconv"
	"strings"
)

// RGB returns the hexadecimal representation of the color (full length) represented
// as rgb by the arguments
func RGB(values ...int) string {
	// rounds values to the closest integer between 0-255
	for i := 0; i < len(values); i++ {
		if values[i] < 0 {
			values[i] = 0
		}
		if values[i] > 255 {
			values[i] = 255
		}
	}

	var r, g, b string

	if len(strconv.FormatInt(int64(values[0]), 16)) < 2 {
		r = "0" + strconv.FormatInt(int64(values[0]), 16)
	} else {
		r = strconv.FormatInt(int64(values[0]), 16)
	}

	if len(strconv.FormatInt(int64(values[1]), 16)) < 2 {
		g = "0" + strconv.FormatInt(int64(values[1]), 16)
	} else {
		g = strconv.FormatInt(int64(values[1]), 16)
	}

	if len(strconv.FormatInt(int64(values[2]), 16)) < 2 {
		b = "0" + strconv.FormatInt(int64(values[2]), 16)
	} else {
		b = strconv.FormatInt(int64(values[2]), 16)
	}

	return strings.ToUpper(r + g + b)
}
