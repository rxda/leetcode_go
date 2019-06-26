package src

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{3,2,2,3},
				3,
			},
			2,
		},
		{
			"2",
			args{
				[]int{0,1,2,2,3,0,4,2},
				2,
			},
			5,
		},
		{
			"3",
			args{
				[]int{1},
				1,
			},
			0,
		},
		{
			"4",
			args{
				[]int{1},
				2,
			},
			1,
		},
		{
			"5",
			args{
				[]int{1,3},
				2,
			},
			2,
		},
		{
			"6",
			args{
				[]int{1,3},
				3,
			},
			1,
		},
		{
			"7",
			args{
				[]int{1,3},
				1,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement2(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
