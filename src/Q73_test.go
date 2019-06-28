package src

import "testing"

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				[][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
		},
		{
			"2",
			args{
				[][]int{
					{0,1,2,0},
					{3,4,5,2},
					{1,3,1,5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
		})
	}
}



func BenchmarkSetZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}
		setZeroes(a)
	}
}

func BenchmarkSetZeroes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}
		setZeroes2(a)
	}
}