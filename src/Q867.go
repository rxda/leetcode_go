package src

func transpose(A [][]int) [][]int {
	row := len(A)
	if row == 0 {
		return [][]int{}
	}
	coloumn := len(A[0])
	var result = make([][]int, coloumn)

	for i := 0; i < coloumn; i++ {
		result[i] = make([]int ,row)
		for rowK, rowV := range A {
			result[i][rowK] = rowV[i]
		}
	}
	return result
}
