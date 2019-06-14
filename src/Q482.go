package src

import "strings"

func licenseKeyFormatting(S string, K int) string {
	S = strings.Replace(S, "-", "", -1)
	S = strings.ToUpper(S)
	bytes := []byte(S)
	result := []byte{}
	length := len(S)
	if length == 0 {
		return ""
	}
	if length <= K {
		return S
	}

	//pair := length / K
	remain := length % K
	if remain != 0 {
		for i := 0; i < remain; i++ {
			result = append(result, bytes[i])
		}
		result = append(result, '-')
	}

	t := 0
	for i := remain; i < length; i++ {
		if t == K {
			result = append(result, '-')
			t = 0
		}
		result = append(result, bytes[i])
		t++
	}
	return string(result)
}