package src

import "testing"

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1",
			args{
				[]int{1,0,2,3,0,4,5,0},
			},
		},
		{"1",
			args{
				[]int{1,2,3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.args.arr)
		})
	}
}
