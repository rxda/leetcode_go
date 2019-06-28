package src

func sortColors(nums []int)  {
	for i:=0;i< len(nums); i++{
		for j :=i+1;j<len(nums);j++{
			if nums[i] > nums[j]{
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}


func sortColors2(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	colorMap := make(map[int]int, 3)
	for _,v :=range nums{
		colorMap[v]++
	}
	for k,_ :=range nums{
		if colorMap[0] != 0{
			nums[k] = 0
			colorMap[0]--
			continue
		}
		if colorMap[1] != 0{
			nums[k] = 1
			colorMap[1]--
			continue
		}
		if colorMap[2] != 0{
			nums[k] = 2
			colorMap[2]--
		}
	}
}