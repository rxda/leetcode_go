package src

import (
	"fmt"
	"testing"
)

func TestAllOne_GetMinKey(t *testing.T) {
	a := Constructora()
	a.Inc("aaaa")
	a.Inc("aaaa")
	a.Inc("aaaa")
	a.Inc("bbbb")
	a.Inc("bbbb")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())
	a.Dec("aaaa")
	a.Dec("aaaa")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())

}
