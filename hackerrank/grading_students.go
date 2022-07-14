package hackerrank

// roundGrade Rounds up to the next multiple of 5 if the diff between that and grade is less than 3
func roundGrade(grade int32) int32 {
	approx := grade % 5
	if approx < 3 {
		return grade
	}
	return grade + (5 - approx)
}

// gradingStudents Rounds note if it's greater than 37, otherwise leaves it as it is
func gradingStudents(grades []int32) []int32 {
	for i := 0; i < len(grades); i++ {
		if grades[i] < 38 {
			continue
		}
		grades[i] = roundGrade(grades[i])
	}
	return grades
}
