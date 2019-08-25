package src

import "testing"

func Test_numRollsToTarget(t *testing.T) {
	type args struct {
		d      int
		f      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				d:      1,
				f:      6,
				target: 3,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRollsToTarget(tt.args.d, tt.args.f, tt.args.target); got != tt.want {
				t.Errorf("numRollsToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}