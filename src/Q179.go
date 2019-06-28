package src

import (
	"sort"
	"strconv"
)

type customCompareString []string

func (s customCompareString) Len() int {
	return len(s)
}
func (s customCompareString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s customCompareString) Less(i, j int) bool {
	temp1 := s[i] + s[j]
	temp2 := s[j] + s[i]
	for p := 0; p < len(temp1); p++ {
		if diff := int(temp1[p]) - int(temp2[p]); diff < 0 {
			return false
		} else {
			if diff > 0 {
				return true
			} else {
				continue
			}
		}
	}
	return true
}

func largestNumber(nums []int) string {
	numStrs := make([]string, len(nums))
	for k, v := range nums {
		numStrs[k] = strconv.Itoa(v)
	}
	sort.Sort(customCompareString(numStrs))
	res := ""
	for _, v := range numStrs {
		res += v
	}
	bytes := []byte(res)
	length := len(bytes)
	for k,v := range bytes{
		if v == '0' && k != length - 1{
			bytes = bytes[1:]
		}else{
			break
		}
	}

	return string(bytes)
}
