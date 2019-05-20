package src

func IntToRoman(num int) string {
	res := ""
	romanUnit := []int{900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	// romanUnit =[]int{5,10,50,100,500,1000}

	numX := num - 1000*(num/1000)
	for k, v := range romanUnit {
		romanCount := numX / v
		numX -= romanCount * v
		if romanCount != 4 {
			for i := 0; i < romanCount; i++ {
				res += roman[k]
			}
		} else {
			res += roman[k-1] + roman[k]
		}

	}
	for i := 0; i < num/1000; i++ {
		res = "M" + res
	}
	return res
}
