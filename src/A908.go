package main

import (
	"fmt"
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

func main() {
	A := []int{0, 10}
	fmt.Println(smallestRangeI(A, 2))
}
