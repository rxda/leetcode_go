package src

func IsPalindrome(x int) bool {
	//s := strconv.Itoa(x)
	//sf :=[]rune(s)
	//var res []rune
	//for i:= len(sf)-1;i>=0;i--{
	//	res = append(res, sf[i])
	//}
	//if string(res) == s {
	//	return true
	//}else{
	//	return false
	//}


	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber * 10 + x % 10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10

}
