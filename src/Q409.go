package src

func longestPalindrome(s string) int {
	bytes := []byte(s)
	byteMap := make(map[byte]int, len(bytes))
	for _,v :=range bytes{
		if _,ok:=byteMap[v];!ok{
			byteMap[v] = 1
		}else{
			byteMap[v]++
		}
	}
	single := false
	result := 0
	for _,v :=range byteMap{
		if s :=v/2; s>0 {
			result += s*2
		}
		if v%2 != 0{
			single =true
		}
	}
	if single{
		result++
	}
	return result
}
