package src

func isMonotonic(A []int) bool {
	if len(A) <= 1{
		return true
	}
	res := true
	setFlag := 0
	for i := 1; i < len(A); i++ {

		if A[i]-A[i-1] == 0 {
			continue
		}
		if A[i]-A[i-1] < 0 {
			if setFlag == 0 {
				setFlag = -1 //des
			}
			if setFlag == 1 {
				res = false
				break
			}

		}

		if A[i]-A[i-1] > 0 {
			if setFlag == 0 {
				setFlag = 1 //des
			}
			if setFlag == -1 {
				res = false
				break
			}
		}

	}
	return res
}
