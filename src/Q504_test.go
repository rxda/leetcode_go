package src

import (
	"strconv"
	"testing"
)

func TestConvertToBase7(t *testing.T) {
	tests := []struct {
		name string
		args int
		want string
	}{
		// TODO: Add test cases.
		{"1",
			0,
			strconv.FormatInt(0,7),
		},
		{"0",
			6,
			strconv.FormatInt(6,7),
		},
		{"100",
			100,
			strconv.FormatInt(100,7),
		},
		{"200",
			200,
			strconv.FormatInt(200,7),
		},
		{"310",
			310,
			strconv.FormatInt(310,7),
		},
		{"100001",
			100001,
			strconv.FormatInt(100001,7),
		},
		{"49",
			49,
			strconv.FormatInt(49,7),
		},
		{"-100",
			-100,
			strconv.FormatInt(-100,7),
		},
		{"-100005",
			-100005,
			strconv.FormatInt(-100005,7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToBase7(tt.args); got != tt.want {
				t.Errorf("ConvertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}
