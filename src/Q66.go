package src

func PlusOne(digits []int) []int {
	digits = append([]int{0},digits...)
	length :=len(digits)
	var result = make([]int,length)
	for i:=length-1;i>=0;i-- {
		if i==length-1{
			t := digits[i]+1
			if t==10{
				digits[i-1]++
				digits[i] = 0
			}else{
				digits[i]=t
				result = digits
				break
			}
		}else{
			if digits[i]==10{
				digits[i-1]++
				digits[i] = 0
			}else{
				result = digits
				break
			}
		}
	}
	if result[0]==0{
		return result[1:]
	}else{
		return result
	}

}