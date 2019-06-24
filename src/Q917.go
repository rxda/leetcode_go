package src

func reverseOnlyLetters(S string) string {
	res := make([]byte, len(S))
	chars := []byte{}
	bytes := []byte(S)

	for k, v := range bytes {
		if isChar(v) {
			chars = append(chars, v)
		} else {
			res[k] = v
		}
	}

	courser := len(chars) - 1
	for k, v := range res {
		if v == 0 {
			res[k] = chars[courser]
			courser--
		}
	}
	return string(res)
}




func reverseOnlyLetters2(S string) string {
	bytes := []byte(S)
	i := 0
	j := len(S) -1
	for i<j{
		if !isChar(bytes[i]){
			i++
			continue
		}
		if !isChar(bytes[j]){
			j--
			continue
		}
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

func isChar(c byte) bool{
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}else{
		return false
	}
}