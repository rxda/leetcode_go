package src

func hammingWeight(num uint32) int {
	var result int
	for num > 0 {
		num &= num-1
		result ++
	}
	return result
}

func hammingWeight2(num uint32) int {
	var a uint32 = 1
	sum :=0
	for i:=0;i<32;i++{
		if a&num !=0{
			sum++
		}
		a <<= 1
	}
	return sum
}