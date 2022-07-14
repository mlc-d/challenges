package hackerrank

func lonelyinteger(a []int32) (res int32) {
	myMap := make(map[int32]uint8)
	for _, v := range a {
		myMap[v]++
	}
	for i, v := range myMap {
		if v == 1 {
			res = i
		}
	}
	return
}
