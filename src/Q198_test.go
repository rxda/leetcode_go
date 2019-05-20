package src

import "testing"

func Test_rob(t *testing.T) {
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
				[]int{2,1,1,1,2},
			},
			4,
		},
		//{
		//	"1",
		//	args{
		//		[]int{2,7,9,3,1},
		//	},
		//	5,
		//},
		//{
		//	"1",
		//	args{
		//		[]int{1,2,3,4,5,6,7,8,9},
		//	},
		//	5,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
