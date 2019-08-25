package src

func calculateTime(keyboard string, word string) int {
	chars := make([]int, 26)
	for k, v := range keyboard {
		chars[v - 'a'] = k
	}
	sum := 0
	var last = int32(keyboard[0]) - 'a'
	for _, v := range word {
		if diff := chars[v- 'a'] -chars[last- 'a'];diff >0{
			sum += diff
		}else {
			sum -= diff
		}
		last = v- 'a'
	}
	return sum
}
