package src

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				"/../",
			},
			"/",
		},
		{
			"2",
			args{
				"/a/b/c/..",
			},
			"/a/b",
		},
		{
			"3",
			args{
				"/a/./b/../../c/",
			},
			"/c",
		},
		{
			"4",
			args{
				"/a//b////c/d//././/..",
			},
			"/a/b/c",
		},
		{
			"5",
			args{
				"/",
			},
			"/",
		},
		{
			"6",
			args{
				"/a/../../b/../c//.//",
			},
			"/c",
		},
		{
			"7",
			args{
				"///////..//////..////././////////../",
			},
			"/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}


func BenchmarkSimplifyPath(b *testing.B) {
	a := "/a/../../b/../c//.//"
	for i := 0; i < b.N; i++ {
		simplifyPath(a)
	}
}

func BenchmarkSimplifyPath2(b *testing.B) {
	a := "/a/../../b/../c//.//"
	for i := 0; i < b.N; i++ {
		simplifyPath2(a)
	}
}

func BenchmarkSimplifyPath3(b *testing.B) {
	a := "/a/../../b/../c//.//"
	for i := 0; i < b.N; i++ {
		simplifyPath3(a)
	}
}