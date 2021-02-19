package src

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_createList(t *testing.T) {
	a := createList([]int{5, 4, 3, 2, 1})
	fmt.Println(a)
}

func Test_deleteDuplicates2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		//{
		//	"1",
		//	args{head: createList([]int{1, 1, 1, 2, 3, 4})},
		//	createList([]int{2, 3, 4}),
		//},
		{
			"2",
			args{head: createList([]int{1, 1, 1, 2, 3, 4, 5, 5})},
			createList([]int{2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}
