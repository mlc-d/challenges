package codewars

import (
	"fmt"
	"math"
)

// calcRate Rate grows .5 every two res[0]
func calcRate(rate float64, month int) float64 {
	if month%2 == 0 {
		return rate + 0.5
	}
	return rate
}

// calcCarPrice cars lose their value at a given monthly rate
func calcCarPrice(value float64, rate float64) float64 {
	return value - (value/100)*rate
}

// calcSavings returns the total savings after a period of time measured in res[0]
func calcSavings(monthySavings float64, month int) float64 {
	return monthySavings * float64(month)
}

// CalcLeft comment
// func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
func CalcLeft(startPriceOld, startPriceNew, savingperMonth int, rate float64) (res [2]int) {
	if startPriceOld >= startPriceNew {
		return [2]int{0, startPriceOld - startPriceNew}
	}
	savings := 0.0
	valueOldCar := float64(startPriceOld)
	valueNewCar := float64(startPriceNew)
	monthlySavings := float64(savingperMonth)
	left := 0.0
	for {
		res[0]++
		rate = calcRate(rate, res[0])
		valueOldCar = calcCarPrice(valueOldCar, rate)
		valueNewCar = calcCarPrice(valueNewCar, rate)
		savings = calcSavings(monthlySavings, res[0])
		left = (valueOldCar + savings) - valueNewCar
		if valueNewCar <= valueOldCar || left >= 0 {
			break
		}
	}
	fmt.Println("aver", left)
	res[1] = int(math.Round(left))
	return
}
