package src

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				1,
			},
			0,
		},
		{
			"6",
			args{
				6,
			},
			1,
		},
		{
			"13",
			args{
				13,
			},
			2,
		},
		{
			"30",
			args{
				30,
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
