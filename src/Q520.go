package src

func DetectCapitalUse(word string) bool {
	runes := []byte(word)
	length := len(runes)
	first := runes[0] >= 'A' && runes[0] <= 'Z'
	if length ==1 {
		return true
	}
	upperSum := 0
	for i := 1; i < length; i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			upperSum++
		}
	}
	if first {
		if upperSum == length-1 || upperSum==0{
			return true
		}else{
			return false
		}
	}else{
		if upperSum == 0 {
			return true
		}else{
			return false
		}
	}
}
