package src

func addBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	bytesA, bytesB := []byte(a), []byte(b)
	diff := lenA - lenB
	length :=0
	if diff > 0 {
		length = lenA+1
		temp :=make([]byte, diff+1)
		for k,_ :=range temp{
			temp[k] ='0'
		}
		bytesB = append(temp, bytesB...)
		bytesA = append([]byte{'0'}, bytesA...)
	} else {
		length = lenB+1
		temp :=make([]byte, -diff+1)
		for k,_ :=range temp{
			temp[k] ='0'
		}
		bytesA = append(temp, bytesA...)
		bytesB = append([]byte{'0'}, bytesB...)
	}
	lastOver := false
	for i:=length -1 ;i>=0;i--{
		//全1
		if bytesA[i] == '1' && bytesB[i] == '1'{
			if !lastOver{
				bytesA[i] ='0'
			}else{
				bytesA[i] ='1'
			}
			lastOver =true
			continue
		}
		//全0
		if bytesA[i] == '0' && bytesB[i] == '0'{
			if !lastOver{
				bytesA[i] ='0'
			}else{
				bytesA[i] ='1'
			}
			lastOver =false
			continue
		}
		//1和0
		if !lastOver{
			bytesA[i] ='1'
			lastOver =false
		}else{
			bytesA[i] ='0'
			lastOver =true
		}
	}
	if bytesA[0] == '0'{
		bytesA = bytesA[1:]
	}
	return string(bytesA)
}
