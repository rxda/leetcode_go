package src

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				"  a    good   example   ",
			},
			"example good a",
		},
		{
			"2",
			args{
				"  hello world!  ",
			},
			"world! hello",
		},
		{
			"4",
			args{
				"the sky is blue",
			},
			"blue is sky the",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
