package src

func numPairsDivisibleBy60(time []int) int {
	res :=map[int]int{}
	for _, v := range time {
		remainder :=v%60
		if _,ok :=res[remainder];ok{
			res[remainder]++
		}else{
			res[remainder]=1
		}
	}
	count :=0
	for k,v := range res{
		if k>30 {
			continue
		}
		if k==0 ||k==30{
			if v<=1{
				continue
			}else{
				if v==2 {
					count+=1
				}else{
					count+= fact(v)/(2*fact(v-2))
				}
			}


		}else{
			count+=v*res[60-k]
		}
	}
	return count
}

func fact(n int) int{
	ji:=1
	for i:=n;i>1;i--{
		ji*=i
	}
	return ji
}