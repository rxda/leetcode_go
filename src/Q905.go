package src



func sortArrayByParity(A []int) []int {
	count := len(A)
	result := make([]int, count)
	i,j:=0,count-1
	for _, v := range A {
		if v%2 == 0 {
			result[i] = v
			i++
		} else {
			result[j] = v
			j--
		}
	}
	return result
}

func sortArrayByParity2(A []int) []int {

	odd,even := []int{},[]int{}

	for _, v := range A {
		if v%2 == 0 {
			even =append(even,v)
		} else {
			odd = append(odd,v)
		}
	}
	return append(even,odd...)
}