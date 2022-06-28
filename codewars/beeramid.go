package codewars

func Beeramid(bonus int, price float64) int {
	cans := int(float64(bonus) / price)
	var i = 0
	for cans-i*i >= 0 {
		cans -= i * i
		i++
	}
	return i - 1
}
