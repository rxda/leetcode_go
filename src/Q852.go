package src

func PeakIndexInMountainArray(A []int) int {
	val := getVal(A)
	for k,v :=range A {
		if v == val {
			return k
		}
	}
	return -1
}

func getVal(A []int) int{
	length :=len(A)
	if length ==1{
		return A[0]
	}
	if length ==2{
		if A[0]>A[1]{
			return A[0]
		}else{
			return A[1]
		}
	}

	mid :=len(A)/2
	//top in left
	if A[mid]>A[mid+1] {
		return getVal(A[:mid+1])
	}else{
		return getVal(A[mid:])
	}
}

func PeakIndexInMountainArray3(A []int) int {
	length := len(A)
	if length == 1 {
		return 0
	}
	if length == 2 {
		if A[0] > A[1] {
			return 0
		} else {
			return 1
		}
	}
	for i, j := 0, length-1; i < length-1 && j > 0; i, j = i+1, j-1 {
		if A[i+1] < A[i] {
			return i
		}
		if A[j-1] < A[j] {
			return j
		}
	}
	return 0
}
