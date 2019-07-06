package src

func trailingZeroes(n int) int {
	//1  1
	//2  2
	//3  6
	//4  24
	//5  120
	//6  720
	//8  5760
	//9  51480
	//10 514800
	//15 1,307,674,368,000
	if n < 5 {
		return 0
	}
	count :=0
	for n!= 0{
		n/=5
		count+=n
	}
	return count
}