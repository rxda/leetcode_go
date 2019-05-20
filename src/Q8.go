package src

import (
	"math"
	"strconv"
)

func MyAtoi(str string) int {
	chars :=[]rune(str)
	flag :=1
	firstFlag :=true
	var num []int
	for _,v :=range chars{
		val :=string(v)
		if val==" "&&firstFlag{
			continue
		}
		if val=="-"&&firstFlag{
			flag =-flag
			firstFlag = false
			continue
		}

		if val=="+"&&firstFlag{
			firstFlag = false
			continue
		}

		if i, err := strconv.Atoi(val);err!=nil{
			break
		}else{
			firstFlag = false
			num =append(num,i)
		}
	}
	res :=0
	if firstFlag ==true{
		return 0
	}
	length :=len(num)-1
	for _,v :=range num{
		res +=v*int(math.Pow(10,float64(length)))
		length-=1
	}
	res*=flag
	if res > 2147483647{
		return 2147483647
	}

	if len(str)>10 && flag ==1 {
		return 2147483647
	}

	if res < -2147483648{
		return -2147483648
	}
	if len(str)>11 &&flag == -1 {
		return -2147483647
	}

	return res
}