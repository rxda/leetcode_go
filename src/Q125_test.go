package src

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				"A man, a plan, a canal: Panama",
			},
			true,
		},
		{
			"2",
			args{
				"aaaa",
			},
			true,
		},
		{
			"3",
			args{
				"aabaa",
			},
			true,
		},
		{
			"4",
			args{
				"aabab",
			},
			false,
		},
		{
			"5",
			args{
				"\",",
			},
			true,
		},
		{
			"5",
			args{
				"0P",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
