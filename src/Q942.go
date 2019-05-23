package src



func diStringMatch(S string) []int {
	length := len(S)
	result := []int{}
	max := length
	min := 0
	for _, v := range S {
		if string(v) == "I" {
			result = append(result, min)
			min++
		} else {
			result = append(result, max)
			max--
		}

	}
	result = append(result, min)
	return result
}
