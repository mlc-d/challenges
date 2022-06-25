package codewars

func calcFibNumber(args ...float64) (res float64) {
	for _, v := range args {
		res += v
	}
	return
}

func Tribonacci(signature [3]float64, n int) (res []float64) {
	res = signature[:]
	for len(res) < n {
		res = append(res, calcFibNumber(res[len(res)-3:]...))
	}
	return res[:n] // if n < len(signature), only values before n index should be returned
}
