package src

func MaxSubArray(nums []int) int {
	length :=len(nums)
	for i :=1;i<length ; i++{
		if nums[i-1]>0 {
			nums[i] +=nums[i-1]
		}
	}
	max :=-10000
	for _,v :=range nums{
		if v>max {
			max =v
		}
	}
	return max
}