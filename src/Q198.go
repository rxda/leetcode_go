package src

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	prePre, pre := 0, nums[0]

	for k, v := range nums {
		if k ==0{
			continue
		}
		cur := max(prePre + v, pre)
		prePre =pre
		pre = cur
	}
	return pre
}

func max(x, y int) int{
	if x> y{
		return x
	}else{
		return y
	}
}
