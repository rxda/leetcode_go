package main

import (
	"fmt"
	"regexp"
)

func main(){
	decodeString("")
}

func decodeString(s string)  {
	reg, _ := regexp.Compile(`[a-z]*(\d\[[a-z]+])+[a-z]*`)
	a:= reg.FindAllStringSubmatch("3[a2[c]]",-1)
	for _,v :=range a[0]{
		fmt.Println(v)
	}
}