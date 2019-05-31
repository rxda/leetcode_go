package src

func pivotIndex(nums []int) int {
	switch len(nums){
	case 0:
		return -1
	case 1:
		return 0
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum -nums[0] == 0{
		return 0
	}
	leftSum, rightSum := nums[0], sum-nums[0]

	for k, v := range nums[1:] {
		rightSum -= v
		if leftSum == rightSum {
			return k + 1
		}
		leftSum += v
	}
	return -1
}

