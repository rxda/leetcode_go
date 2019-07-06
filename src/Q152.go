package src

func maxProduct(nums []int) int {
	min, max, res := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			min , max= intMin(nums[i], nums[i]*min), intMax(nums[i], nums[i]*max)
		}else{
			min , max= intMin(nums[i], nums[i]*max), intMax(nums[i], nums[i]*min)
		}
		res = intMax(res, max)
	}
	return res
}

func intMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func intMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
