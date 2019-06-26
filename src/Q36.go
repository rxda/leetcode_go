package src

func IsValidSudoku(board [][]byte) bool {

	numMap := map[byte]bool{
		'1': false,
		'2': false,
		'3': false,
		'4': false,
		'5': false,
		'6': false,
		'7': false,
		'8': false,
		'9': false,
	}
	//row
	for _, rowV := range board {
		tempNumMap := copyMap(numMap)
		for _, v := range rowV {
			if v == '.' {
				continue
			}
			if val, ok := tempNumMap[v]; !ok {
				return false
			} else {
				if val {
					return false
				}
				tempNumMap[v] = true
			}
		}
	}
	//coloumn
	for i := 0; i < 9; i++ {
		tempNumMap := copyMap(numMap)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if val, ok := tempNumMap[board[j][i]]; !ok {
				return false
			} else {
				if val {
					return false
				}
				tempNumMap[board[j][i]] = true
			}
		}
	}
	//block
	block := [][]int{
		{0, 0}, {0, 3}, {0, 6}, {3, 0}, {3, 3}, {3, 6}, {6, 0}, {6, 3}, {6, 6},
	}
	for _, b := range block {
		tempNumMap := copyMap(numMap)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[b[0]+i][b[1]+j] == '.' {
					continue
				}
				if val, ok := tempNumMap[board[b[0]+i][b[1]+j]]; !ok {
					return false
				} else {
					if val {
						return false
					}
				}
				tempNumMap[board[b[0]+i][b[1]+j]] = true
			}
		}

	}
	return true
}

func copyMap(source map[byte]bool) map[byte]bool {
	newMap := make(map[byte]bool, len(source))
	for k, v := range source {
		newMap[k] = v
	}
	return newMap
}

func IsValidSudoku2(board [][]byte) bool {
	rowMap := [9]map[byte]bool{
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
	}

	coloumnMap := [9]map[byte]bool{
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
		make(map[byte]bool, 9),
	}

	blockMap := map[[2]int]map[byte]bool{
		{0, 0}: make(map[byte]bool, 9),
		{1, 0}: make(map[byte]bool, 9),
		{2, 0}: make(map[byte]bool, 9),
		{0, 1}: make(map[byte]bool, 9),
		{1, 1}: make(map[byte]bool, 9),
		{2, 1}: make(map[byte]bool, 9),
		{0, 2}: make(map[byte]bool, 9),
		{1, 2}: make(map[byte]bool, 9),
		{2, 2}: make(map[byte]bool, 9),
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//row
			if board[i][j] == '.'{
				continue
			}
			if val, ok := rowMap[i][board[i][j]]; ok {
				if val {
					return false
				}
			}
			rowMap[i][board[i][j]] = true
			//coloumn
			if val, ok := coloumnMap[j][board[i][j]]; ok {
				if val {
					return false
				}
			}
			coloumnMap[j][board[i][j]] = true
			//block
			blockX := i / 3
			blockY := j / 3
			if val, ok := blockMap[[2]int{blockX, blockY}][board[i][j]]; ok {
				if val {
					return false
				}
			}
			blockMap[[2]int{blockX, blockY}][board[i][j]] = true
		}
	}

	return true
}


func IsValidSudoku3(board [][]byte) bool {
	rowMap := [9]map[byte]bool{
	}

	coloumnMap := [9]map[byte]bool{
	}

	blockMap := map[[2]int]map[byte]bool{
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//row
			if board[i][j] == '.'{
				continue
			}
			if rowMap[i] == nil{
				rowMap[i] = make(map[byte]bool, 9)
			}
			if val, ok := rowMap[i][board[i][j]]; ok {
				if val {
					return false
				}
			}
			rowMap[i][board[i][j]] = true
			//coloumn
			if coloumnMap[j] == nil{
				coloumnMap[j] = make(map[byte]bool, 9)
			}
			if val, ok := coloumnMap[j][board[i][j]]; ok {
				if val {
					return false
				}
			}
			coloumnMap[j][board[i][j]] = true
			//block
			blockX := i / 3
			blockY := j / 3
			if blockMap[[2]int{blockX, blockY}] == nil{
				blockMap[[2]int{blockX, blockY}] = make(map[byte]bool, 9)
			}
			if val, ok := blockMap[[2]int{blockX, blockY}][board[i][j]]; ok {
				if val {
					return false
				}
			}
			blockMap[[2]int{blockX, blockY}][board[i][j]] = true
		}
	}

	return true
}
