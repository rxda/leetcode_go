package src

import "fmt"

func duplicateZeros(arr []int)  {
	length := len(arr)
	for i:=0;i< len(arr);i++{
		if arr[i] == 0{
			arr = append(append(arr[0:i],0),arr[i:]...)
			i++
		}
	}
	arr = arr[:length]
	fmt.Print(arr)
}

func duplicateZeros2(arr []int)  {
	var res []int
	for i:=0;i< len(arr);i++{
		if arr[i] == 0{
			res = append(res, 0,0)
		}else{
			res = append(res, arr[i])
		}
	}
	arr = res
}