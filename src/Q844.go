package src

func backspaceCompare(S string, T string) bool {
	var stackS []int32
	var stackT []int32
	for _,v :=range S{
		if v == '#'{
			if length := len(stackS); length >0 {
				stackS = stackS[:length-1]
			}
		}else{
			stackS = append(stackS, v)
		}
	}

	for _,v :=range T{
		if v == '#'{
			if length := len(stackT); length >0 {
				stackT = stackT[:length-1]
			}
		}else{
			stackT = append(stackT, v)
		}
	}

	return string(stackS) == string(stackT)
}