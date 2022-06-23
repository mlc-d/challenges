package codewars

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
	codewars.DiffLists([]int{1, 2}, []int{1})       // {2}
	codewars.DiffLists([]int{1, 2, 2}, []int{1})    // {2, 2}
	codewars.DiffLists([]int{1, 2, 2}, []int{2})    // {1}
	codewars.DiffLists([]int{1, 2, 2}, []int{})     // {1, 2, 2}
	codewars.DiffLists([]int{}, []int{1, 2})        // {}
	codewars.DiffLists([]int{1, 2, 3}, []int{1, 2}) // {3}
*/

// DiffLists Returns the first slice after removing all items
// also contained in the second slice.
func DiffLists(a, b []int) []int {
	for _, v := range b {
		size := len(a)
		for i := 0; i < size; i++ {
			if a[i] == v {
				// remove item by reslicing
				a = append(a[:i], a[i+1:]...)
				// since item was removed, decreases the index by 1
				i -= 1
				size = len(a)
			}
		}
	}
	return a
}

/*
	{20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5} = 5
    {1,1,2,-2,5,2,4,4,-1,-2,5} = -1
    {20,1,1,2,2,3,3,5,5,4,20,4,5} = 5
    {10} = 10
    {1,1,1,1,1,1,10,1,1,1,1} = 10
    {5,4,3,2,1,5,4,3,2,10,10} = 1
*/

// FindOdd returns the element in the slice that appears an odd amount of times.
// It's guaranteed that there's always going to be 1 element meeting that condition.
func FindOdd(seq []int) int {
	numbers := make(map[int]int)
	for _, v := range seq {
		numbers[v] += 1
	}
	for i, v := range numbers {
		if v%2 != 0 {
			return i
		}
	}
	return 0
}

func intAbsValue(a int) int {
	if a >= 0 {
		return a
	}
	return a * -1
}

// MaxLengthDiff given two slices of strings, MaxLengthDiff() returns the difference between
// the longest string in one slice and the shortest in the other slice
func MaxLengthDiff(a1, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	var diff int
	for _, x := range a1 {
		for _, y := range a2 {
			if intAbsValue(len(x)-len(y)) > diff {
				diff = intAbsValue(len(x) - len(y))
			}
		}
	}
	return diff
}

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
func AccumAlt(s string) string {
	parts := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
	}
	return strings.Join(parts, "-")
}

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

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(input []int) PosPeaks {
	length := len(input)
	pos := make([]int, 0)
	peaks := make([]int, 0)
	// var altIndex int

	if length < 3 {
		return PosPeaks{pos, peaks}
	}

	for i := 1; i < length-1; i++ {
		fmt.Println("testing:", input[i])
		if input[i] > input[i-1] && input[i] > input[i+1] {
			pos = append(pos, i)
			peaks = append(peaks, input[i])
		}
		if input[i] > input[i-1] && input[i] == input[i+1] {
			for j := i + 1; j < length; j++ {
				fmt.Println("testing j:", input[j])
				if input[j] > input[i] {
					break
				}
				if input[j] < input[i] {
					pos = append(pos, i)
					peaks = append(peaks, input[i])
					i = j
					break
				}
			}

		}
	}

	return PosPeaks{pos, peaks}
}

/*
func PickPeaks(array []int) PosPeaks {
	length := len(array)
	var end int
	end = array[length-1]
	climbing := false
	var localPeak int
	pos := make([]int, 0)
	peaks := make([]int, 0)
	for i := 1; i < length; i++ {
		if array[i] > array[i-1] {
			climbing = true
			localPeak = array[i]
		} else {
			if climbing {
				pos = append(pos, i-1)
				peaks = append(peaks, localPeak)
			}
			climbing = false
		}
	}
	lengthPeaks := len(peaks)
	if lengthPeaks > 0 {
		if peaks[lengthPeaks-1] <= end && !climbing {
			peaks = peaks[:lengthPeaks-1]
			pos = pos[:len(pos)-1]
		}
	}
	return PosPeaks{Pos: pos, Peaks: peaks}
}
*/
