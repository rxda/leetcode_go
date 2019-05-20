package src

import "fmt"

func MaxSumTwoNoOverlap(A []int, L int, M int) int{
	length := len(A)
	r1 := make(map[[2]int]int)
	rf := 0
	if L==1 {
		for k,v :=range A {
			r1[[2]int{k,k}] =v
		}
	}else{
		for i := 0; i < L; i++ {
			rf += A[i]
		}
		for i := 0; i <= length-L; i++ {
			if i == 0 {
				r1[[2]int{i, i + L-1}] = rf
			} else {
				rf = rf - A[i-1] + A[i-1+L]
				r1[[2]int{i, i + L-1}] = rf
			}
		}
	}


	r2 := make(map[[2]int]int)
	rf = 0
	if M==1 {
		for k,v :=range A {
			r2[[2]int{k,k}] =v
		}
	}else{
		for i := 0; i < M; i++ {
			rf += A[i]
		}
		for i := 0; i <= length-M; i++ {
			if i == 0 {
				r2[[2]int{i, i + M-1}] = rf
			} else {
				rf = rf - A[i-1] + A[i-1+M]
				r2[[2]int{i, i + M-1}] = rf

			}
		}
	}


	max := 0
	for k1, v1 := range r1 {
		for k2, v2 := range r2 {
			if !(k2[0]<=k1[1] && k1[0]<=k2[1]) {
				if s := v1 + v2; s > max {
					fmt.Println(k1,k2)
					max = s
				}
			}
		}
	}
	return max
}
