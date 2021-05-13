package src

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "5",
			args: args{numRows: 5},
			want: [][]int{
				{1},
				{1,1},
				{1,2,1},
				{1,3,3,1},
				{1,4,6,4,1},
			},
		},
		{
			name: "1",
			args: args{numRows: 1},
			want: [][]int{
				{1},
			},
		},
		{
			name: "6",
			args: args{numRows: 6},
			want: [][]int{
				{1},
				{1,1},
				{1,2,1},
				{1,3,3,1},
				{1,4,6,4,1},
				{1,5,10,10,5,1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
