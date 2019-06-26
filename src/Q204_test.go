package src

import "testing"

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"10",
			args{
				10,
			},
			4,
		},
		{
			"20",
			args{
				20,//2,3,5,7,11,13,17,19
			},
			8,
		},
		{
			"30",
			args{
				30,//2,3,5,7,11,13,17,19,23,29
			},
			10,
		},
		{
			"40",
			args{
				40,//2,3,5,7,11,13,17,19,23,29,31,37
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
