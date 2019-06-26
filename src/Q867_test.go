package src

import (
	"reflect"
	"testing"
)

func Test_transpose(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			[][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
		{
			"2",
			args{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			[][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
		{
			"3",
			args{
				[][]int{
					{1},
					{4},
				},
			},
			[][]int{
				{1, 4},
			},
		},
		{
			"4",
			args{
				[][]int{
					{1},
				},
			},
			[][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transpose(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
