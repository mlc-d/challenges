package hackerrank

func twoArrays(k int32, A []int32, B []int32) string {
	length := len(A)

	RadixSort(A)
	RadixSort(B)
	for i := 0; i < length; i++ {
		if A[i]+B[length-(i+1)] < k {
			return "NO"
		}
	}
	return "YES"
}
