package src

func isPowerOfTwo(n int) bool {
	if n == 1{
		return true
	}
	for{
		r := n/2
		//not integer
		if r*2 != n{
			return false
		}else{
			n = r
		}
		if n == 1{
			return true
		}
	}
}
