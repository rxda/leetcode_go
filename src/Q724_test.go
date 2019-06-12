package src

import "testing"

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				[]int{1,7,3,6,5,6},
			},
			3,
		},
		{
			"2",
			args{
				[]int{1,2,3},
			},
			-1,
		},
		{
			"3",
			args{
				[]int{1,2,3,4,5,6,7,8,9},
			},
			3,
		},
		{
			"4",
			args{
				[]int{1,1,1,1,1,1,1},
			},
			3,
		},
		{
			"4",
			args{
				[]int{-5,-4,-3,-2,-1,0,-1,-2,-3,-4,-5},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
