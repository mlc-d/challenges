package hackerrank

func towerBreakers(n int32, m int32) int32 {
	if m == 1 || n%2 == 0 {
		return 2
	}
	return 1
}
