package src

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				S:"Qedo1ct-eeLg=ntse-T!",
			},
			"Test1ng-Leet=code-Q!",
		},
		{
			"2",
			args{
				S:"!",
			},
			"!",
		},
		{
			"3",
			args{
				S:"a-bC-dEf-ghIj",
			},
			"j-Ih-gfE-dCba",
		},
		{
			"4",
			args{
				S:"",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters2(tt.args.S); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
