package src

import (
	"bytes"
	"strings"
)

func ReverseWords2(s string) string {
	strs :=strings.Split(s," ")
	res :=""
	for k,v :=range strs{
		if k==0{
			res = res +Reverse(v)
		}else{
			res = res+ " " +Reverse(v)
		}


	}
	return res
}

func Reverse(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseWords(s string) string {
	sr := Reverse(s)
	strs :=strings.Split(sr," ")
	var buffer bytes.Buffer
	length := len(strs)
	for i:=length-1;i>=0;i--{
		if i == 0{
			buffer.WriteString(strs[i])
		}else{
			buffer.WriteString(strs[i]+ " ")
		}
	}
	return buffer.String()
}