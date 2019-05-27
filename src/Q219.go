package src

func containsNearbyDuplicate(nums []int, k int) bool {
	nMap := make(map[int]*[2]int)
	for nk,v :=range nums{
		if val,ok := nMap[v];!ok{
			nMap[v] = &[2]int{nk,-1}
		}else{
			if val[1] == -1 {
				nMap[v][1] = nk
			}else{
				nMap[v][0], nMap[v][1] = nMap[v][1],nk
			}
		}
		if nMap[v][1] -nMap[v][0] <=k && nMap[v][1] != -1{
			return true
		}
	}
	return false
}