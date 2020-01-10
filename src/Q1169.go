package src

func invalidTransactions(transactions []string) []string {
	//if len(transactions) == 1 {
	//	return []string{transactions[0]}
	//}
	//cityOneMap, cityTwoMap := make(map[string]string), make(map[string]string)
	//
	//var split = make([][]string, 16)
	//for i := 0; i < len(transactions); i++ {
	//	temp := strings.Split(transactions[i], ",")
	//	split = append(split, temp)
	//}
	//var result = make(map[string]struct{})
	//cityOneMap[split[0][1]] = transactions[0]
	//cityOneMap[split[0][3]] = transactions[0]
	//for i := 1; i < len(split); i++ {
	//	i1, _ := strconv.Atoi(split[i][1])
	//	i2, _ := strconv.Atoi(split[i-1][1])
	//	v1, ok1 := cityOneMap[split[i][1]]
	//	v2, ok2 := cityOneMap[split[i][1]]
	//	if !ok1 {
	//		result[v1] = struct{}{}
	//		result[transactions[i]] = struct{}{}
	//	}
	//	if !ok2 {
	//		result[v2] = struct{}{}
	//		result[transactions[i]] = struct{}{}
	//	}
	//	if abs(i2 -i1) <60 {
	//		result[v2] = struct{}{}
	//	}
	//
	//}
	//return result
	return []string{}
}
