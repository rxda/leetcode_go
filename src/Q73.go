package src

func setZeroes(matrix [][]int)  {
	set := map[[2]int]struct{}{}
	x := len(matrix)//row
	y := len(matrix[0])//column
	for k1,v1 :=range matrix{
		for k2,v2 := range v1{
			if v2 == 0 {
				set[[2]int{k1,k2}] = struct{}{}
			}
		}
	}


	for k, _ :=range set{
		//行
		matrix[k[0]] = make([]int, y)
		//列
		for i:=0;i<x;i++{
			matrix[i][k[1]] = 0
		}

	}
}

func setZeroes2(matrix [][]int)  {
	xset := map[int]bool{}
	yset := map[int]bool{}
	x := len(matrix)//row
	y := len(matrix[0])//column
	for k1,v1 :=range matrix{
		for k2,v2 := range v1{
			if v2 == 0 {
				xset[k1] = true
				yset[k2] = true
			}
		}
	}

	for k, _ :=range xset{
		//行
		matrix[k] = make([]int, y)
	}
	for k, _ :=range yset{
		//列
		for i:=0;i<x;i++{
			matrix[i][k] = 0
		}
	}
}