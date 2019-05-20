package src

import (
	"strings"
)

func LengthOfLastWord(s string) int {

	s = strings.Trim(s, " ")
	if s == ""{
		return 0
	}
	ss := strings.Split(s, " ")

	return len(ss[len(ss)-1])
}
