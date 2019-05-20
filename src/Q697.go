package src

import "math"

func FindShortestSubArray(nums []int) int {
	if len(nums) ==1 {
		return 1
	}
	if len(nums) == 0{
		return 0
	}
	var imap = make(map[int][]int)
	for k,v :=range nums{
		imap[v] = append(imap[v],k)
	}
	maxlen :=0
	for _,v :=range imap{
		if maxlen< len(v){
			maxlen =len(v)
		}
	}
	if maxlen ==1 {
		return 1
	}
	maxk :=[]int{}

	for k,v :=range imap{
		if maxlen== len(v){
			maxk =append(maxk,k)
		}
	}
	minlength := math.MaxInt32
	for _,v :=range maxk{
		if len(imap[v])>1 {
			length :=imap[v][len(imap[v])-1]-imap[v][0]+1
			if  length < minlength{
				minlength = length
			}
		}

	}
	return minlength
}

