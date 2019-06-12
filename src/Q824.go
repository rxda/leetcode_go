package src

import "strings"

func ToGoatLatin(S string) string {
	//元音
	vowels := map[uint8]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}

	words := strings.Split(S," ")
	for k,v :=range words{
		if _, ok := vowels[v[0]]; !ok{
			v= v[1:] + string(v[0])
		}
		v += "ma"
		for i:=0;i<k+1;i++{
			v += "a"
		}
		words[k] =v
	}
	return strings.Join(words, " ")
}
