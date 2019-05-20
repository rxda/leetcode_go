package src

func mySqrt(x int) int {
	min :=0
	max :=x
	for {
		i:= (max+min)/2
		if i*i< x {
			min = i
			max = x
		}
	}
}
