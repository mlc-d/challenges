package codewars

func validateLine(line []int) bool {
	//fmt.Println(line)
	myMap := make(map[int]int, 9)
	for _, v := range line {
		myMap[v]++
		if myMap[v] > 1 || v == 0 {
			return false
		}
	}
	return true
}

func validateInnerSquare(square [3][3]int) bool {
	subArray := make([]int, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			subArray[3*i+j] = square[i][j]
		}
	}
	return validateLine(subArray)
}

func ValidateSolution(sudoku [][]int) bool {
	var i, j int

	for _, v := range sudoku {
		if !validateLine(v) {
			return false
		}
	}

	column := make([]int, 9)
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			column[j] = sudoku[j][i]
		}
		if !validateLine(column) {
			return false
		}
	}

	var innerSquares [9][3][3]int
	for h := 0; h < 9; h++ {
		for i = 0; i < 3; i++ {
			for j = 0; j < 3; j++ {
				innerSquares[h][i][j] = sudoku[i][j]
			}
		}
		if !validateInnerSquare(innerSquares[h]) {
			return false
		}
	}
	return true
}
