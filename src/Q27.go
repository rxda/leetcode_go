package src

func removeElement(nums []int, val int) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		if nums[0] == val{
			return 0
		}else{
			return 1
		}
	}

	for i:=0;i <len(nums);i++{
		if nums[i] == val{
			for j:=i;j <len(nums) -1;j++{
				nums[j] = nums[j+1]
			}
			nums = nums[:len(nums)-1]
			i--
		}
	}
	return len(nums)
}


func removeElement2(nums []int, val int) int {
	i :=0
	for j:=0;j<len(nums);j++{
		if nums[j] != val{
			nums[i] = nums[j]
			i++
		}
	}
	return i
}