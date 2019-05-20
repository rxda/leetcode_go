package src

func largeGroupPositions(S string) [][]int {
	if len(S) < 3{
		return [][]int{}
	}
	bytes := []byte(S)

	count := 0
	length := len(bytes)
	start := 0
	end := 0
	result :=[][]int{}
	for i := 0; i < length-1; i++ {
		if bytes[i] == bytes[i+1] {
			if count == 1 || count ==0 {
				count = 2
			}else{
				count ++
				if count == 3 {
					start = i - 3 + 2
					end = i+1
				}
				if count > 3 {
					end = i+1
				}
				if i == length -2 {
					result = append(result, []int{start, end})
				}
			}
		} else {
			if count >= 3 {
				result = append(result, []int{start, end})
			}
			count = 1
		}
	}
	return result
}
