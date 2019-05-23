package src

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for k,v :=range nums{
		numMap[v]=k
	}
	res :=make([]int,2)
	for k,v :=range nums{
		if val,ok:= numMap[target-v];!ok&&val!=k{
			res[0]=k
			res[1]=val
			break
		}
	}
	return res
}


//func twoSum2(nums []int, target int) []int {
//	numMap := make(map[int]int)
//	for k,v :=range nums{
//		diff := target - v
//		if _,ok:= numMap[diff];ok{
//			return []int{k,numMap[diff]}
//		}
//		numMap[v]=k
//		strings.tr
//	}
//	return []int{}
//}