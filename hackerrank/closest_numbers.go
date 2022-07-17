package hackerrank

import (
	"fmt"
	"math"
	"retos/utils"
)

// getDiff returns the distance between two numbers.
func getDiff(a, b int32) int32 {
	if b < 0 {
		a, b = a*-1, b*-1
	}
	return utils.Abs(b - a)
}

// closestNumbers returns an array of the adjacent numbers (after sorting) that have the smallest.
func closestNumbers(arr []int32) []int32 {
	// vars used through the function
	l := len(arr)
	i := 0
	min := int32(math.MaxInt32)
	// find the smallest value in the array
	for ; i < l; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	// if the minimum value is negative, make it positive, so we can use it to normalize the array
	min = utils.Abs(min)
	for i = 0; i < l; i++ {
		arr[i] += min // normalize the array so every value is positive (stable)
	}
	utils.RadixSort(arr) // sort the normalized array
	for i = 0; i < l; i++ {
		arr[i] -= min // return every value to its original value
	}
	// find the smallest difference between array's items
	minDiff := int32(math.MaxInt32)
	for i = 0; i <= l-2; i++ {
		if getDiff(arr[i], arr[i+1]) < minDiff {
			minDiff = getDiff(arr[i], arr[i+1])
		}
	}
	// fill a new array with each pair that meets the minimum difference
	r := make([]int32, 0)
	for i = 0; i <= l-2; i++ {
		fmt.Printf("Index: %d\tValue: %d.\n", i, arr[i])
		if getDiff(arr[i], arr[i+1]) == minDiff {
			r = append(r, arr[i])
			r = append(r, arr[i+1])
		}
	}
	return r
}
