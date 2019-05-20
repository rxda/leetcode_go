package src

import (
	"sort"
)

func AllCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {

	var m = make(map[int][][]int)
	var keys []int
	for i:=0;i<R;i++{
		for j:=0;j<C;j++{
			t :=abs(i-r0)+abs(j-c0)
			m[t] = append(m[t], []int{i,j})
			keys = append(keys, t)
		}
	}
	sort.Ints(keys)
	l :=-1
	var result [][]int
	for _, v := range keys {
		if v==l{
			continue
		}else{
			for _,pair :=range m[v]{
				if len(pair)!=0{
					result = append(result, pair)
				}
			}
			l=v
		}

	}
	return result
}