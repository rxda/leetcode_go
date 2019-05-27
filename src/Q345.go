package src

//aeiou -> uoiea

func reverseVowels(s string) string {
	bytes := []byte(s)
	length := len(bytes)
	if length <= 1{
		return s
	}
	i,j := 0, length - 1
	for i < j{
		if isVowel(bytes[i]){
			if isVowel(bytes[j]) {
				bytes[i],bytes[j] = bytes[j],bytes[i]
				i++
			}
			j--
		}else{
			i++
		}
	}
	return string(bytes)
}

func isVowel(b byte) bool{
	if b== 'a' || b =='e' || b== 'i'|| b=='o'|| b=='u' ||b== 'A' || b =='E' || b== 'I'|| b=='O'|| b=='U'{
		return true
	}else{
		return false
	}
}