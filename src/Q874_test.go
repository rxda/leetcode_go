package src

import "testing"

func Test_robotSim(t *testing.T) {
	type args struct {
		commands  []int
		obstacles [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				commands:  []int{4, -1, 3},
				obstacles: [][]int{},
			},
			25,
		},
		{
			"2",
			args{
				commands:  []int{4,-1,4,-2,4},
				obstacles: [][]int{{2,4}},
			},
			65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotSim(tt.args.commands, tt.args.obstacles); got != tt.want {
				t.Errorf("robotSim() = %v, want %v", got, tt.want)
			}
		})
	}
}