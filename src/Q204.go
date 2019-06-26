package src

import "math"

func countPrimes(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if isPrimes(i) {
			count++
		}
	}
	return count
}

func isPrimes(num int) bool {
	if num < 2{
		return false
	}
	if num ==2 {
		return true
	}
	tmp := int(math.Sqrt(float64(num)))+1
	for i:= 2;i< tmp;i++{
		if num %i == 0{
			return false
		}
	}
	return true
}
