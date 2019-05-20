package src

import "strconv"

func ConvertToBase7(num int) string {
	var result string
	if num <0 {
		num = 0-num
		result = "-"
	}else{
		result = ""
	}

	if num < 7 {
		return result+strconv.Itoa(num)
	}

	baseMap := [12]int{}
	baseMap[0] = 1
	maxBit := 0
	for i := 1; i <= 11; i++ {
		baseMap[i] = 7 * baseMap[i-1]
		if baseMap[i] > num {
			maxBit = i - 1
			break
		} else {
			if baseMap[i] == num {
				result += "1"
				for j := 0; j < i; j++ {
					result += "0"
				}
				return result
			}
		}
	}


	for n := maxBit; n >= 0; n-- {
		t := 0
		firstNotZeroFlag := false
		for i := 1; i <= 6; i++ {
			if i*baseMap[n] > num {
				firstNotZeroFlag = true
				t = i - 1
				num -= t * baseMap[n]
				break
			}
			if i == 6 {
				firstNotZeroFlag = true
				t = 6
				num -= t * baseMap[n]
			}
		}
		if firstNotZeroFlag {
			result += strconv.Itoa(t)
		}
	}

	return result
	//for
}
