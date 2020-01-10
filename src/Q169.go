package src

func majorityElement(nums []int) int {
	//2,1,2,1,1,1,2,2,2
	var candidate int
	var count int

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}
