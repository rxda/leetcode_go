package src

import "strings"

//此题描述有误，测试用例中1A被认为是false,A1A为true
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)
	bytes := []byte(s)
	length := len(bytes)
	i, j := 0, length-1
	for i < j {
		for !isNumOrChar(bytes[i]) {
			i++
			if i > length-1 {
				return true
			}
		}
		for !isNumOrChar(bytes[j]) {
			j--
			if j < 0 {
				return true
			}
		}

		if bytes[i] == bytes[j] {
			i++
			j--
			continue
		} else {
			return false
		}
	}
	return true
}

func isNumOrChar(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')|| (c >= '0' && c <= '9') {
		return true
	}else{
		return false
	}
}
