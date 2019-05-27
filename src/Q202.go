package src

import (
	"strconv"
)


func isHappy(n int) bool {
	has := make(map[int]bool)
	for n != 1 {
		n = intToSum(n)
		if _, ok := has[n]; ok {
			return false
		}else{
			has[n] =true
		}
	}
	return true
}

func intToSum( n int) int{
	str := []byte(strconv.Itoa(n))
	sum := 0
	for _, v := range str {
		sum += int(v - 48)*int(v - 48)
	}
	return sum
}