package src

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result = [][]int{}
	numLen := len(nums)
	for i := 0; i < numLen; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo := i + 1
			hi := numLen - 1
			sum := -nums[i]
			for lo < hi {
				if nums[lo]+nums[hi] == sum {
					result = append(result, []int{nums[i], nums[lo], nums[hi]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else {
					if nums[lo]+nums[hi] < sum {
						lo++
					} else {
						hi--
					}
				}
			}
		}
	}
	return result
}
