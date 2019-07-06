package src

import (
	"strings"
)

//a good   example
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	wordsBytes := []byte(s)
	resultBytes := make([]byte, 0, len(s))

	wordsEnd := len(wordsBytes) - 1
	beginLog := true
	for i := len(wordsBytes) - 1; i >= 0; i-- {
		if wordsBytes[i] == ' ' {
			if beginLog {
				resultBytes = append(resultBytes, wordsBytes[i+1:wordsEnd+1]...)
				resultBytes = append(resultBytes, ' ')
				//fmt.Println(string(resultBytes))
			}
			beginLog = false
			wordsEnd = i - 1
		} else {
			beginLog = true
		}

		if i == 0 && wordsBytes[i] != ' ' {
			resultBytes = append(resultBytes, wordsBytes[0:wordsEnd+1]...)
		}
	}

	return string(resultBytes)
}
