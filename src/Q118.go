package src

func generate(numRows int) [][]int {
	// 行从1开始,列从0开始
	result := make([][]int, 0, numRows)
	result = append(result, []int{1})
	if numRows == 1{
		return result
	}
	result = append(result, []int{1, 1})
	if numRows == 2{
		return result
	}

	for i := 3; i <= numRows; i++ {
		index := i-1
		l := make([]int, i)
		l[0], l[i-1] = 1, 1
		for j := 1; j < i-1; j++ {
			l[j] = result[index-1][j-1] + result[index-1][j]
		}
		result = append(result, l)
	}
	return result
}
