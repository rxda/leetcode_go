package src

import "fmt"

func FindSubstring(s string, words []string) []int {
	length :=len(words)
	pair := make([]string,length*length)
	for i:=0;i<length;i++{
		for j:=0;j<length &&j!=i;j++{
			pair = append(pair,words[i]+words[j])
		}
	}
	fmt.Println(pair)
	result :=[]int{}
	for i:=0;i<len(s)-length*2;i++{
		for _,v :=range pair{
			if s[i:i+length*2]==v{
				result = append(result,i)
			}
		}

	}
	return result

}

