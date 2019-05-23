package src

import (
	"sort"
)

func smallestRangeI(A []int, K int) int {
	sort.Ints(A)
	length := len(A)
	if A[length-1]-A[0] <= 2*K {
		return 0
	} else {
		return A[length-1] - A[0] - 2*K
	}
}


