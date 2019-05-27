package src

func hammingWeight(num uint32) int {
	var result int
	for num > 0 {
		num &= (num-1)
		result ++
	}
	return result
}