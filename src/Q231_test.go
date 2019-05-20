package src

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		// TODO: Add test cases.
		{
			"123",
			123,
			false,
		},
		{
			"22",
			22,
			false,
		},
		{
			"123",
			1024,
			true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
