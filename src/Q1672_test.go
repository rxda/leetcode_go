package src

import (
	"fmt"
	"testing"
)

func Test_maximumWealth(t *testing.T) {
	accounts := [][]int{
		{2,8,7},
		{7,1,3},
		{1,9,5},
	}
	res := maximumWealth(accounts)
	fmt.Println(res)
}
