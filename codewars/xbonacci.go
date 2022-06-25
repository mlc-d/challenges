package codewars

func Xbonacci(signature []float64, n int) []float64 {
	solution := signature[:]
	lengthSignature := len(signature)
	solution = append(solution, make([]float64, n-lengthSignature)...)
	for i := lengthSignature; i < n; i++ {
		solution[i] = func() (sum float64) {
			for _, v := range solution[(i)-lengthSignature:] {
				sum += v
			}
			return
		}()
	}
	return solution[:n] // if n < len(signature), only values before n index should be returned

	//solution from @akar-0. This solutions is more efficient since it allocates the memory only once
	//l := len(signature)
	//if l >= n {
	//	return signature[:n]
	//}
	//o := signature[:]
	//o = append(o, make([]float64, n-l)...)
	//for _, x := range signature {
	//	o[l] += x
	//}
	//for i := l + 1; i < n; i++ {
	//	o[i] = o[i-1]*2 - o[i-l-1]
	//}
	//fmt.Println(len(o))
	//return o
}
