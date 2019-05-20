package src

import (
	"reflect"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	tests := []struct {
		name string
		args string
		want [][]int
	}{
		// TODO: Add test cases.
		//{
		//	"1",
		//	"abbxxxxzyy",
		//	[][]int{{3, 6}},
		//},
		//{
		//	"2",
		//	"abbxxxxxxxxxxxxzyy",
		//	[][]int{{3, 14}},
		//},
		//{
		//	"2",
		//	"",
		//	[][]int{},
		//},
		//{
		//	"2",
		//	"s",
		//	[][]int{},
		//},
		//{
		//	"2",
		//	"abcdddeeeeaabbbcd",
		//	[][]int{{3, 5},{6, 9},{12,14}},
		//},
		{
			"2",
			"aaabbb",
			[][]int{{0, 2},{3,5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largeGroupPositions(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}
