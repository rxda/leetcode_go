package src

import (
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	//fmt.Println(PeakIndexInMountainArray([]int{0,1,2,3,4,5,6,1}))
	data := []struct {
		input  []int
		ans int
	}{
		{[]int{0,1,2,3,4,5,6,1},6},
		{[]int{0,8,7,6,5,4,3,2},1},
		{[]int{0},0},
		{[]int{1,2},1},
		{[]int{2,1},0},
		{[]int{8,7,6,5,4,3,2,1},0},
		{[]int{1,2,3,4,5,6,7,8,7,6,5,2},7},

	}
	for _,v :=range data{
		if c:=PeakIndexInMountainArray(v.input); c!=v.ans {
			t.Error(v,c)
		}
	}
}

func BenchmarkPeakIndexInMountainArray(b *testing.B) {
	in := []int{}
	for i:=0;i<1000;i++{
		in =append(in,i)
	}
	for i:=800;i>0;i--{
		in =append(in,i)
	}
	ans :=999
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		if c:=PeakIndexInMountainArray(in);c!=ans {
			b.Error(c,ans)
		}
	}
}
