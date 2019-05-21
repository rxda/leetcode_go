package main

func main() {
	
}

func sortArrayByParityII(A []int) []int {
	result :=make([]int,len(A))
	odd,even :=0,1
	for _,v :=range A{
		if v%2==0{
			result[even]=v
			even+=2
		}else{
			result[odd]=v
			odd+=2
		}
	}
	return  result
}