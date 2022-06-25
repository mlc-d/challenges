package hackerrank

// https://www.hackerrank.com/challenges/three-month-preparation-kit-divisible-sum-pairs/
func divisibleSumPairs(n int32, k int32, ar []int32) (res int32) {
	var i, j int32
	for ; i < n-1; i++ {
		j = i + 1
		for ; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				res++
			}
		}
	}
	return
}
