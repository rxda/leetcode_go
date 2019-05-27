package src

import "testing"

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		name string
		args uint32
		want int
	}{
		// TODO: Add test cases.
		{
			"00000000000000000000000101001011",
			00000000000000000000000101001011,
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
