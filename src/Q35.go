package main

import "fmt"

func main() {
	a :=searchInsert([]int{1,3,4,6},5)
	fmt.Println(a)
}

func searchInsert(nums []int, target int) int {
	if nums == nil {
		return -1 // Err
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	return low

}